package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type LogListReq struct {
	Current  int64
	PageSize int64
}

type Log struct {
	Id             int64  // 编号
	UserName       string // 用户名
	Status         string // 登录状态（online:在线，登录初始状态，方便统计在线人数；login:退出登录后将online置为login；logout:退出登录）
	Ip             string // IP地址
	CreateBy       string // 创建人
	CreateTime     string // 创建时间
	LastUpdateBy   string // 更新人
	LastUpdateTime string // 更新时间
}

type LogDTO struct {
	UserName string
	Status   string
	Ip       string
	CreateBy string
}

type LogRepo interface {
	CreateLog(context.Context, *LogDTO) error
	GetLog(ctx context.Context, id int64) error
	UpdateLog(context.Context, *LogDTO) error
	ListLog(ctx context.Context, req *LogListReq) ([]*Log, error)
	DeleteLog(ctx context.Context, id int64) error
}

type LogUseCase struct {
	repo LogRepo
	log  *log.Helper
}

func NewLogUseCase(repo LogRepo, logger log.Logger) *LogUseCase {
	return &LogUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (l *LogUseCase) LoginLogAdd(ctx context.Context, user *LogDTO) error {

	return l.repo.CreateLog(ctx, user)

}

func (l *LogUseCase) GetLog(ctx context.Context, id int64) error {
	panic("implement me")
}

func (l *LogUseCase) UpdateLog(ctx context.Context, user *LogDTO) error {
	panic("implement me")
}

func (l *LogUseCase) ListLog(ctx context.Context, pageNum, pageSize int64) ([]*Log, error) {
	panic("implement me")
}

func (l *LogUseCase) DeleteLog(ctx context.Context, id int64) error {
	panic("implement me")
}
