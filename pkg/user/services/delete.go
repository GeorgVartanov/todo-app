package services

import "github.ru/GeorgVartanov/todo-app/pkg/user/models"

func (s *service) Delete(userIn *models.UserIn) error {
	var user models.User
	user.ConvertToUserFromUserIn(userIn)
	err := s.repository.Delete(&user)
	if err != nil {
		return err
	}
	return nil
}
