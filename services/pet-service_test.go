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
	petService := NewPetService(petRepo, orgRepo)

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
	petService := NewPetService(petRepo, orgRepo)

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
	require.Error(t, err, "expected an error of invalid org id when creating pet.")
}

func TestGetPetsFromCity(t *testing.T) {
	orgRepo := inMemory.NewOrgRepository()
	petRepo := inMemory.NewPetRepository()
	petService := NewPetService(petRepo, orgRepo)

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

	mockPet := types.CreatePet{
		Name:    "Ronaldo",
		City:    "Recife",
		Species: "Dog",
		Breed:   "Shitzu",
		Height:  0.3,
		Weight:  7.3,
		OrgId:   2,
	}

	_, err := petRepo.Create(ctx, pet)
	_, err = petRepo.Create(ctx, mockPet)
	pets, err := petService.GetFromCity(ctx, "João Pessoa")
	require.NoError(t, err, "error when getting pets from city")
	assert.Equal(t, pet.Name, pets[0].Name, "pet name should be the same")
	assert.Equal(t, 1, len(pets), "invalid length of pets")
}
