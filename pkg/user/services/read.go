package services

import "github.ru/GeorgVartanov/todo-app/pkg/user/models"

func (s *service) Read(userIn *models.UserIn) (*models.UserOut, error) {
	var user models.User
	user.ConvertToUserFromUserIn(userIn)
	userDB, err := s.repository.Select(&user)
	if err != nil {
		return nil, err
	}
	userOut := userDB.ConvertToUserOut()
	return userOut, nil
}

func (s *service) ReadAll() ([]models.UserOut, error) {
	users, err := s.repository.SelectAll()
	if err != nil {
		return nil, err
	}
	var allUsers []models.UserOut
	for _, v := range users {
		userOut := v.ConvertToUserOut()
		allUsers = append(allUsers, *userOut)
	}
	return allUsers, nil
}
