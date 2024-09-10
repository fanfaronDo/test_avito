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

func (t *TenderRepo) CreateTender(tender domain.Tender) (domain.Tender, error) {
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
		tender.OrganizationID,
		tender.CreatorID,
		tender.Version).Scan(&uuid)

	if err != nil {
		log.Fatalf("%s: %v", ErrCreationTender, err)
		return domain.Tender{}, ErrCreationTender
	}
	tender.ID = uuid
	return tender, nil
}

func (t *TenderRepo) GetTenders(limit, offset int, serviceType string) ([]domain.Tender, error) {
	var tenders []domain.Tender

	query := `SELECT id, name, description, 
       				service_type, status, 
       				organization_id, creator_id, 
       				version, created_at, updated_at 
			  FROM tenders;`

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
		log.Fatalf("%s: %v", ErrTenderNotFound, err)
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
			&tender.OrganizationID,
			&tender.CreatorID,
			&tender.Version,
			&tender.CreatedAt,
			&tender.UpdatedAt)

		if err != nil {
			log.Fatalf("%s: %v", ErrScanDataTender, err)
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
       				organization_id, creator_id, 
       				version, created_at, updated_at 
				FROM tenders WHERE creator_id = $1 ORDER BY name DESC LIMIT $2 OFFSET $3;`

	ctx, cancelFn := context.WithTimeout(context.Background(), timeuotCtx)
	defer cancelFn()
	rows, err := t.db.QueryContext(ctx, query, uuid, limit, offset)
	if err != nil {
		log.Fatalf("%s: %v", ErrTenderNotFound, err)
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
			&tender.OrganizationID,
			&tender.CreatorID,
			&tender.Version,
			&tender.CreatedAt,
			&tender.UpdatedAt)

		if err != nil {
			log.Fatalf("%s: %v", ErrScanDataTender, err)
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
		log.Fatalf("%s: %v", ErrInFailedTransaction, err)
		return domain.Tender{}, ErrInFailedTransaction
	}
	ctx, cancelFn := context.WithTimeout(context.Background(), timeuotCtx)
	defer cancelFn()

	updateStatusQuery := `UPDATE tenders SET status = $1 WHERE creator_id = $2 AND id = $3;`
	_, err = tx.ExecContext(ctx, updateStatusQuery, status, userUUID, tenderUUID)
	if err != nil {
		tx.Rollback()
		log.Fatalf("%s: %v", ErrUpdatedTender, err)
		return domain.Tender{}, ErrUpdatedTender
	}

	getTenderByIdQuery := `SELECT id, name, description, 
        service_type, status, 
        organization_id, creator_id, 
        version, created_at, updated_at 
        FROM tenders WHERE id = $1;`

	err = tx.QueryRowContext(ctx, getTenderByIdQuery, tenderUUID).Scan(
		&tender.ID,
		&tender.Name,
		&tender.Description,
		&tender.ServiceType,
		&tender.Status,
		&tender.OrganizationID,
		&tender.CreatorID,
		&tender.Version,
		&tender.CreatedAt,
		&tender.UpdatedAt,
	)
	if err != nil {
		tx.Rollback()
		if err == sql.ErrNoRows {
			log.Fatalf("%s: %v", ErrTenderNotFound, err)
			return domain.Tender{}, ErrTenderNotFound
		}
		log.Fatalf("%s: %v", ErrFetchingTender, err)
		return domain.Tender{}, ErrFetchingTender
	}

	if err := tx.Commit(); err != nil {
		log.Fatalf("%s: %v", ErrCommittingTransaction, err)
		return domain.Tender{}, ErrCommittingTransaction
	}

	return tender, nil
}

//func (t *TenderRepo) UpdateStatusTenderById(tenderUUID, status, userUUID string) (domain.Tender, error) {
//	var tender domain.Tender
//	tx, err := t.db.Begin()
//	if err != nil {
//		log.Fatalf("%s: %v", ErrInFailedTransaction, err)
//		return domain.Tender{}, ErrInFailedTransaction
//	}
//	ctx, cancelFn := context.WithTimeout(context.Background(), timeuotCtx)
//	defer cancelFn()
//
//	updateStatusQuery := `UPDATE tenders SET status = $1 WHERE creator_id = $2 AND id = $3;`
//	err = tx.QueryRowContext(ctx, updateStatusQuery, status, userUUID, tenderUUID).Err()
//	if err != nil {
//		tx.Rollback()
//		log.Fatalf("%s: %v", ErrUpdatedTender, err)
//		return domain.Tender{}, ErrUpdatedTender
//	}
//
//	getTenderByIdQuery := `SELECT id, name, description,
//       				service_type, status,
//       				organization_id, creator_id,
//       				version, created_at, updated_at
//				FROM tenders WHERE id = $1;`
//
//	row := t.db.QueryRowContext(ctx, getTenderByIdQuery, tenderUUID)
//	err = row.Scan(&tender.ID,
//		&tender.Name,
//		&tender.Description,
//		&tender.ServiceType,
//		&tender.Status,
//		&tender.OrganizationID,
//		&tender.CreatorID,
//		&tender.Version,
//		&tender.CreatedAt,
//		&tender.UpdatedAt)
//
//	if err != nil {
//		tx.Rollback()
//		log.Printf("%s: %v", ErrScanDataTender, err)
//		return domain.Tender{}, ErrScanDataTender
//	}
//
//	tx.Commit()
//
//	return tender, nil
//}
