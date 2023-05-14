MIGRATION_DIR=./db/migrations

run:
	swag init && go mod tidy && go mod vendor && go run .

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

gen-test:
	gotests -template_dir="util/testUtils/templates/custom"  -w -all $(ARG)