package models

import (
	"database/sql"
	"log"
	"time"
)

type DBHandler struct {
	db *sql.DB
}

func (h *DBHandler) Insert(username string, passwordHash string, tenant string) {
	stmt, err := h.db.Prepare("INSERT credentials SET username=?,password=?,tenant=?,created=?,modified=?")
	checkErr(err)

	var datetime = time.Now()
	datetime.Format(time.RFC3339)
	res, err := stmt.Exec(username, passwordHash, tenant, datetime, datetime)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	log.Println("credentials inserted", id)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
