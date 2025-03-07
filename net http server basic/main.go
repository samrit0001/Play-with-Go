package main

import (
	"fmt"
	"net/http"
)

func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func getusers(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not allowed only GET", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "list of users....")
}

func createuser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only Post Method allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "user is created ")
}

func main() {
	http.HandleFunc("/", hellohandler)
	http.HandleFunc("/getusers", getusers)
	http.HandleFunc("/create", createuser)

	fmt.Printf("Starting the server")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
