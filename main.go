package main

import (
	"encoding/json"
	"net/http"
)

var PORT = "localhost:8080"

type User struct {
	ID   string
	Name string
}

var userData = []User{
	{ID: "1", Name: "fadli"},
	{ID: "2", Name: "joko"},
}

func main() {

	http.HandleFunc("/getuserdata", getUserData)
	http.HandleFunc("/createuser", createUser)
	http.ListenAndServe(PORT, nil)

}

func getUserData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(userData)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)

}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")

		newUser := User{
			ID:   id,
			Name: name,
		}

		userData = append(userData, newUser)

		json.NewEncoder(w).Encode(newUser)
		return
	}

}
