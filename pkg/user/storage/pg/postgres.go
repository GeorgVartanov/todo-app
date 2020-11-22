package pg

import (
	"github.com/go-pg/pg/v10/orm"
	"github.ru/GeorgVartanov/todo-app/pkg/database/postgres"
	"github.ru/GeorgVartanov/todo-app/pkg/user/models"
)

type Storage struct {
	DB *postgres.DB
}

func NewStorage(DB *postgres.DB) (*Storage, error) {
	//var user models.User
	err := DB.Model((*models.User)(nil)).CreateTable(&orm.CreateTableOptions{
		Temp: false,
		IfNotExists: true,
	})
	if err != nil {
		return nil, err
	}
	return &Storage{DB}, err
}
