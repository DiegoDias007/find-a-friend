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

func (r *PetRepository) GetFromCity(context context.Context, city string) ([]types.Pet, error) {
	rows, err := r.db.Query(context, "SELECT * FROM pet WHERE city = $1", city)
	if err != nil {
		return nil, fmt.Errorf("error when fetching pets: %v", err)
	}
	defer r.db.Close(context)

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
