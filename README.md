# Distributed File Storage Server

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/tushar/Assignment-HyperHire.git
   cd Assignment-HyperHire
   ```
## To access the docker-compose and docker file 
2. cd cmd/fileStorageService/


3. Build and run the application using Docker Compose:
   ```sh
   docker-compose up --build
   ```

## Running the Application

Once the application is running, it will be available at `http://localhost:8080`.


## API Endpoints

- **Upload File**: `POST /upload`
  - Uploads a file and returns a unique file ID.
  - Example:
    ```sh
    curl -F "file=@/path/to/your/file" http://localhost:8080/upload
    ```

- **Get Uploaded Files**: `GET /files`
  - Returns a list of all uploaded file IDs.
  - Example:
    ```sh
    curl http://localhost:8080/files
    ```

- **Download File**: `GET /download?id=<fileID>`
  - Downloads the file with the given file ID.
  - Example:
    ```sh
    curl -O http://localhost:8080/download?id=<fileID>
    ```
