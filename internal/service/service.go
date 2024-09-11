package service

import (
	"github.com/fanfaronDo/test_avito/internal/domain"
	"github.com/fanfaronDo/test_avito/internal/repo"
)

type Auth interface {
	GetUserId(username string) (string, error)
	GetUserChargeId(username string) (string, error)
	CheckUserCharge(userUUID, organisationid string) (string, error)
	CreateUserCharge(userUUID, username string) (string, error)
	IsUserChargeExist(username string) bool
}

type Tender interface {
	CreateTender(tenderCreator domain.TenderCreator, userUUID string) (domain.Tender, error)
	GetTenders(limit, offset int, serviceType string) ([]domain.Tender, error)
	GetTendersByUserUUID(limit, offset int, userUUID string) ([]domain.Tender, error)
	GetStatusTenderByTenderID(tenderID, userUUID string) (string, error)
	UpdateStatusTender(tenderUUID, status, userUUID string) (domain.Tender, error)
	EditTender(tenderUUID, userUUID string, tenderEditor *domain.TenderEditor) (domain.Tender, error)
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
