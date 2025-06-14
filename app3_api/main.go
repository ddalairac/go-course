package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/diegodalairac/go-course/app3_api/animals"
	"github.com/diegodalairac/go-course/app3_api/conn"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("App2 starting")
	// Initialize database connection
	conn.InitDB()

	// Create router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/animals", animals.GetAll).Methods("GET")
	router.HandleFunc("/animals", animals.Post).Methods("POST")
	router.HandleFunc("/animals/{id}", animals.Get).Methods("GET")
	router.HandleFunc("/animals/{id}", animals.Put).Methods("PUT")
	router.HandleFunc("/animals/{id}", animals.Delete).Methods("DELETE")

	// Start server
	fmt.Println("Server is running on port", conn.Port)
	log.Fatal(http.ListenAndServe(conn.Port, router))
}
