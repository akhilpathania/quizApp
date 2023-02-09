package main

import (
	"fmt"
	"quizapp/src/package/db_conn"
)

func main() {
	db, err := db_conn.InitDB()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
}
