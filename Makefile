include .env
export $(shell sed 's/=.*//g' .env)

up:
	go run cmd/server/main.go

migrate-up:
	goose -dir migrations postgres "postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)" up

migrate-down:
	goose -dir migrations postgres "postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)" reset

swagger:
	swag init --parseDependency --parseInternal -g cmd/server/main.go

dbuild:
	docker build -t image-optimization-api .