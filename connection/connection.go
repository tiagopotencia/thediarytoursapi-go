package connection

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"fmt"
)

func GetConnection() (db *sqlx.DB) {
	//PROD
	db, err := sqlx.Connect("postgres", "host=ec2-54-163-251-104.compute-1.amazonaws.com user=hmojojwszvhfyi password=ma0K_yrpxftus3Yp5PP_tWcOy3 port=5432 dbname=d8oigdpqn2tb6l sslmode=require	")
	//DEV
	// db, err := sqlx.Connect("postgres", "user=postgres password=root port=5433 dbname=thediarytours sslmode=disable")
	if err != nil {
		fmt.Print(err.Error())
	}

	return db

}