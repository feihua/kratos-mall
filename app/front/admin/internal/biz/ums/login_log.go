package ums

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type LoginLogListReq struct {
	Current  int64
	PageSize int64
}

type LoginLog struct {
	Id         int64
	MemberId   int64
	CreateTime string
	Ip         string
	City       string
	LoginType  int // 登录类型：0->PC；1->android;2->ios;3->小程序
	Province   string
}
type LoginLogListResp struct {
	Total int64
	List  []*LoginLog
}

type LoginLogRepo interface {
	CreateLoginLog(context.Context, *LoginLog) error
	GetLoginLog(ctx context.Context, id int64) error
	UpdateLoginLog(context.Context, *LoginLog) error
	ListLoginLog(ctx context.Context, req *LoginLogListReq) (*LoginLogListResp, error)
	DeleteLoginLog(ctx context.Context, id int64) error
}

type LoginLogUseCase struct {
	repo LoginLogRepo
	log  *log.Helper
}

func NewLoginLogUseCase(repo LoginLogRepo, logger log.Logger) *LoginLogUseCase {
	return &LoginLogUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (l LoginLogUseCase) CreateLoginLog(ctx context.Context, loginLog *LoginLog) error {
	panic("implement me")
}

func (l LoginLogUseCase) GetLoginLog(ctx context.Context, id int64) error {
	panic("implement me")
}

func (l LoginLogUseCase) UpdateLoginLog(ctx context.Context, loginLog *LoginLog) error {
	panic("implement me")
}

func (l LoginLogUseCase) ListLoginLog(ctx context.Context, req *LoginLogListReq) (*LoginLogListResp, error) {
	return l.repo.ListLoginLog(ctx, req)
}

func (l LoginLogUseCase) DeleteLoginLog(ctx context.Context, id int64) error {
	panic("implement me")
}
