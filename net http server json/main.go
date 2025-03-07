package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID   int    `json:"id"` //struct tags
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "Amrit"},
	{ID: 2, Name: "Motti"},
}

func getusers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "only GET method", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json") //set header to tell client what we sending
	json.NewEncoder(w).Encode(users)                   // this converts the Go struct info to json using Encode and write to Response using the NewEncoder
}

func createuser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only POST method", http.StatusMethodNotAllowed)
		return
	}
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	users = append(users, user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

}

func main() {
	http.HandleFunc("/getusers", getusers)
	http.HandleFunc("/createuser", createuser)

	fmt.Println("Started the server")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
