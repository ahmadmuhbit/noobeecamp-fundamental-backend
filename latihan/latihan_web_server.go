package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Class string `json:"class"`
}

// Untuk menampilkan ke website
func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users := []User{
		User{Name: "User 1", Class: "Flower A"},
		User{Name: "User 2", Class: "Flower B"},
		User{Name: "User 3", Class: "Flower C"},
		User{Name: "User 4", Class: "Flower D"},
		User{Name: "User 5", Class: "Flower E"},
	}

	usersJSON, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	w.Write(usersJSON)
}

// Untuk melakukan routing
func handleRoute() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(struct {
			Name  string `json:"name"`
			Class string `json:"class"`
		}{"User 1", "Flower A"})
	})
}

// Untuk membuat web server
func startServer(port string) {
	handleRoute()
	http.ListenAndServe(port, nil)
}

func main() {
	port := ":8080"
	startServer(port)
}
