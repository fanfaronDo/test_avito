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

func (a *AuthRepo) CreateUserCharge(userid, username string) (string, error) {
	var createdUUID string
	query := `INSERT INTO user_charges(user_id, username) VALUES ($1, $2) RETURNING id;`
	ctx, cancelFn := context.WithTimeout(context.Background(), timeuotCtx)
	defer cancelFn()
	err := a.db.QueryRowContext(ctx, query, userid, username).Scan(&createdUUID)
	if err != nil {
		return "", err
	}

	return createdUUID, nil
}

func (a *AuthRepo) CheckUserCharge(userid, organisationid string) (string, error) {
	query := `SELECT user_id FROM organization_responsible WHERE user_id = $1 AND organization_id = $2;`

	var uuid string
	ctx, cancelFn := context.WithTimeout(context.Background(), timeuotCtx)
	defer cancelFn()
	err := a.db.QueryRowContext(ctx, query, userid, organisationid).Scan(&uuid)
	if err != nil {
		log.Debugf("%s: %v", ErrUserChargeNotFound, err)
		return "", ErrUserChargeNotFound
	}

	return uuid, nil
}

func (a *AuthRepo) GetUserChargeUUID(username string) (string, error) {
	var uuid string
	ctx, cancelFn := context.WithTimeout(context.Background(), timeuotCtx)
	defer cancelFn()
	query := `SELECT id FROM user_charges WHERE username = $1;`
	err := a.db.QueryRowContext(ctx, query, username).Scan(&uuid)
	if err != nil {
		log.Debugf("%s: %v", ErrUserNotFound, err)
		return "", ErrUserNotFound
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
