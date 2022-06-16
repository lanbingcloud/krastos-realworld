package service

import (
	"context"
	pb "realworld/api/helloworld/v1"
)

func (s *RealWorldService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.UserReply, error) {
	email := req.User.Email
	return &pb.UserReply{
		User: &pb.UserReply_User{
			Username: "bool",
			Email:    email,
		},
	}, nil
}
func (s *RealWorldService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.UserReply, error) {
	u, err := s.uc.Register(ctx, req.User.Username, req.User.Email, req.User.Password)
	if err != nil {
		return nil, err
	}
	return &pb.UserReply{
		User: &pb.UserReply_User{
			Username: u.Username,
			Email:    u.Email,
			Token:    u.Token,
		},
	}, nil
}
func (s *RealWorldService) Getcurrentuser(ctx context.Context, req *pb.GetCurrentUserRequest) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}
func (s *RealWorldService) Updateuser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}
func (s *RealWorldService) Followuser(ctx context.Context, req *pb.FollowUserRequest) (*pb.ProfileReply, error) {
	return &pb.ProfileReply{}, nil
}
func (s *RealWorldService) Unfollowuser(ctx context.Context, req *pb.UnFollowUserRequest) (*pb.ProfileReply, error) {
	return &pb.ProfileReply{}, nil
}
