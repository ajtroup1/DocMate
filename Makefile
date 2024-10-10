# Variables
BINARY_NAME=docmate
SOURCE_DIR=./cmd/
OUTPUT_DIR=./bin

# Default target: build the project
build:
	@echo "Building the project..."
	go build -o $(OUTPUT_DIR)/$(BINARY_NAME) $(SOURCE_DIR)
	@echo "Build successful! Binary located at $(OUTPUT_DIR)/$(BINARY_NAME)"

# Clean target: remove the binary
clean:
	@echo "Cleaning up..."
	rm -f $(OUTPUT_DIR)/$(BINARY_NAME)
	@echo "Clean completed!"

# Run target: build and run the project
run: build
	@echo "Running the project..."
	$(OUTPUT_DIR)/$(BINARY_NAME) /path/to/go-project

fmt:
	@echo "Formatting the project..."
	go fmt ./...

# Help target: display available commands
help:
	@echo "Available commands:"
	@echo "  make build     Build the project"
	@echo "  make clean     Clean the project"
	@echo "  make run       Build and run the project"
	@echo "  make help      Display this help message"
