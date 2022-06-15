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
func (s *RealWorldService) Getprofile(ctx context.Context, req *pb.GetProfileRequest) (*pb.ProfileReplay, error) {
	return &pb.ProfileReplay{}, nil
}
func (s *RealWorldService) Followuser(ctx context.Context, req *pb.FollowUserRequest) (*pb.ProfileReplay, error) {
	return &pb.ProfileReplay{}, nil
}
func (s *RealWorldService) Unfollowuser(ctx context.Context, req *pb.UnFollowUserRequest) (*pb.ProfileReplay, error) {
	return &pb.ProfileReplay{}, nil
}
func (s *RealWorldService) Listarticles(ctx context.Context, req *pb.ListArticlesRequest) (*pb.MultipleArticlesReplay, error) {
	return &pb.MultipleArticlesReplay{}, nil
}
func (s *RealWorldService) Feedarticles(ctx context.Context, req *pb.FeedArticlesRequest) (*pb.MultipleArticlesReplay, error) {
	return &pb.MultipleArticlesReplay{}, nil
}
func (s *RealWorldService) Getarticles(ctx context.Context, req *pb.GetArticleRequest) (*pb.MultipleArticlesReplay, error) {
	return &pb.MultipleArticlesReplay{}, nil
}
func (s *RealWorldService) Createarticles(ctx context.Context, req *pb.CreateArticleRequest) (*pb.SingleArticleReplay, error) {
	return &pb.SingleArticleReplay{}, nil
}
func (s *RealWorldService) Updatearticles(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.SingleArticleReplay, error) {
	return &pb.SingleArticleReplay{}, nil
}
func (s *RealWorldService) Deletearticles(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.SingleArticleReplay, error) {
	return &pb.SingleArticleReplay{}, nil
}
func (s *RealWorldService) Addcomment(ctx context.Context, req *pb.AddCommentRequest) (*pb.SingleCommentReplay, error) {
	return &pb.SingleCommentReplay{}, nil
}
func (s *RealWorldService) Getcomment(ctx context.Context, req *pb.GetCommentRequest) (*pb.MultipleCommentsReplay, error) {
	return &pb.MultipleCommentsReplay{}, nil
}
func (s *RealWorldService) Deletecomment(ctx context.Context, req *pb.DeleteCommentRequest) (*pb.SingleCommentReplay, error) {
	return &pb.SingleCommentReplay{}, nil
}
func (s *RealWorldService) Favoritearticle(ctx context.Context, req *pb.FavoriteArticleRequest) (*pb.SingleArticleReplay, error) {
	return &pb.SingleArticleReplay{}, nil
}
func (s *RealWorldService) Unfavoritearticle(ctx context.Context, req *pb.UnFavoriteArticleRequest) (*pb.SingleArticleReplay, error) {
	return &pb.SingleArticleReplay{}, nil
}
func (s *RealWorldService) Gettags(ctx context.Context, req *pb.GetTagsRequest) (*pb.ListTagsReplay, error) {
	return &pb.ListTagsReplay{}, nil
}
