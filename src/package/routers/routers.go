package routers

import (
	"database/sql"
	"fmt"
	"net/http"

	"encoding/json"

	"quizapp/src/package/db_conn"
	"quizapp/src/package/models"
	"quizapp/src/package/utilities"

	gMux "github.com/gorilla/mux"
)

var db *sql.DB
var err error

func Init() {
	fmt.Println("Package::routers::Init")

	db, err = db_conn.InitDB()
	if err != nil {
		fmt.Println(err)
	}

	router := gMux.NewRouter()

	router.HandleFunc("/", Index).Methods("GET")
	router.HandleFunc("/signup", Signup).Methods("POST")

	http.ListenAndServe(":5000", router)
}

func Index(q http.ResponseWriter, r *http.Request) {

}

func Signup(q http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	if user.Username == "" || user.Password == "" || user.FirstName == "" || user.LastName == "" {
		q.WriteHeader(400)
		q.Write([]byte("Wrong Response"))
		return
	}
	user.Password = utilities.GeneratePassword(user.Password)
	
}
