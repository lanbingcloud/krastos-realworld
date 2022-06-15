package service

import (
	"context"
	pb "realworld/api/helloworld/v1"
)

func (s *RealWorldService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.UserReplay, error) {
	email := req.User.Email
	return &pb.UserReplay{
		User: &pb.UserReplay_User{
			Username: "bool",
			Email:    email,
		},
	}, nil
}
func (s *RealWorldService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.UserReplay, error) {
	return &pb.UserReplay{}, nil
}
func (s *RealWorldService) Getcurrentuser(ctx context.Context, req *pb.GetCurrentUserRequest) (*pb.UserReplay, error) {
	return &pb.UserReplay{}, nil
}
func (s *RealWorldService) Updateuser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserReplay, error) {
	return &pb.UserReplay{}, nil
}
func (s *RealWorldService) Followuser(ctx context.Context, req *pb.FollowUserRequest) (*pb.ProfileReplay, error) {
	return &pb.ProfileReplay{}, nil
}
func (s *RealWorldService) Unfollowuser(ctx context.Context, req *pb.UnFollowUserRequest) (*pb.ProfileReplay, error) {
	return &pb.ProfileReplay{}, nil
}
