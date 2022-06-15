package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// var (
// 	// ErrUserNotFound is user not found.
// 	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
// )

// Repo is a Greater repo.
// type ArticleRepo interface{}
// type CommentRepo interface{}
// type TagRepo interface{}

// GreeterUsecase is a Greeter usecase.
type SocialUsecase struct {
	// ar ArticleRepo
	// cr CommentRepo
	// tr TagRepo

	log *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
// func NewSocialUsecase(ar ArticleRepo, cr CommentRepo, tr TagRepo, logger log.Logger) *SocialUsecase {
// 	return &SocialUsecase{ar: ar, cr: cr, tr: tr, log: log.NewHelper(logger)}
// }
func NewSocialUsecase(logger log.Logger) *SocialUsecase {
	return &SocialUsecase{log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *SocialUsecase) CreateArticle(ctx context.Context) error {
	return nil
}
