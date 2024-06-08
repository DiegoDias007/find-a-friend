package services

import (
	"find-a-friend/repositories"
)

func CreatePetService(petRepo repositories.PetRepository, orgRepo repositories.OrgRepository) *PetService {
	return NewPetService(petRepo, orgRepo)
}