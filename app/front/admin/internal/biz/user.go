package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	sysV1 "kratos-mall/api/sys/v1"
)

type Login struct {
	UserName string
	Password string
}

type UserRepo interface {
	Login(context.Context, *Login) (*sysV1.LoginResp, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUseCase) Login(ctx context.Context, g *Login) (*sysV1.LoginResp, error) {
	return uc.repo.Login(ctx, g)
}
