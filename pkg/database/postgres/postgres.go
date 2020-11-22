package postgres

import (
	"context"
	"github.com/go-pg/pg/v10"
)

type DB struct {
	*pg.DB
}

func NewDB() *DB {
	var db DB
	db.Connect()
	return &db
}

func (db *DB) Connect() {
	pgDB := pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "postgres",
		Database: "todo-app",
	})
	//_, err := db.Exec("SELECT 1")
	ctx := context.Background()
	err := pgDB.Ping(ctx)
	if err != nil {
		panic(err)
	}
	db.DB = pgDB
}
