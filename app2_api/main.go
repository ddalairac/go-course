package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	Port     = ":8080"
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = ""
	dbname   = "postgres"
)

type Animal struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var db *sql.DB

func initDB() {
	fmt.Println("initDB")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to database:", dbname)
}

func AnimalGetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AnimalGetAll")
	query := `SELECT id, name FROM animals`
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var animals []Animal
	for rows.Next() {
		var animal Animal
		err := rows.Scan(&animal.ID, &animal.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		animals = append(animals, animal)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(animals)
}

func AnimalPost(w http.ResponseWriter, r *http.Request) {
	var animal Animal
	err := json.NewDecoder(r.Body).Decode(&animal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO animals (name) VALUES ($1) RETURNING id`
	err = db.QueryRow(query, animal.Name).Scan(&animal.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(animal)
}

func AnimalGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var animal Animal
	query := `SELECT id, name FROM animals WHERE id = $1`
	err = db.QueryRow(query, id).Scan(&animal.ID, &animal.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Animal not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(animal)
}

func AnimalDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	query := `DELETE FROM animals WHERE id = $1`
	result, err := db.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Animal not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	fmt.Println("App2 starting")

	// Initialize database connection
	initDB()

	// Create router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/animals", AnimalGetAll).Methods("GET")
	router.HandleFunc("/animals", AnimalPost).Methods("POST")
	router.HandleFunc("/animals/{id}", AnimalGet).Methods("GET")
	router.HandleFunc("/animals/{id}", AnimalDelete).Methods("DELETE")

	// Start server
	fmt.Println("Server is running on port", Port)
	log.Fatal(http.ListenAndServe(Port, router))
}
