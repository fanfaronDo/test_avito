package repo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/fanfaronDo/test_avito/internal/domain"
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
		return domain.Tender{}, fmt.Errorf("Error %s when inserting row into tenders table", err)
	}
	tender.ID = uuid
	return tender, nil
}

func (t *TenderRepo) GetTenders(limit, offset int, serviceType string) ([]domain.Tender, error) {
	var tenders []domain.Tender

	query := "SELECT id, name, description, service_type, status, organization_id, creator_id, version, created_at, updated_at FROM tenders"
	params := []interface{}{}
	if serviceType != "" {
		query += " WHERE service_type = $1"
		params = append(params, serviceType)
	}
	if limit != 0 && offset != 0 {
		query += " LIMIT $2 OFFSET $3"
		params = append(params, limit, offset)
	}

	ctx, cancelFn := context.WithTimeout(context.Background(), timeuotCtx)
	defer cancelFn()
	rows, err := t.db.QueryContext(ctx, query, params...)
	if err != nil {
		return []domain.Tender{}, fmt.Errorf("Error %s when getting rows from tenders table", err)
	}

	defer rows.Close()
	for rows.Next() {
		var tender domain.Tender
		err := rows.Scan(&tender.ID,
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
			return []domain.Tender{}, fmt.Errorf("Error %s when getting rows from tenders table", err)
		}
		tenders = append(tenders, tender)
	}
	return tenders, nil
}
