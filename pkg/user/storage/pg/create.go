package pg

import (
	"github.ru/GeorgVartanov/todo-app/pkg/user/models"
)

func (s *Storage) Insert(user *models.User) error {

	//_, err := s.DB.Model(user).Returning("*").Insert()
	_, err := s.DB.Model(user).Insert()
	if err != nil {
		return err
	}
	return nil
}
