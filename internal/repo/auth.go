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

func (a *AuthRepo) CheckOrganizationAffiliation(userid, organisationid string) (string, error) {
	query := `SELECT user_id FROM organization_responsible WHERE user_id = $1 AND organization_id = $2;`

	var uuid string
	ctx, cancelFn := context.WithTimeout(context.Background(), timeuotCtx)
	defer cancelFn()
	err := a.db.QueryRowContext(ctx, query, userid, organisationid).Scan(&uuid)
	if err != nil {
		log.Debugf("%s: %v", ErrUserChargeNotAffiliationThisOrganisation, err)
		return "", ErrUserChargeNotAffiliationThisOrganisation
	}

	return uuid, nil
}

func (a *AuthRepo) CheckUserChargeAffiliation(userUUID string) (string, error) {
	var uuid string
	ctx, cancelFn := context.WithTimeout(context.Background(), timeuotCtx)
	defer cancelFn()
	query := `SELECT user_id FROM organization_responsible WHERE user_id = $1;`
	err := a.db.QueryRowContext(ctx, query, userUUID).Scan(&uuid)
	if err != nil {
		log.Debugf("%s: %v", ErrUserChargeNotAffiliationThisTender, err)
		return "", ErrUserChargeNotAffiliationThisTender
	}

	return uuid, nil
}

func (a *AuthRepo) CheckUserCreatorTender(userUUID, tenderUUID string) (string, error) {
	var uuid string
	ctx, cancelFn := context.WithTimeout(context.Background(), timeuotCtx)
	defer cancelFn()
	query := `SELECT creator_id FROM tenders WHERE creator_id = $1 AND id = $2;`
	err := a.db.QueryRowContext(ctx, query, userUUID, tenderUUID).Scan(&uuid)
	if err != nil {
		log.Debugf("%s: %v", ErrUserNotCreator, err)
		return "", ErrUserNotCreator
	}

	return uuid, nil
}

func (a *AuthRepo) CheckUserCreatorBids(userUUID, bidsUUID string) (string, error) {
	var uuid string
	ctx, cancelFn := context.WithTimeout(context.Background(), timeuotCtx)
	defer cancelFn()
	query := `SELECT author_id FROM bids WHERE author_id = $1 AND id = $2;`
	err := a.db.QueryRowContext(ctx, query, userUUID, bidsUUID).Scan(&uuid)
	if err != nil {
		log.Debugf("%s: %v", ErrUserNotCreator, err)
		return "", ErrUserNotFound
	}

	return uuid, nil
}

func (a *AuthRepo) CheckUserCreatorBidsByTenderId(userUUID, tenderUUID string) (string, error) {
	var uuid string
	ctx, cancelFn := context.WithTimeout(context.Background(), timeuotCtx)
	defer cancelFn()
	query := `SELECT author_id FROM bids WHERE author_id = $1 AND tender_id = $2;`
	err := a.db.QueryRowContext(ctx, query, userUUID, tenderUUID).Scan(&uuid)
	if err != nil {
		log.Debugf("%s: %v", ErrUserNotCreator, err)
		return "", ErrUserNotFound
	}

	return uuid, nil
}

func (a *AuthRepo) CheckUserID(userUUID string) (string, error) {
	var uuid string
	ctx, cancelFn := context.WithTimeout(context.Background(), timeuotCtx)
	defer cancelFn()
	query := `SELECT id FROM employee WHERE id = $1;`
	err := a.db.QueryRowContext(ctx, query, userUUID).Scan(&uuid)
	if err != nil {
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
