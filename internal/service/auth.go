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

func (a *AuthService) GetUserChargeId(username string) (string, error) {
	return a.repo.Auth.GetUserChargeUUID(username)
}

func (a *AuthService) IsUserChargeExist(username string) bool {
	_, err := a.repo.Auth.GetUserChargeUUID(username)
	if err != nil {
		return false
	}
	return true
}

func (a *AuthService) CheckUserCharge(userUUID, organisationid string) (string, error) {
	return a.repo.Auth.CheckUserCharge(userUUID, organisationid)
}

func (a *AuthService) CreateUserCharge(userUUID, username string) (string, error) {
	return a.repo.Auth.CreateUserCharge(userUUID, username)
}
