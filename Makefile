APP_NAME := app

OUTPUT_DIR := /.bin

build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/app/main.go

run: build
	@echo "Running the application..."
	./$(OUTPUT_DIR)/$(APP_NAME)

clean:
	@echo "Cleaning up..."
	rm -f $(OUTPUT_DIR)/$(APP_NAME)

deps:
	@echo "Getting dependencies..."
	go mod tidy

.PHONY: build run clean test deps help
