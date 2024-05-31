# Use the official Golang image to build the app
FROM golang:1.18 as builder

WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o /grpc-user-service server/main.go server/user_service.go server/gateway.go

# Use a smaller image for the final build
FROM alpine:latest

WORKDIR /root/

# Copy the binary from the builder image
COPY --from=builder /grpc-user-service .

# Expose ports for gRPC and HTTP
EXPOSE 50051 8080

# Command to run the binary
CMD ["./grpc-user-service"]
