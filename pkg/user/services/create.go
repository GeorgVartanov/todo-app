package services

import "github.ru/GeorgVartanov/todo-app/pkg/user/models"

func (s *service) Create(userIn *models.UserIn) error {
	var user models.User
	user.ConvertToUserFromUserIn(userIn)
	err := s.repository.Insert(&user)
	if err != nil {
		return err
	}
	return nil
}
