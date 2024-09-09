package repo

import (
	"database/sql"
	"github.com/fanfaronDo/test_avito/internal/domain"
)

type Tender interface {
	CreateTender(tender domain.Tender) (domain.Tender, error)
	GetUserUUIDCharge(username, organisation_id string) (string, error)
}

type Repository struct {
	Tender
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		Tender: NewTenderRepo(db),
	}
}
