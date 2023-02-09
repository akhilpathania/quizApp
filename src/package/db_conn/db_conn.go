package db_conn

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "akhil"
	dbname   = "first_db"
)

func InitDB() (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port = %d user= %s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)

	return db, err
}
