# gRPC User Service

This is a gRPC-based User Service implemented in Go. It allows you to get user details, list users by IDs, and search for users based on criteria.

## Prerequisites

- Go 1.18 or higher
- Docker (if running in a container)

## Running the Server

```sh
go run server/main.go

## Docker
- docker build -t grpc-user-service .
- docker run -p 50051:50051 -p 8080:8080 grpc-user-service
