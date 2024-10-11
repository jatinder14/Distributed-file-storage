package fileStorageService

import (
	"encoding/json"
	"net/http"
)

// UploadFileHandler handles file uploads
func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fileID, err := UploadFileService(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fileID))
}

// GetUploadedFilesHandler returns all uploaded files
func GetUploadedFilesHandler(w http.ResponseWriter, r *http.Request) {
	files := GetAllFilesService()
	if files == nil {
		http.Error(w, "Failed to retrieve files", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(files)
}

// DownloadFileHandler handles file downloads by ID
func DownloadFileHandler(w http.ResponseWriter, r *http.Request) {
	fileID := r.URL.Query().Get("id")
	fileContent, err := DownloadFileService(fileID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+fileID)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileContent)
}
