package fileStorageService

import (
	"log"
)

// SaveChunkToDB saves a file chunk to the database
func SaveChunkToDB(fileID string, chunkData []byte) error {
	query := `INSERT INTO file_chunks (file_id, chunk_data) VALUES ($1, $2)`
	_, err := DB.Exec(query, fileID, chunkData)
	if err != nil {
		log.Println("Error inserting file chunk:", err)
		return err
	}
	return nil
}

// GetFileChunksByID retrieves file chunks by file ID from the database
func GetFileChunksByID(fileID string) []string {
	query := `SELECT chunk_data FROM file_chunks WHERE file_id = $1`
	rows, err := DB.Query(query, fileID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var chunks []string
	for rows.Next() {
		var chunk string
		err := rows.Scan(&chunk)
		if err != nil {
			log.Fatal(err)
		}
		chunks = append(chunks, chunk)
	}

	return chunks
}

// GetAllFileIDs retrieves all file IDs from the database
func GetAllFileIDs() []string {
	query := `SELECT DISTINCT file_id FROM file_chunks`
	rows, err := DB.Query(query)
	if err != nil {
		log.Println("Error retrieving file IDs:", err)
		return nil
	}
	defer rows.Close()

	var fileIDs []string
	for rows.Next() {
		var fileID string
		err := rows.Scan(&fileID)
		if err != nil {
			log.Println("Error scanning file ID:", err)
			continue
		}
		fileIDs = append(fileIDs, fileID)
	}

	return fileIDs
}
