package biz

import (
	"context"
	"errors"
	"fmt"

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

func hashPassword(pwd string) ([]byte, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	fmt.Printf("hashPassword value is %v, %s | ", b, string(b))
	return b, nil
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
	ur UserRepo
	pr ProgfileRepo

	log *log.Helper
}

// NewUserUsecase new a User usecase.
func NewUserUsecase(ur UserRepo, pr ProgfileRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{ur: ur, pr: pr, log: log.NewHelper(logger)}
}

// CreateUser creates a User, and returns the new User.
func (uc *UserUsecase) Register(ctx context.Context, email, username, password string) (*UserLogin, error) {
	h, err := hashPassword(password)
	if err != nil {
		return nil, err
	}
	u := &User{
		Email:        email,
		Username:     username,
		PasswordHash: string(h),
	}
	if err := uc.ur.CreateUser(ctx, u); err != nil {
		return nil, err
	}
	return &UserLogin{
		Email:    email,
		Username: username,
		Token:    "xxx",
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
		Token:    "xxx",
	}, nil
}
