package services

import (
	"context"
	inMemory "find-a-friend/repositories/in-memory"
	"find-a-friend/types"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreatePet(t *testing.T) {
	orgRepo := inMemory.NewOrgRepository()
	petRepo := inMemory.NewPetRepository()
	petService := CreatePetService(petRepo, orgRepo)
	
	ctx := context.Background()

	org := types.CreateOrg{
		Name:     "We love Pets",
		Address:  "John Doe, 500, João Pessoa",
		Whatsapp: "94832-1283",
		Password: "Random Password",
	}


	createdOrg, err := petService.orgRepo.Create(ctx, org)

	pet := types.CreatePet{
		Name:    "Milla",
		City:    "João Pessoa",
		Species: "Dog",
		Breed:   "Shitzu",
		Height:  0.3,
		Weight:  7.3,
		OrgId:   createdOrg.Id,
	}

	_, err = petService.Create(ctx, pet)
	require.NoError(t, err, "error when creating pet.")
	assert.Equal(t, "Milla", pet.Name, "they should be equal.")
}


func TestCreatePetWithInvalidOrgId(t *testing.T) {
	orgRepo := inMemory.NewOrgRepository()
	petRepo := inMemory.NewPetRepository()
	petService := CreatePetService(petRepo, orgRepo)

	ctx := context.Background()

	pet := types.CreatePet{
		Name:    "Milla",
		City:    "João Pessoa",
		Species: "Dog",
		Breed:   "Shitzu",
		Height:  0.3,
		Weight:  7.3,
		OrgId:   2,
	}

	_, err := petService.Create(ctx, pet)
	require.Error(t, err, "expected an error when of invalid org id when creating pet.")
}