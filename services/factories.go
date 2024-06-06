package services

import (
	"find-a-friend/repositories/in-memory"
)

func CreatePetService() *PetService {
	petRepo := inMemory.NewPetRepository()
	return NewPetService(petRepo)
}