package services

import (
	"context"
	"find-a-friend/repositories"
	"find-a-friend/types"
)

type PetService struct {
	petRepo repositories.PetRepository
	orgRepo repositories.OrgRepository
}

func NewPetService(repo repositories.PetRepository, orgRepo repositories.OrgRepository) (*PetService) {
	return &PetService{petRepo: repo, orgRepo: orgRepo}
}

func (s *PetService) Create(ctx context.Context, pet types.CreatePet) (types.Pet, error) {
	_, err := s.orgRepo.GetById(ctx, pet.OrgId)
	if err != nil {
		return types.Pet{}, err
	}

	newPet, err := s.petRepo.Create(ctx, pet)
	if err != nil {
		return types.Pet{}, err
	}

	return newPet, nil
}

func (s *PetService) GetFromCity(ctx context.Context, city string) ([]types.Pet, error) {
	pets, err := s.petRepo.GetFromCity(ctx, city)
	if err != nil {
		return nil, err
	}

	return pets, nil
}

func (s *PetService) GetById(ctx context.Context, id int) (types.Pet, error) {
	pet, err := s.petRepo.GetById(ctx, id)
	if err != nil {
		return types.Pet{}, err
	}

	return pet, nil
}
