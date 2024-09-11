package repo

import (
	"database/sql"
	"github.com/fanfaronDo/test_avito/internal/domain"
)

type Auth interface {
	CheckUserCharge(userid, organisationid string) (string, error)
	CreateUserCharge(userid, username string) (string, error)
	GetUserUUID(username string) (string, error)
	GetUserChargeUUID(username string) (string, error)
}

type Tender interface {
	CreateTender(tender domain.Tender, orgID, creatorID string) (domain.Tender, error)
	GetTenders(limit, offset int, serviceType string) ([]domain.Tender, error)
	GetTendersByUserID(limit, offset int, uuid string) ([]domain.Tender, error)
	GetStatusTenderById(tenderUUID, userUUID string) (string, error)
	UpdateStatusTenderById(tenderUUID, status, userUUID string) (domain.Tender, error)
	UpdateTender(tenderUUID, userUUID string, tenderEditor *domain.TenderEditor) (domain.Tender, error)
	GetTenderById(tenderUUID string) (domain.Tender, error)
	//RollbackTender(tenderUUID, userUUID string, version int) (domain.Tender, error)
}

type Repository struct {
	Auth
	Tender
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Tender: NewTenderRepo(db),
		Auth:   NewAuthRepo(db),
	}
}
