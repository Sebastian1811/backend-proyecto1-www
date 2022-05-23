package service

import (
	"github.com/Sebastian1811/backend-proyecto1-www/entity"
	"github.com/Sebastian1811/backend-proyecto1-www/repository"
)

type UserService interface {
	register(entity.User)
	login(int) entity.User
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		userrepository: repo,
	}
}
