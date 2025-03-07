package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "Amrit"},
	{ID: 2, Name: "Motti"},
}
var mu sync.Mutex

func loggingmiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("method is %s and url is %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)

	})
}

func authmiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apikey := r.Header.Get("X-API-Key")
		if apikey != "secret123" {
			http.Error(w, "unautorized", http.StatusUnauthorized)
			return
		}
		fmt.Println("Authenticated with right API key")
		next.ServeHTTP(w, r)
	})
}

func getusers(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)

}

func createuser(w http.ResponseWriter, r *http.Request) {
	var user User

	json.NewDecoder(r.Body).Decode(&user)

	mu.Lock()
	defer mu.Unlock()

	users = append(users, user)
	// fmt.Fprintf(w, "user created")
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getusers(w, r)
		case http.MethodPost:
			createuser(w, r)
		default:
			http.Error(w, "no method matched", http.StatusMethodNotAllowed)
		}
	})

	loggingmux := loggingmiddleware(mux)
	authmux := authmiddleware(loggingmux)
	fmt.Println("Server started ")
	http.ListenAndServe(":8080", authmux)

}
