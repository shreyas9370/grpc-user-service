package main

import (
	"context"
	"fmt"
	"log"

	pb "grpc-user-service/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	// Call GetUser
	userReq := &pb.GetUserRequest{Id: 1}
	userRes, err := client.GetUser(context.Background(), userReq)
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}
	fmt.Printf("User: %v\n", userRes.User)

	// Call ListUsers
	listReq := &pb.ListUsersRequest{Ids: []int32{1, 2}}
	listRes, err := client.ListUsers(context.Background(), listReq)
	if err != nil {
		log.Fatalf("could not list users: %v", err)
	}
	fmt.Printf("Users: %v\n", listRes.Users)

	// Call SearchUsers
	searchReq := &pb.SearchUsersRequest{City: "LA"}
	searchRes, err := client.SearchUsers(context.Background(), searchReq)
	if err != nil {
		log.Fatalf("could not search users: %v", err)
	}
	fmt.Printf("Search Users: %v\n", searchRes.Users)
}
