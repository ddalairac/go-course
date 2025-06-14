// Package conn handles the database connection setup and provides a global DB instance.
package conn

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Database connection configuration constants
const (
	Port     = ":8080"         // Application port (not used for DB connection)
	host     = "localhost"     // Database host
	port     = 5433            // Database port
	user     = "postgres"      // Database user
	password = ""              // Database password
	dbname   = "postgres"      // Database name
)

// DB is the global database connection pool
var DB *sql.DB

// InitDB initializes the database connection and assigns it to the global DB variable.
func InitDB() {
	fmt.Println("initDB")
	// Build the PostgreSQL connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	// Open a new database connection using the connection string
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		// If an error occurs while opening the connection, log and exit
		log.Fatal(err)
	}

	// Ping the database to verify the connection is alive
	err = DB.Ping()
	if err != nil {
		// If the ping fails, log and exit
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to database:", dbname)
}
