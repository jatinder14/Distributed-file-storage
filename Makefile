# Makefile for Distributed File Storage Service

# Variables
DOCKER_COMPOSE_FILE = cmd/fileStorageService/docker-compose.yml
SERVICE_NAME = file_storage_service
DB_CONTAINER = postgres_db
GO_MAIN = cmd/fileStorageService/main.go

# Default target
.PHONY: all
all: build

# Build the Go application
.PHONY: build
build:
	@echo "Building the Go application..."
	docker-compose -f cmd/fileStorageService/docker-compose.yml build

# Start the services using Docker Compose
.PHONY: up
up:
	@echo "Starting the service..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) up

# Stop the services
.PHONY: down
down:
	@echo "Stopping the service..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) down


# Clean up Docker containers, volumes, and images
.PHONY: clean
clean:
	@echo "Cleaning up Docker containers and images..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) down --rmi all --volumes --remove-orphans

# Rebuild and restart the service
.PHONY: rebuild
rebuild: down build up
