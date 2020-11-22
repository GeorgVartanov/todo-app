package models

import (
	"fmt"
)

//User from model
type User struct {
	Id          uint64 `pg:",pk"`
	Email       string `pg:",unique,notnull, type:varchar(255)"`
	Password    string
	DisplayName string
}

func (u *User) String() string {
	return fmt.Sprintf("User<%d %s %s>", u.Id, u.Email, u.DisplayName)
}

func (u *User) ConvertToUserFromUserIn(user *UserIn) {
	u.Email = user.Email
	u.Password = user.Password
	u.DisplayName = user.DisplayName
}

func (u User) ConvertToUserOut() *UserOut {
	return &UserOut{
		Id:          u.Id,
		Email:       u.Email,
		DisplayName: u.DisplayName,
	}
}