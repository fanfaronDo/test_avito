package service

import (
	"github.com/fanfaronDo/test_avito/internal/domain"
	"github.com/fanfaronDo/test_avito/internal/handler"
	"github.com/fanfaronDo/test_avito/internal/repo"
	"time"
)

type TenderService struct {
	repo repo.Repository
}

func NewTenderService(repo repo.Repository) *TenderService {
	return &TenderService{repo: repo}
}

func (t *TenderService) CreateTender(tenderCreator handler.TenderCreator, uuid string) (domain.Tender, error) {
	var tender domain.Tender
	tender = t.initTender(tenderCreator, uuid)

	return t.repo.CreateTender(tender)
}

func (t *TenderService) initTender(creator handler.TenderCreator, uuid string) domain.Tender {
	return domain.Tender{
		Name:           creator.Name,
		Description:    creator.Description,
		ServiceType:    creator.ServiceType,
		Status:         creator.Status,
		OrganizationID: creator.OrganizationID,
		CreatorID:      uuid,
		Version:        1,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}
