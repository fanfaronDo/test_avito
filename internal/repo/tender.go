package repo

import (
	"context"
	"database/sql"
	"github.com/fanfaronDo/test_avito/internal/domain"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

const (
	timeuotCtx = 5 * time.Second
)

type TenderRepo struct {
	db *sql.DB
}

func NewTenderRepo(db *sql.DB) *TenderRepo {
	return &TenderRepo{db: db}
}

func (t *TenderRepo) CreateTender(tender domain.Tender, orgID, userUUID string) (domain.Tender, error) {
	query := `INSERT INTO tenders (name, description, service_type, status, organization_id, creator_id, version) 
              VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	ctx, cancelFn := context.WithTimeout(context.Background(), timeuotCtx)
	defer cancelFn()
	var uuid string
	err := t.db.QueryRowContext(ctx, query,
		tender.Name,
		tender.Description,
		tender.ServiceType,
		tender.Status,
		orgID,
		userUUID,
		tender.Version).Scan(&uuid)

	if err != nil {
		log.Debugf("%s: %v", ErrCreationTender, err)
		return domain.Tender{}, ErrCreationTender
	}
	tender.ID = uuid
	return tender, nil
}

func (t *TenderRepo) GetTenderById(tenderUUID string) (domain.Tender, error) {
	var tender domain.Tender
	query := `SELECT id, name, description, 
       				service_type, status,  
       				version, created_at 
			  FROM tenders WHERE id = $1;`

	ctx, cancelFn := context.WithTimeout(context.Background(), timeuotCtx)
	defer cancelFn()
	err := t.db.QueryRowContext(ctx, query, tenderUUID).Scan(&tender.ID,
		&tender.Name,
		&tender.Description,
		&tender.ServiceType,
		&tender.Status,
		&tender.Version,
		&tender.CreatedAt)
	if err != nil {
		return domain.Tender{}, ErrTenderNotFound
	}

	return tender, nil
}

func (t *TenderRepo) GetTenders(limit, offset int, serviceType string) ([]domain.Tender, error) {
	var tenders []domain.Tender

	query := `SELECT id, name, description, 
       				service_type, status, 
       				version, created_at 
			  FROM tenders`

	params := []interface{}{}
	paramsIndex := 1
	if serviceType != "" {
		query += " WHERE service_type = $" + strconv.Itoa(paramsIndex)
		params = append(params, serviceType)
		paramsIndex++
	}
	query += " ORDER BY name DESC"
	query += " LIMIT $" + strconv.Itoa(paramsIndex) + " OFFSET $" + strconv.Itoa(paramsIndex+1)
	params = append(params, limit, offset)

	ctx, cancelFn := context.WithTimeout(context.Background(), timeuotCtx)
	defer cancelFn()
	rows, err := t.db.QueryContext(ctx, query, params...)
	if err != nil {
		log.Debugf("%s: %v", ErrTenderNotFound, err)
		return []domain.Tender{}, ErrTenderNotFound
	}

	defer rows.Close()
	for rows.Next() {
		var tender domain.Tender
		err = rows.Scan(&tender.ID,
			&tender.Name,
			&tender.Description,
			&tender.ServiceType,
			&tender.Status,
			&tender.Version,
			&tender.CreatedAt)

		if err != nil {
			log.Debugf("%s: %v", ErrScanDataTender, err)
			return []domain.Tender{}, ErrScanDataTender
		}
		tenders = append(tenders, tender)
	}
	return tenders, nil
}

func (t *TenderRepo) GetTendersByUserID(limit, offset int, uuid string) ([]domain.Tender, error) {
	var tenders []domain.Tender
	query := `SELECT id, name, description, 
       				service_type, status, 
       				version, created_at 
				FROM tenders WHERE creator_id = $1 ORDER BY name DESC LIMIT $2 OFFSET $3;`

	ctx, cancelFn := context.WithTimeout(context.Background(), timeuotCtx)
	defer cancelFn()
	rows, err := t.db.QueryContext(ctx, query, uuid, limit, offset)
	if err != nil {
		log.Debugf("%s: %v", ErrTenderNotFound, err)
		return []domain.Tender{}, ErrTenderNotFound
	}
	defer rows.Close()

	for rows.Next() {
		var tender domain.Tender
		err = rows.Scan(&tender.ID,
			&tender.Name,
			&tender.Description,
			&tender.ServiceType,
			&tender.Status,
			&tender.Version,
			&tender.CreatedAt)

		if err != nil {
			log.Debugf("%s: %v", ErrScanDataTender, err)
			return []domain.Tender{}, ErrScanDataTender
		}
		tenders = append(tenders, tender)
	}

	return tenders, nil
}

func (t *TenderRepo) GetStatusTenderById(tenderUUID, userUUID string) (string, error) {
	var status string

	query := `SELECT status 
				FROM tenders WHERE creator_id = $1 AND id = $2;`

	ctx, cancelFn := context.WithTimeout(context.Background(), timeuotCtx)
	defer cancelFn()

	err := t.db.QueryRowContext(ctx, query, userUUID, tenderUUID).Scan(&status)
	if err != nil {
		log.Printf("%s: %v", ErrTenderStatusNotFound, err)
		return "", ErrTenderStatusNotFound
	}

	return status, nil
}

func (t *TenderRepo) UpdateStatusTenderById(tenderUUID, status, userUUID string) (domain.Tender, error) {
	var tender domain.Tender
	tx, err := t.db.Begin()
	if err != nil {
		log.Debugf("%s: %v", ErrInFailedTransaction, err)
		return domain.Tender{}, ErrInFailedTransaction
	}
	ctx, cancelFn := context.WithTimeout(context.Background(), timeuotCtx)
	defer cancelFn()

	updateStatusQuery := `UPDATE tenders SET status = $1, 
                   version = COALESCE((SELECT MAX(version) FROM tenders_history WHERE tender_id = $3), 0) + 1 
               WHERE creator_id = $2 AND id = $3;`

	_, err = tx.ExecContext(ctx, updateStatusQuery, status, userUUID, tenderUUID)
	if err != nil {
		tx.Rollback()
		log.Debugf("%s: %v", ErrUpdatedTender, err)
		return domain.Tender{}, ErrUpdatedTender
	}

	getTenderByIdQuery := `SELECT id, name, description, 
        service_type, status,  
        version, created_at 
        FROM tenders WHERE id = $1;`

	err = tx.QueryRowContext(ctx, getTenderByIdQuery, tenderUUID).Scan(
		&tender.ID,
		&tender.Name,
		&tender.Description,
		&tender.ServiceType,
		&tender.Status,
		&tender.Version,
		&tender.CreatedAt,
	)
	if err != nil {
		tx.Rollback()
		if err == sql.ErrNoRows {
			log.Debugf("%s: %v", ErrTenderNotFound, err)
			return domain.Tender{}, ErrTenderNotFound
		}
		log.Debugf("%s: %v", ErrFetchingTender, err)
		return domain.Tender{}, ErrFetchingTender
	}

	if err := tx.Commit(); err != nil {
		log.Debugf("%s: %v", ErrCommittingTransaction, err)
		return domain.Tender{}, ErrCommittingTransaction
	}

	return tender, nil
}

func (t *TenderRepo) UpdateTender(tenderUUID, userUUID string, tenderEditor *domain.TenderEditor) (domain.Tender, error) {
	var tender domain.Tender
	tx, err := t.db.Begin()
	if err != nil {
		log.Debugf("%s: %v", ErrInFailedTransaction, err)
		return domain.Tender{}, ErrInFailedTransaction
	}
	ctx, cancelFn := context.WithTimeout(context.Background(), timeuotCtx)
	defer cancelFn()

	updateStatusQuery := `UPDATE tenders SET name = $1, description = $2, service_type = $3, version = COALESCE((SELECT MAX(version) FROM tenders_history WHERE tender_id = $4), 0) + 1 
               WHERE creator_id = $5 AND id = $6;`
	_, err = tx.ExecContext(ctx,
		updateStatusQuery, tenderEditor.Name,
		tenderEditor.Description,
		tenderEditor.ServiceType,
		tenderUUID,
		userUUID, tenderUUID)

	if err != nil {
		tx.Rollback()
		log.Debugf("%s: %v", ErrTenderNotFound, err)
		return domain.Tender{}, ErrTenderNotFound
	}

	getTenderByIdQuery := `SELECT id, name, description, 
        service_type, status, 
        version, created_at, updated_at 
        FROM tenders WHERE id = $1;`

	err = tx.QueryRowContext(ctx, getTenderByIdQuery, tenderUUID).Scan(
		&tender.ID,
		&tender.Name,
		&tender.Description,
		&tender.ServiceType,
		&tender.Status,
		&tender.Version,
		&tender.CreatedAt,
	)
	if err != nil {
		tx.Rollback()
		if err == sql.ErrNoRows {
			log.Debugf("%s: %v", ErrTenderNotFound, err)
			return domain.Tender{}, ErrTenderNotFound
		}
		log.Debugf("%s: %v", ErrFetchingTender, err)
		return domain.Tender{}, ErrFetchingTender
	}

	if err := tx.Commit(); err != nil {
		log.Debugf("%s: %v", ErrCommittingTransaction, err)
		return domain.Tender{}, ErrCommittingTransaction
	}

	return tender, nil
}

func (r *TenderRepo) RollbackTender(tenderUUID, userUUID string, version int) (domain.Tender, error) {
	var tender domain.Tender
	tx, err := r.db.Begin()
	if err != nil {
		log.Debugf("%s: %v", ErrInFailedTransaction, err)
		return domain.Tender{}, ErrInFailedTransaction
	}
	ctx, cancelFn := context.WithTimeout(context.Background(), timeuotCtx)
	defer cancelFn()
	query := `UPDATE tenders t SET name = th.name,
					description = th.description,
					status = th.status,
					service_type = th.service_type
			  FROM (SELECT tender_id, name, description, status, service_type
				    FROM tenders_history
				    WHERE tender_id = $1 AND version = $2) th
			  WHERE t.id = th.tender_id;`

	_, err = tx.ExecContext(ctx, query, tenderUUID, version)
	if err != nil {
		tx.Rollback()
		log.Debugf("%s: %v", ErrTenderNotFound, err)
		return domain.Tender{}, ErrTenderNotFound
	}
	getTenderByIdQuery := `SELECT id, name, description, 
								service_type, status, 
						        version, created_at, updated_at 
						   FROM tenders WHERE id = $1;`

	err = tx.QueryRowContext(ctx, getTenderByIdQuery, tenderUUID).Scan(
		&tender.ID,
		&tender.Name,
		&tender.Description,
		&tender.ServiceType,
		&tender.Status,
		&tender.Version,
		&tender.CreatedAt,
	)
	if err != nil {
		tx.Rollback()
		if err == sql.ErrNoRows {
			log.Debugf("%s: %v", ErrTenderNotFound, err)
			return domain.Tender{}, ErrTenderNotFound
		}
		log.Debugf("%s: %v", ErrFetchingTender, err)
		return domain.Tender{}, ErrFetchingTender
	}

	if err = tx.Commit(); err != nil {
		log.Debugf("%s: %v", ErrCommittingTransaction, err)
		return domain.Tender{}, ErrCommittingTransaction
	}

	return tender, nil
}
