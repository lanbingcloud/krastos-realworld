package service

import (
	"github.com/google/wire"

	v1 "realworld/api/helloworld/v1"
	"realworld/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealWorldService)

// RealWorldService is a greeter service.
type RealWorldService struct {
	v1.UnimplementedRealWorldServer

	sc *biz.SocialUsecase
	uc *biz.UserUsecase
}

// NewRealWorldService new a greeter service.
func NewRealWorldService(sc *biz.SocialUsecase, uc *biz.UserUsecase, logger log.Logger) *RealWorldService {
	return &RealWorldService{sc: sc, uc: uc}
}
