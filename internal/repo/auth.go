package repo

import (
	"context"
	"database/sql"
	log "github.com/sirupsen/logrus"
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
		log.Fatalf("%s: %v", ErrUserChargeNotFound, err)
		return "", ErrUserChargeNotFound
	}

	return uuid, nil
}

func (a *AuthRepo) GetUserUUID(username string) (string, error) {
	var uuid string
	ctx, cancelFn := context.WithTimeout(context.Background(), timeuotCtx)
	defer cancelFn()
	query := `SELECT id FROM employee WHERE username = $1;`
	err := a.db.QueryRowContext(ctx, query, username).Scan(&uuid)
	if err != nil {
		log.Fatalf("%s: %v", ErrUserNotFound, err)
		return "", ErrUserNotFound
	}
	return uuid, nil
}
