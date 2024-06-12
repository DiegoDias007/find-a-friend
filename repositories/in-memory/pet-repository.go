package inMemory

import (
	"context"
	"find-a-friend/types"
	"find-a-friend/utils"
	"fmt"
)

type PetRepository struct {
	pets []types.Pet
}

func NewPetRepository() *PetRepository {
	return &PetRepository{}
}

func (r *PetRepository) Create(ctx context.Context,pet types.CreatePet) (types.Pet, error) {
	id := utils.GenerateRandomNumber(1000)

	newPet := types.Pet{
		Id:      id,
		Name:    pet.Name,
		City:    pet.City,
		Species: pet.Species,
		Breed:   pet.Breed,
		Height:  pet.Height,
		Weight:  pet.Weight,
		OrgId:   pet.OrgId,
	}

	r.pets = append(r.pets, newPet)

	return newPet, nil
}

func (r *PetRepository) GetFromCity(ctx context.Context, city string) ([]types.Pet, error) {
	var pets []types.Pet
	for _, pet := range r.pets {
		if pet.City == city {
			pets = append(pets, pet)
		}
	}

	return pets, nil
}

func (r *PetRepository) GetById(ctx context.Context, id int) (types.Pet, error) {
	for _, pet := range r.pets {
		if pet.Id == id {
			return pet, nil
		}
	}

	return types.Pet{}, fmt.Errorf("error getting pet by id, pet not found.")
}