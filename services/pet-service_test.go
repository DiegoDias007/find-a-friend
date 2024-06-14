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
		Email:    "welovepets@gmail.com",
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

	filter := types.PetFilter{
		City: "João Pessoa",
	}

	pets, err := petService.GetFromCity(ctx, filter)
	require.NoError(t, err, "error when getting pets from city")
	assert.Equal(t, pet.Name, pets[0].Name, "pet name should be the same")
	assert.Equal(t, 1, len(pets), "invalid length of pets")
}

func TestGetPetsUsingFiltering(t *testing.T) {
	orgRepo := inMemory.NewOrgRepository()
	petRepo := inMemory.NewPetRepository()
	petService := NewPetService(petRepo, orgRepo)

	ctx := context.Background()

	org := types.CreateOrg{
		Name:     "We love Pets",
		Address:  "John Doe, 500, João Pessoa",
		Whatsapp: "94832-1283",
		Email:    "welovepets@gmail.com",
		Password: "Random Password",
	}

	createdOrg, err := petService.orgRepo.Create(ctx, org)


	pet1 := types.CreatePet{
		Name:    "Milla",
		City:    "João Pessoa",
		Species: "Dog",
		Breed:   "Shitzu",
		Height:  0.3,
		Weight:  7.3,
		OrgId:   createdOrg.Id,
	}

	pet2 := types.CreatePet{
		Name:    "Ronaldo",
		City:    "João Pessoa",
		Species: "Dog",
		Breed:   "Bulldog",
		Height:  0.3,
		Weight:  8.5,
		OrgId:   createdOrg.Id,
	}

	pet3 := types.CreatePet{
		Name: "Messi",
		City: "João Pessoa",
		Species: "Cat",
		Breed: "Idk",
		Height: 1,
		Weight: 5.5,
		OrgId: createdOrg.Id,
	}

	_, err = petService.Create(ctx, pet1)
	require.NoError(t, err, "got error.")
	_, err = petService.Create(ctx, pet2)
	require.NoError(t, err, "got error.")
	_, err = petService.Create(ctx, pet3)
	require.NoError(t, err, "got error.")

	weight := 7.0

	filter := types.PetFilter{
		City: "João Pessoa",
		Weight: &weight,
	}

	pets, err := petService.GetFromCity(ctx, filter)
	require.NoError(t, err, "got an unexpected error.")
	assert.Equal(t, 2, len(pets), "they should be equal.")
}
