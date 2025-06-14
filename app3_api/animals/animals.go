// Package animals provides HTTP handlers and data access logic
// for managing the Animal entity in the database.
package animals

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/diegodalairac/go-course/app3_api/conn"
	"github.com/gorilla/mux"
)

// Animal represents the animal entity stored in the database.
type Animal struct {
	ID   int    `json:"id"`   // Unique animal ID
	Name string `json:"name"` // Animal name
}

// GetAll retrieves and returns all animals from the database as JSON.
func GetAll(w http.ResponseWriter, r *http.Request) {
	// Prepare the SQL query to get all animals
	query := `SELECT id, name FROM animals`
	rows, err := conn.DB.Query(query)
	if err != nil {
		// If an error occurs while executing the query, return a 500 error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close() // Ensure resources are released at the end

	var animals []Animal
	// Iterate over each row in the result
	for rows.Next() {
		var animal Animal
		// Scan the row values into the Animal struct
		err := rows.Scan(&animal.ID, &animal.Name)
		if err != nil {
			// If an error occurs while scanning, return a 500 error
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		animals = append(animals, animal)
	}

	// Set the response header to JSON
	w.Header().Set("Content-Type", "application/json")
	// Encode and send the list of animals as JSON
	json.NewEncoder(w).Encode(animals)
}

// Post creates a new animal in the database from the JSON received in the body.
func Post(w http.ResponseWriter, r *http.Request) {
	var animal Animal
	// Decode the JSON received in the body into the Animal struct
	err := json.NewDecoder(r.Body).Decode(&animal)
	if err != nil {
		// If the JSON is invalid, return a 400 error
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Prepare the SQL query to insert the new animal
	query := `INSERT INTO animals (name) VALUES ($1) RETURNING id`
	// Execute the query and get the generated ID
	err = conn.DB.QueryRow(query, animal.Name).Scan(&animal.ID)
	if err != nil {
		// If an error occurs while inserting, return a 500 error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the response header to JSON
	w.Header().Set("Content-Type", "application/json")
	// Encode and send the created animal as JSON
	json.NewEncoder(w).Encode(animal)
}

// Get retrieves a specific animal by its ID and returns it as JSON.
func Get(w http.ResponseWriter, r *http.Request) {
	// Get the URL parameters
	vars := mux.Vars(r)
	// Convert the "id" parameter to an integer
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// If the ID is not valid, return a 400 error
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var animal Animal
	// Prepare the SQL query to find the animal by ID
	query := `SELECT id, name FROM animals WHERE id = $1`
	// Execute the query and scan the result into the Animal struct
	err = conn.DB.QueryRow(query, id).Scan(&animal.ID, &animal.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			// If the animal is not found, return a 404 error
			http.Error(w, "Animal not found", http.StatusNotFound)
			return
		}
		// If another error occurs, return a 500 error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the response header to JSON
	w.Header().Set("Content-Type", "application/json")
	// Encode and send the found animal as JSON
	json.NewEncoder(w).Encode(animal)
}

// Delete removes a specific animal by its ID from the database.
func Delete(w http.ResponseWriter, r *http.Request) {
	// Get the URL parameters
	vars := mux.Vars(r)
	// Convert the "id" parameter to an integer
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// If the ID is not valid, return a 400 error
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Prepare the SQL query to delete the animal by ID
	query := `DELETE FROM animals WHERE id = $1`
	// Execute the query
	result, err := conn.DB.Exec(query, id)
	if err != nil {
		// If an error occurs while deleting, return a 500 error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Check how many rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		// If an error occurs while getting the number of rows, return a 500 error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		// If no animal was deleted, return a 404 error
		http.Error(w, "Animal not found", http.StatusNotFound)
		return
	}

	// Return a 204 No Content status if the deletion was successful
	w.WriteHeader(http.StatusNoContent)
}

// Put updates the name of a specific animal by its ID using the data received in the body.
func Put(w http.ResponseWriter, r *http.Request) {
	// Get the URL parameters
	vars := mux.Vars(r)
	// Convert the "id" parameter to an integer
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// If the ID is not valid, return a 400 error
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var animal Animal
	// Decode the JSON received in the body into the Animal struct
	err = json.NewDecoder(r.Body).Decode(&animal)
	if err != nil {
		// If the JSON is invalid, return a 400 error
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Prepare the SQL query to update the animal's name
	query := `UPDATE animals SET name = $1 WHERE id = $2`
	// Execute the update query
	_, err = conn.DB.Exec(query, animal.Name, id)
	if err != nil {
		// If an error occurs while updating, return a 500 error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retrieve the updated animal
	query = `SELECT id, name FROM animals WHERE id = $1`
	err = conn.DB.QueryRow(query, id).Scan(&animal.ID, &animal.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the updated animal as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(animal)
}
