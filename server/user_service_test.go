package main

import (
	"context"
	"testing"

	pb "grpc-user-service/proto"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	server := &UserServiceServer{
		users: []pb.User{
			{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
		},
	}

	req := &pb.GetUserRequest{Id: 1}
	res, err := server.GetUser(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, int32(1), res.User.Id)
}

func TestListUsers(t *testing.T) {
	server := &UserServiceServer{
		users: []pb.User{
			{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
			{Id: 2, Fname: "John", City: "NY", Phone: 9876543210, Height: 6.0, Married: false},
		},
	}

	req := &pb.ListUsersRequest{Ids: []int32{1, 2}}
	res, err := server.ListUsers(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, 2, len(res.Users))
}

func TestSearchUsers(t *testing.T) {
	server := &UserServiceServer{
		users: []pb.User{
			{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
			{Id: 2, Fname: "John", City: "NY", Phone: 9876543210, Height: 6.0, Married: false},
		},
	}

	req := &pb.SearchUsersRequest{City: "LA"}
	res, err := server.SearchUsers(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, 1, len(res.Users))
}
