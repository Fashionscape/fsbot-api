package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host     string = "localhost"
	port     string = "5432"
	user     string = "fsbot"
	password string = ""
	database string = "fashionscape"
)

func Connection() *sqlx.DB {
	db, err := sqlx.Connect("database", fmt.Sprintf("user=%s dbname=%s sslmode=disable", user, database))
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic("Could not connect to database: " + err.Error())
	}

	return db
}
