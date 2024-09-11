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

func (a *AuthRepo) FindUserUUIDCharge(userid string) (string, error) {
	query := `SELECT id FROM organization_responsible WHERE user_id = $1;`

	var uuid string
	ctx, cancelFn := context.WithTimeout(context.Background(), timeuotCtx)
	defer cancelFn()
	err := a.db.QueryRowContext(ctx, query, userid).Scan(&uuid)
	if err != nil {
		log.Debugf("%s: %v", ErrUserChargeNotFound, err)
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
		log.Debugf("%s: %v", ErrUserNotFound, err)
		return "", ErrUserNotFound
	}
	return uuid, nil
}
