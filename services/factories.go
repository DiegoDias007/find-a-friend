package services

import (
	"find-a-friend/repositories/psql"

	"github.com/jackc/pgx/v5"
)

func MakePetService(db *pgx.Conn) *PetService {
	petRepo := psql.NewPetRepository(db)
	orgRepo := psql.NewOrgRepository(db)
	return NewPetService(petRepo, orgRepo)
}