package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Hello string
}

type Beer struct {
	Id          int64
	Name        string
	Description string
	Count       int64
}

type UserRepo interface {
	CreateUser(context.Context, *User) error
	GetUser(ctx context.Context, id int64) error
	UpdateUser(context.Context, *User) error
	ListUser(ctx context.Context, pageNum, pageSize int64) ([]*Beer, error)
	DeleteUser(ctx context.Context, id int64) error
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (u *UserUseCase) CreateUser(ctx context.Context, user *User) error {
	panic("implement me")
}

func (u *UserUseCase) GetUser(ctx context.Context, id int64) error {
	panic("implement me")
}

func (u *UserUseCase) UpdateUser(ctx context.Context, user *User) error {
	panic("implement me")
}

func (u *UserUseCase) ListUser(ctx context.Context, pageNum, pageSize int64) ([]*Beer, error) {
	panic("implement me")
}

func (u *UserUseCase) DeleteUser(ctx context.Context, id int64) error {
	panic("implement me")
}
