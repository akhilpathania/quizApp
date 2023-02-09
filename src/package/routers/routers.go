package routers

import (
	"database/sql"
	"fmt"
	"log"
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
	router.HandleFunc("/shutServer", shutServer).Methods("POST")

	server := &http.Server{
		Addr:    ":5000",
		Handler: router,
	}

	//http.ListenAndServe(":5000", router)
	go func() {
		log.Fatal(server.ListenAndServe())
	}()
	log.Println("server started")

	// stopC := make(chan, os.Interrupt)
	// signal.Notify(stopC, os.Interrupt)
	// <-stopC

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// log.Println("server stopping...")
	// defer cancel()

	// log.Fatal(server.Shutdown(ctx))
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
	q.Header().Set("Content-Type", "application/json")
	json.NewEncoder(q).Encode(user)
	err := db_conn.InsertStmt(user, db)
	if err != nil {
		log.Println(err)
	}

}

func shutServer(q http.ResponseWriter, r *http.Request) {

}
