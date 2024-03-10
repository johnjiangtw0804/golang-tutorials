package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user, err = database.GetUsers()
	if err != nil {
		// no user found code 204
		fmt.Fprint(w, err)
		w.WriteHeader(http.StatusNoContent)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// database provides API for GetUser service
	var user, err = database.GetUser(params["user_id"])
	if err != nil {
		// no user found code 204
		fmt.Fprint(w, err)
		w.WriteHeader(http.StatusNoContent)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	var err = database.CreateUser(user)
	if err != nil {
		// something wrong with the User fields
		fmt.Fprint(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var err = database.DeleteUser(params["user_id"])
	if err != nil {
		fmt.Fprint(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
