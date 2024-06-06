package repositories

import (
	"context"
	"find-a-friend/types"
)

type PetRepository interface {
	Create(ctx context.Context, pet types.CreatePet) (types.Pet, error)
	GetFromCity(ctx context.Context, city string) ([]types.Pet, error)
}