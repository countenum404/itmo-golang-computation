# ITMO Golang Computation
## Author
📝 Denis Shabashov  
Group: K4111c  
Email: countenum404@gmail.com or d.shabashov@icloud.com

## About
This repository contains an example application written in Go that demonstrates how to implement a compution using gRPC and HTTP servers. The application provides two ways to calculate mathematical expressions: synchronous execution via BasicSolver and asynchronous execution via AsyncSolver.

## Features

* Support for basic arithmetic operations like addition, subtraction, multiplication, and division.  
* Two solvers available: BasicSolver (synchronous) and AsyncSolver (asynchronous).  
* REST API endpoint accessible through HTTP requests.  
* GRPC endpoint for remote procedure calls.  
* Swagger documentation generated automatically using Swag library.
* Advanced DI that uber.fx package provides

## Getting Started
**Prerequisites**  
Before getting started, ensure you have the following installed:

Go: Version 1.23+ Docker: For containerization and local development Protobuf compiler: To generate .pb.go files Swag: For generating Swagger documentation
## Installation

Clone this repository:

``` bash
git clone git@github.com:countenum404/itmo-golang-computation.git  
cd itmo-golang-computation
```

## Building and Running Locally

Use the provided Makefile to simplify common tasks:

Build the application:  
`make build`

Run unit tests:  
`make test`

Run static analysis tools:  
`make lint`

Clean up containers and images:  
`make clean`

Start the application using Docker Compose:  
`make up`

## View Swagger UI documentation:  
Open your browser and navigate to http://localhost:8080/swagger/index.html.  

## Curl examples
sample:
```bash
    curl -X POST \
  http://localhost:8080/api \
  -H 'Content-Type: application/json' \
  -d '{
    "operations": [
      {"type": "calc", "op": "+", "var": "sum", "left": "1", "right": "2"},
      {"type": "print", "var": "sum"}
    ]
  }'
```

sample:
```bash
curl --location 'http://localhost:8080/api' \
-H 'Content-Type: application/json' \
-d '{
    "operations": [
        { "type": "calc", "op": "+", "var": "x", "left": "10", "right": "2" },
        { "type": "calc", "op": "*", "var": "y", "left": "x", "right": "5" },
        { "type": "calc", "op": "-", "var": "q", "left": "y", "right": "20" },
        { "type": "calc", "op": "+", "var": "unusedA", "left": "y", "right": "100" },
        { "type": "calc", "op": "*", "var": "unusedB", "left": "unusedA", "right": "2" },
        { "type": "print", "var": "q" },
        { "type": "calc", "op": "-", "var": "z", "left": "x", "right": "15" },
        { "type": "print", "var": "z" },
        { "type": "calc", "op": "+", "var": "ignoreC", "left": "z", "right": "y" },
        { "type": "print", "var": "x" }
    ]
}'

```