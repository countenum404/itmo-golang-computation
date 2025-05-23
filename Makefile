
NAME=app

image:
	docker build -t $(NAME) .

run-image:
	docker run --rm -p 8080:8080 $(NAME)

up:
	docker-compose up

clean:
	rm -rf test-reports
	docker compose rm -f -s
	docker rmi itmo-golang-computation-2025-grpc:latest itmo-golang-computation-2025-http:latest

protoc:
	protoc --go_out=. --go-grpc_out=. proto/calculation.proto

swagger-doc:
	cd internal/handlers/
	swag init --parseDependency --parseInternal -g handlers.go

deps:
	go mod download

build:
	go build -o bin/app ./cmd/main.go

test:
	mkdir -p test-reports
	go test -race -coverprofile=./test-reports/coverage.txt -covermode=atomic ./...

lint:
	@echo "Running linters..."
	go vet $(shell go list ./...)
	staticcheck $(shell go list ./...) || true
	golint $(shell go list ./...) || true