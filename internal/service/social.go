package service

import (
	"context"
	pb "realworld/api/helloworld/v1"
)

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
func (s *RealWorldService) Getprofile(ctx context.Context, req *pb.GetProfileRequest) (*pb.ProfileReplay, error) {
	return &pb.ProfileReplay{}, nil
}
