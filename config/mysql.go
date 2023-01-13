package config

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/ecommerce_api")
	if err != nil {
		panic(err)
		// error 500 msg failed connect to database
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	defer db.Close()

	return db
}