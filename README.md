# Spy Cat Agency (SCA) CRUD Application

## Overview

The Spy Cat Agency (SCA) CRUD Application is a comprehensive system for managing operations at the Spy Cat Agency. This project is designed to demonstrate how to build RESTful APIs in Go, interact with PostgreSQL using sqlx, integrate third-party services, and optionally build user interfaces.

## Features
- **RESTful API**: Exposes endpoints for CRUD operations on the Spy Cat database.
- **PostgreSQL Integration**: Uses sqlx to interact with a PostgreSQL database.
- **Migrations**: Manage database schema changes using goose.
- **Swagger API Documentation**: Automatically generated documentation for the API using swag.
- **Docker Support**: Containerized application for ease of deployment and scalability.

## Getting Started

### Environment Variables

Create a `.env` file with the following configuration:

```dotenv
HTTP_PORT=8080

DB_USERNAME=postgres
DB_PASSWORD=qwerty
DB_HOST=localhost
DB_PORT=5432
DB_NAME=spy_cats
DB_SSLMODE=disable
```

To load environment variables, include this in your shell configuration or use the export command:

export $(shell sed 's/=.*//g' .env)

Running the Application
	1.	Start the application:

make up

This will start the server using go run cmd/server/main.go.

	2.	Run Database Migrations:
To apply database migrations, use the following command:

make migrate-up

To reset the migrations:

make migrate-down


	3.	Swagger Documentation:
To generate Swagger API documentation, use:

make swagger

This will generate the necessary Swagger files for API documentation.

	4.	Build Docker Image (Optional):
To build the Docker image for the application:

make dbuild



Postman Collection

You can import the SpyCat API.postman_collection.json into Postman to test the API endpoints.

Dependencies

The application uses the following key dependencies:
	•	github.com/jmoiron/sqlx: SQL library for Go.
	•	github.com/lib/pq: PostgreSQL driver for Go.
	•	github.com/spf13/viper: Configuration management library.
	•	go.uber.org/zap: Structured logging library.
	•	github.com/pressly/goose: Database migration tool.