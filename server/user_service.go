package main

import (
	"context"
	"errors"
	pb "grpc-user-service/proto"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
	users []pb.User
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	for _, user := range s.users {
		if user.Id == req.Id {
			return &pb.GetUserResponse{User: &user}, nil
		}
	}
	return nil, errors.New("user not found")
}

func (s *UserServiceServer) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	var users []*pb.User
	for _, id := range req.Ids {
		for _, user := range s.users {
			if user.Id == id {
				users = append(users, &user)
			}
		}
	}
	return &pb.ListUsersResponse{Users: users}, nil
}

func (s *UserServiceServer) SearchUsers(ctx context.Context, req *pb.SearchUsersRequest) (*pb.SearchUsersResponse, error) {
	var users []*pb.User
	for _, user := range s.users {
		if (req.City == "" || user.City == req.City) &&
			(req.Phone == 0 || user.Phone == req.Phone) &&
			(req.Married == user.Married) {
			users = append(users, &user)
		}
	}
	return &pb.SearchUsersResponse{Users: users}, nil
}
