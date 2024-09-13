package service

import "github.com/fanfaronDo/test_avito/internal/repo"

type AuthService struct {
	repo *repo.Repository
}

func NewAuthService(repo *repo.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) GetUserId(username string) (string, error) {
	return a.repo.Auth.GetUserUUID(username)
}

func (a *AuthService) CheckOrganizationAffiliation(userid, organisationid string) (string, error) {
	return a.repo.Auth.CheckOrganizationAffiliation(userid, organisationid)
}

func (a *AuthService) CheckUserCreatorTender(userUUID, tenderUUID string) (string, error) {
	return a.repo.Auth.CheckUserCreatorTender(userUUID, tenderUUID)
}

func (a *AuthService) CheckUserChargeAffiliation(userUUID string) (string, error) {
	return a.repo.Auth.CheckUserChargeAffiliation(userUUID)
}

func (a *AuthService) CheckUserCreatorBids(userUUID, bidsUUID string) (string, error) {
	return a.repo.Auth.CheckUserCreatorBids(userUUID, bidsUUID)
}
func (a *AuthService) CheckUserID(username string) (string, error) {
	return a.repo.Auth.CheckUserID(username)
}

func (a *AuthService) CheckUserCreatorBidsByTenderId(userUUID, tenderUUID string) (string, error) {
	return a.repo.Auth.CheckUserCreatorBidsByTenderId(userUUID, tenderUUID)
}
