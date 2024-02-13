package service

import (
	"sessionmanagement/api/model/request"
	"sessionmanagement/api/model/response"
	"sessionmanagement/api/repository"
)

type AuthService interface {
	UserLogin(user request.User) (response.User, error)
}

type authService struct {
	authRepository repository.AuthRepository
}

func NewAuthService(a repository.AuthRepository) AuthService {
	return authService{
		authRepository: a,
	}
}

func (a authService) UserLogin(user request.User) (response.User, error) {
	return a.authRepository.UserLogin(user)
}