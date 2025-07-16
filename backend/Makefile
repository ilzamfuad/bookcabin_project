.PHONY: seed unit-test migrate
seed:
	go run ./cmd/seed/main.go

unit-test:
	go test ./... -v -coverprofile=coverage.out

migrate:
	go run db/migrate.go