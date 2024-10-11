package fileStorageService

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// DB is the global database connection
var DB *sql.DB

// CreateDB creates the file_storage database if it doesn't exist
func CreateDB() {
	// Connect to the default 'postgres' database to create a new database
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=%s sslmode=%s",
		DBHost, DBUser, DBPassword, DBPort, DBSSLMode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to postgres database:", err)
	}
	defer db.Close()

	// Check if the file_storage database already exists
	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE datname = $1)", DBName).Scan(&exists)
	if err != nil {
		log.Fatal("Error checking if database exists:", err)
	}

	// If the database doesn't exist, create it
	if !exists {
		_, err = db.Exec("CREATE DATABASE " + DBName)
		if err != nil {
			log.Fatal("Error creating database:", err)
		}
		log.Printf("Database '%s' created successfully", DBName)
	} else {
		log.Printf("Database '%s' already exists", DBName)
	}
}

// InitDB initializes the database connection and creates the necessary tables
func InitDB() {
	LoadConfig() // Make sure configuration is loaded
	CreateDB()
	ConnectDB()
}

// Call InitDB when the server starts
func init() {
	InitDB()
}

func ConnectDB() {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		DBHost, DBUser, DBPassword, DBName, DBPort, DBSSLMode)
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database connection:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	createFileChunksTable()

	log.Println("Database connected successfully!")
}

func createFileChunksTable() {
	query := `
		CREATE TABLE IF NOT EXISTS file_chunks (
			id SERIAL PRIMARY KEY,
			file_id VARCHAR(255) NOT NULL,
			chunk_data BYTEA NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal("Failed to create file_chunks table:", err)
	}
}
