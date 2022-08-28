package client

import (
	"database/sql"
)

type PostgresDB struct {
	SQL *sql.DB
}

var Postgres = &PostgresDB{}

func Connect(connStr string) *PostgresDB {
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	Postgres.SQL = db

	return Postgres
}