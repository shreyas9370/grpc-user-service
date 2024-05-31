package main

import (
	"fmt"
	pb "grpc-user-service/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	users := []pb.User{
		{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
		{Id: 2, Fname: "John", City: "NY", Phone: 9876543210, Height: 6.0, Married: false},
		// Add more users as needed
	}

	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, &UserServiceServer{users: users})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Println("Server is running on port 50051")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
