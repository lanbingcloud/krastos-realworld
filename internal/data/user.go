package data

import (
	"context"

	"realworld/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) error {
	r.data.db.Create(u)
	return nil
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (*biz.User, error) {
	r.data.db.Get("")
	return nil, nil
}

type profileRepo struct {
	data *Data
	log  *log.Helper
}

// NewProfileRepo .
func NewProfileRepo(data *Data, logger log.Logger) biz.ProgfileRepo {
	return &profileRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// func (r *greeterRepo) Update(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
// 	return g, nil
// }

// func (r *greeterRepo) FindByID(context.Context, int64) (*biz.Greeter, error) {
// 	return nil, nil
// }

// func (r *greeterRepo) ListByHello(context.Context, string) ([]*biz.Greeter, error) {
// 	return nil, nil
// }

// func (r *greeterRepo) ListAll(context.Context) ([]*biz.Greeter, error) {
// 	return nil, nil
// }
