# Spy Cat Agency (SCA) CRUD Application

## Overview
This project is a robust CRUD application for managing operations at the Spy Cat Agency (SCA). It showcases building RESTful APIs in Go, interacting with PostgreSQL using sqlx, integrating third-party services, and optionally creating user interfaces.

## Requirements
Ensure you have the following dependencies installed:
- go: 1.22 version
- [cleanenv](https://github.com/ilyakaznacheev/cleanenv): Reads environment variables into Go structs.
- [godotenv](https://github.com/joho/godotenv): Loads environment variables from a `.env` file.
- [zap](https://go.uber.org/zap): High-performance logging library.
- [sqlx](https://github.com/jmoiron/sqlx): Database library with extensions for Go.
- [pq](https://github.com/lib/pq): PostgreSQL driver for Go.
- [goose](https://github.com/pressly/goose): CLI tool for managing Go migrations.

## Installation
1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd <repository-directory>
   ```

2. **Install dependencies:**
   ```bash
   go get -u github.com/ilyakaznacheev/cleanenv
   go get github.com/joho/godotenv
   go get -u go.uber.org/zap
   go get github.com/jmoiron/sqlx
   go get github.com/lib/pq
   go get -u github.com/pressly/goose/cmd/goose
   ```

3. **Setup environment variables:**
   Create a `.env` file with your parameters:
   ```dotenv
   SERVICE_HOST=
   SERVICE_PORT=
   SERVICE_READ_TIMEOUT_SECONDS=
   SERVICE_WRITE_TIMEOUT_SECONDS=
   SERVICE_MAX_HEADER_BYTES=
  
   POSTGRES_HOST=
   POSTGRES_PORT=
   POSTGRES_USERNAME=
   POSTGRES_PASSWORD=
   POSTGRES_DBNAME=
   POSTGRES_SSL=
   ```

4. **Run PostgreSQL in Docker:**
   Example setup:
   ```bash
   sudo docker run --name=sca-db -e POSTGRES_PASSWORD=qwerty -p 5436:5432 -d postgres
   sudo docker exec -it <postgresql_container_id> /bin/bash
   psql -U postgres
   CREATE DATABASE sca;
   ```

5. **Apply database migrations:**
   ```bash
   goose -dir migrations postgres "user=postgres dbname=postgres sslmode=disable host=localhost port=5436 password=qwerty" up
   ```

## Usage
- Modify the `.env` file with your specific configurations.
- Start the application:
  ```bash
  make build
  make run
  ```

## Tests
- Find postman tests in postman_tests

