
NAME=app

image:
	docker build -t $(NAME) .

run-image:
	docker run --rm -p 8080:8080 $(NAME)

up:
	docker-compose up

clean:
	rm -rf pkg/app
	docker compose rm -f -s && docker rmi $(NAME)

protoc:
	protoc --go_out=. --go-grpc_out=. proto/calculation.proto