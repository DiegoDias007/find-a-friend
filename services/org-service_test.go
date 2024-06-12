package services

import (
	"context"
	inMemory "find-a-friend/repositories/in-memory"
	"find-a-friend/types"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateOrg(t *testing.T) {
	orgRepo := inMemory.NewOrgRepository()
	orgService := NewOrgService(orgRepo)

	ctx := context.Background()

	org := types.CreateOrg{
		Name:     "We love Pets",
		Address:  "John Doe, 500, João Pessoa",
		Whatsapp: "94832-1283",
		Password: "Random Password",
	}

	createOrg, err := orgService.Create(ctx, org)
	require.NoError(t, err, "got an unexpected error when creating org.")
	assert.Equal(t, createOrg.Name, "We love Pets", "expected them to be equal")
}

func TestGetOrgById(t *testing.T) {
	orgRepo := inMemory.NewOrgRepository()
	orgService := NewOrgService(orgRepo)

	ctx := context.Background()
	
	org := types.CreateOrg{
		Name:     "We love Pets",
		Address:  "John Doe, 500, João Pessoa",
		Whatsapp: "94832-1283",
		Password: "Random Password",
	}

	createOrg, err := orgRepo.Create(ctx, org)
	require.NoError(t, err, "got an unexpected error.")

	fetchedOrg, err := orgService.GetById(ctx, createOrg.Id)
	require.NoError(t, err, "got an error when getting org by id.")
	assert.Equal(t, fetchedOrg.Name, "We love Pets")
	assert.Equal(t, createOrg.Id, fetchedOrg.Id)
}
