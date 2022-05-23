package service

import (
	"github.com/Sebastian1811/backend-proyecto1-www/entity"
	"github.com/Sebastian1811/backend-proyecto1-www/repository"
)

type UserService interface {
	Register(entity.User)
	Login(int) entity.User
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		userRepository: repo,
	}
}

func (service *userService) Register(user entity.User) {
	service.userRepository.Register(user)
}

func (service *userService) Login(id int) entity.User {
	return service.userRepository.Login(id)
}
