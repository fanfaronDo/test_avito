package service

import (
	"github.com/fanfaronDo/test_avito/internal/domain"
	"github.com/fanfaronDo/test_avito/internal/handler"
	"github.com/fanfaronDo/test_avito/internal/repo"
)

type Tender interface {
	CreateTender(tenderCreator handler.TenderCreator) (domain.Tender, error)
}

type Service struct {
	Tender
}

func NewService(repo repo.Repository) *Service {
	return &Service{
		NewTenderService(repo),
	}
}
