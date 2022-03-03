package main

import (
	"fmt"
	"net/http"
)

// Untuk menampilkan ke website
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Index called ...")
	hello := "Hello World"
	w.Write([]byte(hello))
}

// Untuk melakukan routing
func handleRoute() {
	http.HandleFunc("/", index)
}

func main() {
	// port bebas, yang penting belum digunakan sama sekali.
	// saya rekomendasikan membuat port yang simple
	// seperti 3000, 3002, 9000, 9090, 9999, dll..
	port := ":9999"
	startServer(port)
}

// Untuk membuat web server

func startServer(port string) {
	handleRoute()
	fmt.Println("Server running at port ", port)
	http.ListenAndServe(port, nil)
}
