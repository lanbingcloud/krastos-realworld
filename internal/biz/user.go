package biz

import (
	"context"
	"errors"
	"fmt"

	"realworld/internal/conf"
	"realworld/internal/pkg/middlewire/auth"

	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email        string
	Username     string
	Bio          string
	Image        string
	PasswordHash string
}

type UserLogin struct {
	Email    string
	Username string
	Token    string
	Bio      string
	Image    string
}

func hashPassword(pwd string) string {
	b, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", b)
	return string(b)
}

func verifyPassword(hashed, pwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pwd)); err != nil {
		return false
	}
	return true
}

// Repo is a User repo.
type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type ProgfileRepo interface{}

// UserUsecase is a User usecase.
type UserUsecase struct {
	ur   UserRepo
	pr   ProgfileRepo
	jwtc *conf.JWT

	log *log.Helper
}

// NewUserUsecase new a User usecase.
func NewUserUsecase(ur UserRepo, pr ProgfileRepo, jwtc *conf.JWT, logger log.Logger) *UserUsecase {
	return &UserUsecase{ur: ur, pr: pr, jwtc: jwtc, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) generateToken(username string) string {
	return auth.GenerateToken(string(uc.jwtc.Secret), username)
}

// CreateUser creates a User, and returns the new User.
func (uc *UserUsecase) Register(ctx context.Context, email, username, password string) (*UserLogin, error) {
	// u := &User{
	// 	Email:        email,
	// 	Username:     username,
	// 	PasswordHash: hashPassword(password),
	// }
	// if err := uc.ur.CreateUser(ctx, u); err != nil {
	// 	return nil, err
	// }
	return &UserLogin{
		Email:    email,
		Username: username,
		Token:    uc.generateToken(username),
	}, nil
}

func (uc *UserUsecase) Login(ctx context.Context, email string, password string) (*UserLogin, error) {
	u, err := uc.ur.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if verifyPassword(u.PasswordHash, password) {
		return nil, errors.New("login fail")
	}
	return &UserLogin{
		Email:    u.Email,
		Username: u.Username,
		Bio:      u.Bio,
		Image:    u.Image,
		Token:    uc.generateToken(u.Username),
	}, nil
}
