MIGRATION_DIR=./db/migrations

create-migration:
	migrate create -ext sql -dir $(MIGRATION_DIR) -seq create_schema

migrateup:
	go run ./db/main.go up

migratedown:
	go run ./db/main.go down

build:
	go build -o bin/backend-guchitter-app -v .

build-migration:
	go build -o bin/migration -v ./db