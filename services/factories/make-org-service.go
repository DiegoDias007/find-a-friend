package factories

import (
	"find-a-friend/repositories/psql"
	"find-a-friend/services"

	"github.com/jackc/pgx/v5"
)

func MakeOrgService(db *pgx.Conn) *services.OrgService {
	orgRepo := psql.NewOrgRepository(db)
	service := services.NewOrgService(orgRepo)
	return service
}