package service

import (
	"github.com/fanfaronDo/test_avito/internal/domain"
	"github.com/fanfaronDo/test_avito/internal/repo"
)

type Auth interface {
	CheckUserCharge(username, organizationID string) (string, bool)
}

type Tender interface {
	CreateTender(tenderCreator domain.TenderCreator, uuid string) (domain.Tender, error)
	GetTenders(limit, offset int, serviceType string) ([]domain.Tender, error)
	GetTendersByUsername(limit, offset int, username string) ([]domain.Tender, error)
	GetStatusTenderByTenderID(tenderID, username string) (string, error)
	UpdateStatusTender(tenderUUID, status, username string) (domain.Tender, error)
	EditTender(tenderUUID, username string, tenderEditor *domain.TenderEditor) (domain.Tender, error)
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
