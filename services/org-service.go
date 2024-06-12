package services

import (
	"context"
	"find-a-friend/repositories"
	"find-a-friend/types"
	"find-a-friend/utils"
)

type OrgService struct {
	orgRepo repositories.OrgRepository
}

func NewOrgService(orgRepo repositories.OrgRepository) *OrgService {
	return &OrgService{orgRepo: orgRepo}
}

func (r *OrgService) Create(ctx context.Context, org types.CreateOrg) (types.Org, error) {
	hashedPassword, err := utils.HashPassword(org.Password, 10)
	if err != nil {
		return types.Org{}, err
	}

	org.Password = hashedPassword

	newOrg, err := r.orgRepo.Create(ctx, org)
	if err != nil {
		return types.Org{}, err
	}

	return newOrg, nil
}


func (r *OrgService) GetById(ctx context.Context, id int) (types.Org, error) {
	org, err := r.orgRepo.GetById(ctx, id)
	if err != nil {
		return types.Org{}, err
	}

	return org, nil
}