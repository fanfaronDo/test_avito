package service

import "github.com/fanfaronDo/test_avito/internal/repo"

type AuthService struct {
	repo *repo.Repository
}

func NewAuthService(repo *repo.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) CheckUserExists(username string) (string, error) {
	userUUID, err := a.repo.GetUserUUID(username)
	if err != nil {
		return "", ErrUserNotFound
	}
	return userUUID, err
}

func (a *AuthService) CheckUserCharge(userUUID string) (string, error) {
	_, err := a.repo.FindUserUUIDCharge(userUUID)
	if err != nil {
		return "", ErrUnauthorizedError
	}

	return userUUID, nil
}
