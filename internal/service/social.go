package service

import (
	"context"
	pb "realworld/api/helloworld/v1"
)

func (s *RealWorldService) Listarticles(ctx context.Context, req *pb.ListArticlesRequest) (*pb.MultipleArticlesReply, error) {
	return &pb.MultipleArticlesReply{}, nil
}
func (s *RealWorldService) Feedarticles(ctx context.Context, req *pb.FeedArticlesRequest) (*pb.MultipleArticlesReply, error) {
	return &pb.MultipleArticlesReply{}, nil
}
func (s *RealWorldService) Getarticles(ctx context.Context, req *pb.GetArticleRequest) (*pb.MultipleArticlesReply, error) {
	return &pb.MultipleArticlesReply{}, nil
}
func (s *RealWorldService) Createarticles(ctx context.Context, req *pb.CreateArticleRequest) (*pb.SingleArticleReply, error) {
	return &pb.SingleArticleReply{}, nil
}
func (s *RealWorldService) Updatearticles(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.SingleArticleReply, error) {
	return &pb.SingleArticleReply{}, nil
}
func (s *RealWorldService) Deletearticles(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.SingleArticleReply, error) {
	return &pb.SingleArticleReply{}, nil
}
func (s *RealWorldService) Addcomment(ctx context.Context, req *pb.AddCommentRequest) (*pb.SingleCommentReply, error) {
	return &pb.SingleCommentReply{}, nil
}
func (s *RealWorldService) Getcomment(ctx context.Context, req *pb.GetCommentRequest) (*pb.MultipleCommentsReply, error) {
	return &pb.MultipleCommentsReply{}, nil
}
func (s *RealWorldService) Deletecomment(ctx context.Context, req *pb.DeleteCommentRequest) (*pb.SingleCommentReply, error) {
	return &pb.SingleCommentReply{}, nil
}
func (s *RealWorldService) Favoritearticle(ctx context.Context, req *pb.FavoriteArticleRequest) (*pb.SingleArticleReply, error) {
	return &pb.SingleArticleReply{}, nil
}
func (s *RealWorldService) Unfavoritearticle(ctx context.Context, req *pb.UnFavoriteArticleRequest) (*pb.SingleArticleReply, error) {
	return &pb.SingleArticleReply{}, nil
}
func (s *RealWorldService) Gettags(ctx context.Context, req *pb.GetTagsRequest) (*pb.ListTagsReply, error) {
	return &pb.ListTagsReply{}, nil
}
func (s *RealWorldService) Getprofile(ctx context.Context, req *pb.GetProfileRequest) (*pb.ProfileReply, error) {
	return &pb.ProfileReply{}, nil
}
