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

run: build
	go run cmd/main.go