package repositories

import (
	"context"
	"find-a-friend/types"
)

type PetRepository interface {
	Create(ctx context.Context, pet types.CreatePet) (types.Pet, error)
	GetFromCity(ctx context.Context, filter types.PetFilter) ([]types.Pet, error)
	GetById(ctx context.Context, id int) (types.Pet, error)
}

type OrgRepository interface {
	Create(ctx context.Context, org types.CreateOrg) (types.Org, error)
	GetByEmail(ctx context.Context, email string) (types.Org, error)
	GetById(ctx context.Context, id int) (types.Org, error)
}
