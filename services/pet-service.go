package services

import (
	"context"
	"find-a-friend/repositories"
	"find-a-friend/types"
	"fmt"
)

type PetService struct {
	petRepo repositories.PetRepository
}

func NewPetService(repo repositories.PetRepository) (*PetService) {
	return &PetService{petRepo: repo}
}

func (s *PetService) Create(ctx context.Context, pet types.CreatePet) (types.Pet, error) {
	// todo: check if org id is valid -> create org
	// test creating a pet with an invalid org id
	newPet, err := s.petRepo.Create(ctx, pet)
	if err != nil {
		return types.Pet{}, fmt.Errorf("error when creating a new pet: %v", err)
	}

	return newPet, nil
}
