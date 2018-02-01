package models

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB(connectionString string) *DBHandler {
	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		log.Panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	err = db.Ping()
	if err != nil {
		log.Panic(err.Error()) // proper error handling instead of panic in your app
	}
	log.Println("db connection ready")

	return &DBHandler{db: db}
}
