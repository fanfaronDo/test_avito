package repo

import (
	"database/sql"
	"github.com/fanfaronDo/test_avito/internal/domain"
)

type Auth interface {
	GetUserUUIDCharge(username, organisation_id string) (string, error)
}

type Tender interface {
	CreateTender(tender domain.Tender) (domain.Tender, error)
}

type Repository struct {
	Auth
	Tender
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		Tender: NewTenderRepo(db),
		Auth:   NewAuthRepo(db),
	}
}
