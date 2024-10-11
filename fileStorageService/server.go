package fileStorageService

import (
	"log"
	"net/http"
)

func StartServer() error {
	InitDB()
	http.HandleFunc("/upload", UploadFileHandler)
	http.HandleFunc("/files", GetUploadedFilesHandler)
	http.HandleFunc("/download", DownloadFileHandler)

	log.Println("Starting server on :8080")
	return http.ListenAndServe(":8080", nil)
}
