package fileStorageService

import (
	"errors"
	"net/http"
	"sync"
)

// UploadFileService handles the business logic of file uploads
func UploadFileService(r *http.Request) (string, error) {
	file, _, err := r.FormFile("file")
	if err != nil {
		return "", errors.New("failed to read file")
	}
	defer file.Close()

	fileID := GenerateFileID()
	chunks, err := SplitFile(file, fileID)
	if err != nil {
		return "", err
	}

	var wg sync.WaitGroup
	errChan := make(chan error, len(chunks))

	for _, chunk := range chunks {
		wg.Add(1)
		go func(chunk []byte) {
			defer wg.Done()
			if err := SaveChunkToDB(fileID, chunk); err != nil {
				errChan <- err
			}
		}([]byte(chunk))
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return "", err
		}
	}

	return fileID, nil
}

// Add this function to the existing file
func GetAllFilesService() []string {
	return GetAllFileIDs()
}

// DownloadFileService handles the business logic of file downloads
func DownloadFileService(fileID string) ([]byte, error) {
	chunks := GetFileChunksByID(fileID)
	if len(chunks) == 0 {
		return nil, errors.New("file not found")
	}

	// Combine chunks into a single file
	var fileContent []byte
	for _, chunk := range chunks {
		fileContent = append(fileContent, []byte(chunk)...)
	}

	return fileContent, nil
}
