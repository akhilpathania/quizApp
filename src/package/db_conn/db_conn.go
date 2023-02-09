package db_conn

import (
	"database/sql"
	"fmt"
	"quizapp/src/package/models"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "akhil"
	dbname   = "quizapp"
)

func InitDB() (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port = %d user= %s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)

	return db, err
}

func InsertStmt(user models.User, dBase *sql.DB) error {

	insertStmt := `insert into "User" ("ID", "Username", "Password", "FirstName", 
		"LastName", "TotalScore", "Admin") values($1, $2, $3, $4, $5, $6,$7)`
	_, e := dBase.Exec(insertStmt, user.ID, user.Username, user.Password, user.FirstName, user.LastName, user.TotalScore, user.Admin)
	return e
}
