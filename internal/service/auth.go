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

//func (a *AuthService) GetUserCharge(userid string) (string, error) {
//	return a.repo.Auth.GetUserCharge(userid)
//}

func (a *AuthService) CheckUserCreatorTender(userUUID, tenderUUID string) (string, error) {
	return a.repo.Auth.CheckUserCreatorTender(userUUID, tenderUUID)
}
