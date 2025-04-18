
NAME=app

image:
	docker build -t $(NAME) .

run-image:
	docker run --rm -p 8080:8080 $(NAME)

up:
	docker-compose up

clean:
	docker compose rm -f -s
	docker rmi itmo-golang-computation-2025-grpc:latest itmo-golang-computation-2025-http:latest

protoc:
	protoc --go_out=. --go-grpc_out=. proto/calculation.proto

swagger-doc:
	cd internal/handlers/
	swag init --parseDependency --parseInternal -g handlers.go