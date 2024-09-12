package service

import (
	"github.com/fanfaronDo/test_avito/internal/domain"
	"github.com/fanfaronDo/test_avito/internal/repo"
)

type Auth interface {
	GetUserId(username string) (string, error)
	CheckOrganizationAffiliation(userid, organisationid string) (string, error)
	CheckUserCreatorTender(userUUID, tenderUUID string) (string, error)
	//GetUserCharge(userid string) (string, error)
}

type Tender interface {
	CreateTender(tenderCreator domain.TenderCreator, userUUID string) (domain.Tender, error)
	GetTenders(limit, offset int, serviceType string) ([]domain.Tender, error)
	GetTendersByUserUUID(limit, offset int, userUUID string) ([]domain.Tender, error)
	GetStatusTenderByTenderID(tenderID, userUUID string) (string, error)
	UpdateStatusTender(tenderUUID, status, userUUID string) (domain.Tender, error)
	EditTender(tenderUUID, userUUID string, tenderEditor *domain.TenderEditor) (domain.Tender, error)
	RollbackTender(tenderUUID, userUUID string, version int) (domain.Tender, error)
}

type Service struct {
	Auth
	Tender
}

func NewService(repo *repo.Repository) *Service {
	return &Service{
		NewAuthService(repo),
		NewTenderService(repo),
	}
}
