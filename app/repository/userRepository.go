package repository

import (
	"errors"
	appError "github.com/tomoyasuzuki/sample-api/app/error"
	"github.com/tomoyasuzuki/sample-api/app/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUser(id string) (model.User, error)
	GetUsers() ([]model.User, error)
	CreateUser(user model.User) error
	UpdateUser(user model.User) error
}

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db}
}

func (u *UserRepository) GetUser(id string) (model.User, error) {
	var user model.User

	if err := u.Db.
		Model(&model.User{}).
		Preload("Tasks").
		First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, appError.New(appError.RecordNotFound)
		} else {
			return user, appError.New(appError.Unknown)
		}
	}

	return user, nil
}

func (u *UserRepository) GetUsers() ([]model.User, error) {
	var users []model.User

	if err := u.Db.
		Model(&model.User{}).
		Preload("Tasks").
		Find(&users).Error; err != nil {
		return users, appError.New(appError.Unknown)
	}
	return users, nil
}

func (u *UserRepository) CreateUser(user model.User) error {
	if err := u.Db.Create(&user).Error; err != nil {
		return appError.New(appError.Unknown)
	}
	return nil
}

func (u *UserRepository) UpdateUser(user model.User) error {
	if err := u.Db.Save(&user).Error; err != nil {
		return appError.New(appError.Unknown)
	}
	return nil
}
