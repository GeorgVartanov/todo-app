package services

import "github.ru/GeorgVartanov/todo-app/pkg/user/models"

type Repository interface {
	Insert(user *models.User) error
	Select(user *models.User) (*models.User, error)
	SelectAll() ([]models.User, error)
	Update(user *models.User) error
	Delete(user *models.User) error
}

type Service interface {
	Create(userIn *models.UserIn) error
	Read(userIn *models.UserIn) (*models.UserOut, error)
	ReadAll() ([]models.UserOut, error)
	Update(userIn *models.UserIn) error
	Delete(userIn *models.UserIn) error
}

type service struct {
	repository Repository
	User       *models.User
}

func NewService(repo Repository) Service {
	return &service{repository: repo}
}
