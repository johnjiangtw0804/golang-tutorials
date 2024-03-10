package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Global variriables
var router = mux.NewRouter()
var database Storage

// Server init
func init() {
	var err error

	// Use memory as storage space (mimic the database action)
	database, err = NewStorage(Memory)
	if err != nil {
		log.Fatal(err)
	}

	// init the router
	router.HandleFunc("/usersList", GetUsers).Methods("GET")
	router.HandleFunc("/user/{user_id}", GetUser).Methods("GET")
	router.HandleFunc("/user", CreateUser).Methods("POST")
	router.HandleFunc("/user/{user_id}", DeleteUser).Methods("DELETE")
}

func main() {
	// starting the server
	log.Println("Server starting... listening at port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
