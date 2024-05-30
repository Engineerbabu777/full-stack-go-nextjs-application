package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// main function!
func main() {
	// connect to database!
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal("Could not connect to database!")
	}

	defer db.Close()

	// CREATE TABLE IF NOT EXISTS!
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name VARCHAR(255), email VARCHAR(255))")

	if err != nil {
		log.Fatal(err)
	}

	// router:
	router := mux.NewRouter()
	router.HandleFunc("/api/go/users", getUsers(db)).Methods("GET");
	router.HandleFunc("/api/go/users", createUser(db)).Methods("POST");
	router.HandleFunc("/api/go/users/{id}", getUser(db)).Methods("GET");
	router.HandleFunc("/api/go/users/{id}", updateUser(db)).Methods("PUT");
	router.HandleFunc("/api/go/users/{id}", deleteUser(db)).Methods("DELETE");


 // WRAP!
 enhancedRouter := enableCORS(jsonContentTypeMiddleware(router));

 // start server!
 log.Fatal(http.ListenAndServe(":8000", enhancedRouter))



}
