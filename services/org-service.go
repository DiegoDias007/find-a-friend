package services

import (
	"context"
	"find-a-friend/repositories"
	"find-a-friend/types"
	"find-a-friend/utils"
	"fmt"
)

type OrgService struct {
	orgRepo repositories.OrgRepository
}

func NewOrgService(orgRepo repositories.OrgRepository) *OrgService {
	return &OrgService{orgRepo: orgRepo}
}

func (s *OrgService) Create(ctx context.Context, org types.CreateOrg) (string, error) {
	hashedPassword, err := utils.HashPassword(org.Password, 10)
	if err != nil {
		return "", err
	}

	org.Password = hashedPassword

	newOrg, err := s.orgRepo.Create(ctx, org)
	if err != nil {
		return "", err
	}

	token, err := utils.CreateToken(newOrg.Name, newOrg.Id)
	if err != nil {
		return "", fmt.Errorf("error creating token.")
	}

	return token, nil
}

func (s *OrgService) Login(ctx context.Context, email, password string) (string, error) {
	org, err := s.orgRepo.GetByEmail(ctx, email)
	if err != nil {
		return "", fmt.Errorf("invalid credentials.")
	}

	isValidPassword := utils.CheckHashedPassword(org.Password, password)
	if !isValidPassword {
		return "", fmt.Errorf("invalid credentials.")
	}

	token, err := utils.CreateToken(org.Name, org.Id)
	if err != nil {
		return "", err
	}

	return token, nil
}


func (s *OrgService) GetById(ctx context.Context, id int) (types.Org, error) {
	org, err := s.orgRepo.GetById(ctx, id)
	if err != nil {
		return types.Org{}, err
	}

	return org, nil
}