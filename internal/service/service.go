package service

import (
	"github.com/fanfaronDo/test_avito/internal/domain"
	"github.com/fanfaronDo/test_avito/internal/repo"
)

type Auth interface {
	CheckUserCharge(username string) (string, error)
	CheckUserExists(username string) (string, error)
}

type Tender interface {
	CreateTender(tenderCreator domain.TenderCreator, uuid string) (domain.Tender, error)
	GetTenders(limit, offset int, serviceType string) ([]domain.Tender, error)
	GetTendersByUserUUID(limit, offset int, uuid string) ([]domain.Tender, error)
	GetStatusTenderByTenderID(tenderID, uuid string) (string, error)
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
