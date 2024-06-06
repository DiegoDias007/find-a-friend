package services

import (
	"context"
	"find-a-friend/types"
	"testing"
)

func TestCreatePet(t *testing.T) {
	petService := CreatePetService()

	pet := types.CreatePet{
		Name:    "Milla",
		City:    "João Pessoa",
		Species: "Dog",
		Breed:   "Shitzu",
		Height:  0.3,
		Weight:  7.3,
		OrgId:   2,
	}

	createdPet, err := petService.Create(context.Background(), pet)
	if err != nil {
		t.Errorf("error creating pet: %v", err)
	}

	if createdPet.Name != "Milla" {
		t.Errorf("unexpected pet name.")
	}
}
