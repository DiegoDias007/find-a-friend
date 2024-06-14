package psql

import (
	"context"
	"find-a-friend/types"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type PetRepository struct {
	db *pgx.Conn
}

func NewPetRepository(db *pgx.Conn) *PetRepository {
	return &PetRepository{
		db: db,
	}
}

func (r *PetRepository) Create(ctx context.Context, pet types.CreatePet) (types.Pet, error) {
	var createdPet types.Pet

	query := `
        INSERT INTO pet (name, city, species, breed, height, weight, org_id)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id, name, city, species, breed, height, weight, org_id
    `

	err := r.db.QueryRow(ctx, query, pet.Name, pet.City, pet.Species, pet.Breed, pet.Height, pet.Weight, pet.OrgId).
		Scan(&createdPet.Id, &createdPet.Name, &createdPet.City, &createdPet.Species, &createdPet.Breed, &createdPet.Height, &createdPet.Weight, &createdPet.OrgId)
	if err != nil {
		return types.Pet{}, fmt.Errorf("error creating pet: %v", err)
	}
	defer r.db.Close(ctx)

	return createdPet, nil
}

func (r *PetRepository) GetFromCity(ctx context.Context, filter types.PetFilter) ([]types.Pet, error) {
	query := "SELECT * FROM pet WHERE city = $1"
	args := []interface{}{filter.City}
	argIdx := 2

	if filter.Species != nil {
		query += fmt.Sprintf(" AND species = $%d", argIdx)
		args = append(args, *filter.Species)
		argIdx++
	}
	if filter.Breed != nil {
		query += fmt.Sprintf(" AND breed = $%d", argIdx)
		args = append(args, *filter.Breed)
		argIdx++
	}
	if filter.Height != nil {
		query += fmt.Sprintf(" AND height >= $%d", argIdx)
		args = append(args, *filter.Height)
		argIdx++
	}
	if filter.Weight != nil {
		query += fmt.Sprintf(" AND weight >= $%d", argIdx)
		args = append(args, *filter.Weight)
		argIdx++
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error when fetching pets: %v", err)
	}
	defer rows.Close()

	var pets []types.Pet
	for rows.Next() {
		var pet types.Pet
		err := rows.Scan(&pet.Id, &pet.Name, &pet.City, &pet.Species, &pet.Breed, &pet.Height, &pet.Weight, &pet.OrgId)
		if err != nil {
			return nil, fmt.Errorf("error scanning pet row: %w", err)
		}
		pets = append(pets, pet)
	}

	return pets, nil
}


func (r *PetRepository) GetById(context context.Context, id int) (types.Pet, error) {
	var pet types.Pet
	err := r.db.QueryRow(context, "SELECT * FROM pet WHERE id = $1", id).
		Scan(&pet.Id, &pet.Name, &pet.City, &pet.Species, &pet.Breed, &pet.Height, &pet.Weight, pet.OrgId)

	if err != nil {
		return types.Pet{}, fmt.Errorf("error when fetching pets: %v", err)
	}
	defer r.db.Close(context)

	return pet, nil
}
