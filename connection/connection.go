package connection

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"fmt"
)

func GetConnection() (db *sqlx.DB) {
	db, err := sqlx.Connect("postgres", "user=postgres password=root port=5433 dbname=thediarytours sslmode=disable")
	if err != nil {
		fmt.Print(err.Error())
	}

	return db

}