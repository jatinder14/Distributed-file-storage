package main

import (
	"log"

	"github.com/tushar/Assignment-HyperHire/fileStorageService"
)

func main() {
	// Initialize the database connection
	fileStorageService.InitDB()

	// Start the server
	log.Println("Server running on :8080")
	err := fileStorageService.StartServer()
	if err != nil {
		log.Fatal(err)
	}
}
