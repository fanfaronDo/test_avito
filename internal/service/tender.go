package service

import (
	"github.com/fanfaronDo/test_avito/internal/domain"
	"github.com/fanfaronDo/test_avito/internal/handler"
	"github.com/fanfaronDo/test_avito/internal/repo"
	"time"
)

const (
	defaultLimit = 5
)

type TenderService struct {
	repo *repo.Repository
}

func NewTenderService(repo *repo.Repository) *TenderService {
	return &TenderService{repo: repo}
}

func (t *TenderService) CreateTender(tenderCreator handler.TenderCreator, uuid string) (domain.Tender, error) {
	var tender domain.Tender
	tender = t.initTender(tenderCreator, uuid)

	return t.repo.CreateTender(tender)
}

func (t *TenderService) GetTenders(limit, offset int, serviceType string) ([]domain.Tender, error) {
	return t.repo.Tender.GetTenders(limit, offset, serviceType)
}

func (t *TenderService) GetTendersByUsername(limit, offset int, username string) ([]domain.Tender, error) {
	uuid, err := t.repo.GetUserUUID(username)
	if err != nil {
		return nil, ErrTenderNotFound
	}
	return t.repo.GetTendersByUserID(limit, offset, uuid)
}

func (t *TenderService) GetStatusTenderByTenderID(tenderID, username string) (string, error) {
	uuid, err := t.repo.GetUserUUID(username)
	if err != nil {
		return "", ErrUserNotFound
	}
	status, err := t.repo.Tender.GetStatusTenderById(tenderID, uuid)
	if err != nil {
		return "", ErrTenderNotFound
	}
	return status, nil
}

func (t *TenderService) UpdateStatusTender(tenderUUID, status, username string) (domain.Tender, error) {
	uuid, err := t.repo.GetUserUUID(username)
	if err != nil {
		return domain.Tender{}, ErrUserNotFound
	}
	return t.repo.UpdateStatusTenderById(tenderUUID, status, uuid)
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

//func (t *TenderService) UpdateTender(tender domain.Tender) error {}
