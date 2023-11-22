package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Database struct {
	connection *sql.DB
}

var instance *Database

func GetDatabaseInstance() *Database {
	if instance == nil {
		connStr := "host=dpg-cl9sfa5o7jlc73fipc30-a.oregon-postgres.render.com port=5432 user=restorandb_user password=w9w2KqtmDC2VNO8JwDw74sEWjXxUFCJV dbname=restorandb sslmode=require"
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}
		instance = &Database{connection: db}
	}

	return instance
}
