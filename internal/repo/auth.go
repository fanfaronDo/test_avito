package repo

import (
	"context"
	"database/sql"
	"fmt"
)

type AuthRepo struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) *AuthRepo {
	return &AuthRepo{db: db}
}

func (a *AuthRepo) GetUserUUIDCharge(username, organisation_id string) (string, error) {
	query := `SELECT e.id FROM organization_responsible o 
				LEFT JOIN employee e ON o.user_id = e.id 
				WHERE e.username = $1 AND o.organization_id = $2;`

	var uuid string
	ctx, cancelFn := context.WithTimeout(context.Background(), timeuotCtx)
	defer cancelFn()
	err := a.db.QueryRowContext(ctx, query, username, organisation_id).Scan(&uuid)
	if err != nil {
		return "", fmt.Errorf("Error %s when checking user charge", err)
	}

	return uuid, nil
}
