package service

import "github.com/fanfaronDo/test_avito/internal/repo"

type AuthService struct {
	repo *repo.Repository
}

func NewAuthService(repo *repo.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) CheckUserCharge(username, organizationID string) (string, bool) {
	uuid, err := a.repo.GetUserUUIDCharge(username, organizationID)
	if err != nil {
		return "", false
	}
	return uuid, true
}
