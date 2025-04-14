FROM golang:1.23.8-alpine3.21 AS build

WORKDIR /home/build

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/main.go

FROM scratch

WORKDIR /opt/backend-service

COPY --from=build /home/build/app app

# if SERVER_TYPE is http
EXPOSE 8080
# if SERVER_TYPE is grpc
EXPOSE 50051

CMD [ "./app" ]