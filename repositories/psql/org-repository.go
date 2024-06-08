package psql

import (
	"context"
	"find-a-friend/types"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type OrgRepository struct {
	db *pgx.Conn
}

func NewOrgRepository(db *pgx.Conn) *OrgRepository {
	return &OrgRepository{db: db}
}

func (r *OrgRepository) Create(ctx context.Context, org types.CreateOrg) (types.Org, error) {
	var createdOrg types.Org
	query := `
		INSERT INTO org (name, address, whatsapp, password)
		VALUES ($1, $2, $3, $4)
		RETURNING id, name, address, whatsapp, password
	`

	err := r.db.QueryRow(ctx, query, org.Name, org.Address, org.Whatsapp, org.Password).
				Scan(&createdOrg.Id, &createdOrg.Name, &createdOrg.Address, &org.Whatsapp, &org.Password)
	if err != nil {
		return types.Org{}, fmt.Errorf("error when creating org: %v", err)
	}

	return createdOrg, nil
}

func (r *OrgRepository) GetById(ctx context.Context, id int) (types.Org, error) {
	var org types.Org
	
	query := `
		SELECT * FROM org WHERE id = $1
		RETURNING id, name, address, whatsapp, password
	`

	err := r.db.QueryRow(ctx, query, id).Scan(&org.Id, &org.Name, &org.Address, &org.Whatsapp, &org.Password)
	if err != nil {
		return types.Org{}, fmt.Errorf("error when getting org by id: %v", err)
	}

	return org, nil
}