package fileStorageService

import (
	"crypto/rand"
	"fmt"
	"mime/multipart"
	"io"
)

const chunkSize = 1024 * 1024 // 1MB per chunk

// SplitFile splits a file into smaller chunks
func SplitFile(file multipart.File, tempDir string) ([][]byte, error) {
	var chunks [][]byte
	buffer := make([]byte, chunkSize)

	for {
		bytesRead, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if bytesRead == 0 {
			break
		}
		// Append the chunk to the slice
		chunkCopy := make([]byte, bytesRead)
		copy(chunkCopy, buffer[:bytesRead])
		chunks = append(chunks, chunkCopy)
	}

	return chunks, nil
}

// GenerateFileID creates a unique file IDs
func GenerateFileID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
