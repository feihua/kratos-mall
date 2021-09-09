package ums

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type LoginLogListReq struct {
	Current  int64
	PageSize int64
}

type LoginLog struct {
	Id         int64
	MemberId   int64
	CreateTime time.Time
	Ip         string
	City       string
	LoginType  int // 登录类型：0->PC；1->android;2->ios;3->小程序
	Province   string
}

type LoginLogRepo interface {
	CreateLoginLog(context.Context, *LoginLog) error
	GetLoginLog(ctx context.Context, id int64) error
	UpdateLoginLog(context.Context, *LoginLog) error
	ListLoginLog(ctx context.Context, req *LoginLogListReq) ([]*LoginLog, error)
	DeleteLoginLog(ctx context.Context, id int64) error
}

type LoginLogUseCase struct {
	repo LoginLogRepo
	log  *log.Helper
}

func NewLoginLogUseCase(repo LoginLogRepo, logger log.Logger) *LoginLogUseCase {
	return &LoginLogUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (r *LoginLogUseCase) CreateLoginLog(ctx context.Context, user *LoginLog) error {
	panic("implement me")
}

func (r *LoginLogUseCase) GetLoginLog(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r *LoginLogUseCase) UpdateLoginLog(ctx context.Context, user *LoginLog) error {
	panic("implement me")
}

func (r *LoginLogUseCase) ListLoginLog(ctx context.Context, pageNum, pageSize int64) ([]*LoginLog, error) {
	panic("implement me")
}

func (r *LoginLogUseCase) DeleteLoginLog(ctx context.Context, id int64) error {
	panic("implement me")
}
