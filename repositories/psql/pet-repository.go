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

func (p *PetRepository) Create(ctx context.Context, pet types.CreatePet) (types.Pet, error) {
    var createdPet types.Pet

    query := `
        INSERT INTO pet (name, city, species, breed, height, weight, org_id)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id, name, city, species, breed, height, weight, org_id
    `

    err := p.db.QueryRow(ctx, query, pet.Name, pet.City, pet.Species, pet.Breed, pet.Height, pet.Weight, pet.Org_id).
        Scan(&createdPet.Id, &createdPet.Name, &createdPet.City, &createdPet.Species, &createdPet.Breed, &createdPet.Height, &createdPet.Weight, &createdPet.Org_id)
    if err != nil {
        return types.Pet{}, fmt.Errorf("error creating pet: %w", err)
    }

    return createdPet, nil
}

func (p *PetRepository) GetFromCity(context context.Context, city string) ([]types.Pet, error) {
	rows, err := p.db.Query(context, "SELECT * FROM pet WHERE city = $1", city)
	if err != nil {
		return nil, fmt.Errorf("error when fetching pets: %v", err)
	}
	defer p.db.Close(context)

	var pets []types.Pet
	for rows.Next() {
		var pet types.Pet
		err := rows.Scan(&pet.Id, &pet.Name, &pet.City, &pet.Species, &pet.Breed, &pet.Height, &pet.Weight, &pet.Org_id)
		if err != nil {
			return nil, fmt.Errorf("error scanning pet row: %w", err)
		}
		pets = append(pets, pet)
	}

	return pets, nil

}
