package services

import (
	"context"
	inMemory "find-a-friend/repositories/in-memory"
	"find-a-friend/types"
	"find-a-friend/utils"
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
		Email:    "welovepets@gmail.com",
		Password: "Random Password",
	}

	token, err := orgService.Create(ctx, org)
	require.NoError(t, err, "got an unexpected error when creating org.")
	_, isValidToken := utils.ValidateToken(token)
	require.True(t, isValidToken, "token was expected to be true.")
}

func TestLoginOrg(t *testing.T) {
	orgRepo := inMemory.NewOrgRepository()
	orgService := NewOrgService(orgRepo)

	ctx := context.Background()

	org := types.CreateOrg{
		Name:     "We love Pets",
		Address:  "John Doe, 500, João Pessoa",
		Whatsapp: "94832-1283",
		Email:    "welovepets@gmail.com",
		Password: "Random Password",
	}

	_, err := orgService.Create(context.Background(), org)
	require.NoError(t, err, "got an error when creating org.")
	
	_, err = orgService.Login(ctx, "welovepets@gmail.com", "Random Password")
	require.NoError(t, err, "got an error when logging in.")

	_, err = orgService.Login(ctx, "welove@gmail.com", "Wrong Password")
	require.Error(t, err, "got no error when logging in with invalid credentials.")

}

func TestGetOrgById(t *testing.T) {
	orgRepo := inMemory.NewOrgRepository()
	orgService := NewOrgService(orgRepo)

	ctx := context.Background()

	org := types.CreateOrg{
		Name:     "We love Pets",
		Address:  "John Doe, 500, João Pessoa",
		Whatsapp: "94832-1283",
		Email:    "welovepets@gmail.com",
		Password: "Random Password",
	}

	createOrg, err := orgRepo.Create(ctx, org)
	require.NoError(t, err, "got an unexpected error.")

	fetchedOrg, err := orgService.GetById(ctx, createOrg.Id)
	require.NoError(t, err, "got an error when getting org by id.")
	assert.Equal(t, fetchedOrg.Name, "We love Pets")
	assert.Equal(t, createOrg.Id, fetchedOrg.Id)
}
