#build:
#	@go build -o bin/com .\cmd\main.go
#
#test:
#	@go test -v ./...
#
#run: build
#	@./bin/ecom

build:
	go build -o bin/com cmd/main.go

test:
	go test -v ./...

run: build
	go run cmd/main.go

migration:
	migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	go run cmd/migrate/main.go up

migrate-down:
	go run cmd/migrate/main.go down