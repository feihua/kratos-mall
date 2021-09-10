package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type HomeAdvertiseListReq struct {
	Current  int64
	PageSize int64
}

type HomeAdvertise struct {
	Id         int64
	Name       string
	Type       int // 轮播位置：0->PC首页轮播；1->app首页轮播
	Pic        string
	StartTime  string
	EndTime    string
	Status     int    // 上下线状态：0->下线；1->上线
	ClickCount int    // 点击数
	OrderCount int    // 下单数
	Url        string // 链接地址
	Note       string // 备注
	Sort       int    // 排序
}

type HomeAdvertiseRepo interface {
	CreateHomeAdvertise(context.Context, *HomeAdvertise) error
	GetHomeAdvertise(ctx context.Context, id int64) error
	UpdateHomeAdvertise(context.Context, *HomeAdvertise) error
	ListHomeAdvertise(ctx context.Context, req *HomeAdvertiseListReq) ([]*HomeAdvertise, error)
	DeleteHomeAdvertise(ctx context.Context, id int64) error
}

type HomeAdvertiseUseCase struct {
	repo HomeAdvertiseRepo
	log  *log.Helper
}

func NewHomeAdvertiseUseCase(repo HomeAdvertiseRepo, logger log.Logger) *HomeAdvertiseUseCase {
	return &HomeAdvertiseUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (r *HomeAdvertiseUseCase) CreateHomeAdvertise(ctx context.Context, user *HomeAdvertise) error {
	panic("implement me")
}

func (r *HomeAdvertiseUseCase) GetHomeAdvertise(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r *HomeAdvertiseUseCase) UpdateHomeAdvertise(ctx context.Context, user *HomeAdvertise) error {
	panic("implement me")
}

func (r *HomeAdvertiseUseCase) ListHomeAdvertise(ctx context.Context, pageNum, pageSize int64) ([]*HomeAdvertise, error) {
	panic("implement me")
}

func (r *HomeAdvertiseUseCase) DeleteHomeAdvertise(ctx context.Context, id int64) error {
	panic("implement me")
}
