package pg

import "github.ru/GeorgVartanov/todo-app/pkg/user/models"

func (s *Storage) Update(user *models.User) error {

	_, err := s.DB.Model(user).Where("email=?", user.Email).Update()
	if err != nil {
		return err
	}
	return nil
}
