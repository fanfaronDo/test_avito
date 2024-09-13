package repo

import (
	"database/sql"
	"github.com/fanfaronDo/test_avito/internal/domain"
)

type Auth interface {
	CheckOrganizationAffiliation(userid, organisationid string) (string, error)
	CheckUserChargeAffiliation(userUUID string) (string, error)
	CheckUserCreatorBids(userUUID, bidsUUID string) (string, error)
	CheckUserID(userUUID string) (string, error)
	CheckUserCreatorTender(userUUID, tenderUUID string) (string, error)
	GetUserUUID(username string) (string, error)
}

type Tender interface {
	CreateTender(tender domain.Tender, orgID, creatorID string) (domain.Tender, error)
	GetTenders(limit, offset int, serviceType string) ([]domain.Tender, error)
	GetTendersByUserID(limit, offset int, uuid string) ([]domain.Tender, error)
	GetStatusTenderById(tenderUUID, userUUID string) (string, error)
	UpdateStatusTenderById(tenderUUID, status, userUUID string) (domain.Tender, error)
	UpdateTender(tenderUUID, userUUID string, tenderEditor *domain.TenderEditor) (domain.Tender, error)
	GetTenderById(tenderUUID string) (domain.Tender, error)
	RollbackTender(tenderUUID, userUUID string, version int) (domain.Tender, error)
}

type Bid interface {
	CreateBids(tenderUUID, descr string, bid domain.Bid) (domain.Bid, error)
}

type Repository struct {
	Auth
	Tender
	Bid
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Tender: NewTenderRepo(db),
		Auth:   NewAuthRepo(db),
		Bid:    NewBidsRepo(db),
	}
}
