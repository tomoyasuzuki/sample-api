package service

import (
	"github.com/tomoyasuzuki/sample-api/app/model"
	"github.com/tomoyasuzuki/sample-api/app/repository"
)

type IUserService interface {
	GetUser(id string) (model.User, error)
	GetUsers() ([]model.User, error)
	CreateUser(user model.User) error
	UpdateUser(user model.User) error
}

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(repo *repository.IUserRepository) IUserService {
	return &UserService{*repo}
}

func (us *UserService) GetUser(id string) (model.User, error) {
	return us.repo.GetUser(id)
}

func (us *UserService) GetUsers() ([]model.User, error) {
	return us.repo.GetUsers()
}

func (us *UserService) CreateUser(user model.User) error {
	return us.repo.CreateUser(user)
}

func (us *UserService) UpdateUser(user model.User) error {
	return us.repo.UpdateUser(user)
}
