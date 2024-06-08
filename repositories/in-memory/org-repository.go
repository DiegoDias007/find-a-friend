package inMemory

import (
	"context"
	"find-a-friend/types"
	"find-a-friend/utils"
	"fmt"
)

type OrgRepository struct {
	orgs []types.Org
}

func NewOrgRepository() *OrgRepository {
	return &OrgRepository{}
}

func (r *OrgRepository) Create(ctx context.Context, org types.CreateOrg) (types.Org, error) {
	randomId := utils.GenerateRandomNumber(1000)

	newOrg := types.Org{
		Id: randomId,
		Name: org.Name,
		Address: org.Address,
		Whatsapp: org.Whatsapp,
		Password: org.Password,
	}

	r.orgs = append(r.orgs, newOrg)

	return newOrg, nil
}

func (r *OrgRepository) GetById(ctx context.Context, id int) (types.Org, error) {
	for _, org := range(r.orgs) {
		if org.Id == id {
			return org, nil
		}
	}

	return types.Org{}, fmt.Errorf("org not found.")
}