package factories

import (
	"find-a-friend/repositories/psql"
	"find-a-friend/services"

	"github.com/jackc/pgx/v5"
)

func MakePetService(db *pgx.Conn) *services.PetService {
	petRepo := psql.NewPetRepository(db)
	orgRepo := psql.NewOrgRepository(db)
	return services.NewPetService(petRepo, orgRepo)
}