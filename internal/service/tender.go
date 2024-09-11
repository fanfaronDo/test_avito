package service

import (
	"errors"
	"github.com/fanfaronDo/test_avito/internal/domain"
	"github.com/fanfaronDo/test_avito/internal/repo"
	log "github.com/sirupsen/logrus"
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

func (t *TenderService) CreateTender(tenderCreator domain.TenderCreator, userUUID string) (domain.Tender, error) {
	tenderCreator.Status = "Created"
	if !checkServiceType(tenderCreator.ServiceType) {
		log.Debugf("%s: %v", ErrStatusError, errors.New("service type not found"))
		return domain.Tender{}, ErrServiceTypeError
	}

	var tender domain.Tender
	tender = t.initTender(tenderCreator, userUUID)
	return t.repo.CreateTender(tender, tenderCreator.OrganizationID, userUUID)
}

func (t *TenderService) GetTenders(limit, offset int, serviceType string) ([]domain.Tender, error) {
	return t.repo.Tender.GetTenders(limit, offset, serviceType)
}

func (t *TenderService) GetTendersByUserUUID(limit, offset int, uuid string) ([]domain.Tender, error) {
	return t.repo.GetTendersByUserID(limit, offset, uuid)
}

func (t *TenderService) GetStatusTenderByTenderID(tenderID, userUUID string) (string, error) {
	status, err := t.repo.Tender.GetStatusTenderById(tenderID, userUUID)
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

func (t *TenderService) EditTender(tenderUUID, username string, tenderEditor *domain.TenderEditor) (domain.Tender, error) {
	if tenderEditor == nil {
		return t.repo.GetTenderById(tenderUUID)
	}
	if !checkServiceType(tenderEditor.ServiceType) {
		return domain.Tender{}, ErrServiceTypeError
	}

	uuid, err := t.repo.GetUserUUID(username)
	if err != nil {
		return domain.Tender{}, ErrUserNotFound
	}
	return t.repo.UpdateTender(tenderUUID, uuid, tenderEditor)
}

func (t *TenderService) initTender(creator domain.TenderCreator, uuid string) domain.Tender {
	return domain.Tender{
		Name:        creator.Name,
		Description: creator.Description,
		ServiceType: creator.ServiceType,
		Status:      creator.Status,
		Version:     1,
		CreatedAt:   time.Now(),
	}
}
