package pg

import "github.ru/GeorgVartanov/todo-app/pkg/user/models"

func (s *Storage) SelectAll() ([]models.User, error) {
	var user []models.User
	//_, err := s.DB.Model(user).Returning("*").Insert()
	err := s.DB.Model(&user).Select()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Storage) Select(user *models.User) (*models.User, error) {
	err := s.DB.Model(user).Where("email=?", user.Email).Select()
	if err != nil {
		return nil, err
	}
	return user, nil
}
