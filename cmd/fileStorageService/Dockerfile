# Start with the official Golang image
FROM golang:1.23-alpine

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go module files first, then download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app
RUN go build -o main ./cmd/fileStorageService/main.go

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
