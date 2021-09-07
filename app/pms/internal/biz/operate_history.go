package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type OperateHistoryListReq struct {
	Current  int64
	PageSize int64
}

type OperateHistory struct {
	Id               int64
	ProductId        int64
	PriceOld         string
	PriceNew         string
	SalePriceOld     string
	SalePriceNew     string
	GiftPointOld     int
	GiftPointNew     int
	UsePointLimitOld int
	UsePointLimitNew int
	OperateMan       string
	CreateTime       time.Time
}

type OperateHistoryRepo interface {
	CreateOperateHistory(context.Context, *OperateHistory) error
	GetOperateHistory(ctx context.Context, id int64) error
	UpdateOperateHistory(context.Context, *OperateHistory) error
	ListOperateHistory(ctx context.Context, req *OperateHistoryListReq) ([]*OperateHistory, error)
	DeleteOperateHistory(ctx context.Context, id int64) error
}

type OperateHistoryUseCase struct {
	repo OperateHistoryRepo
	log  *log.Helper
}

func NewOperateHistoryUseCase(repo OperateHistoryRepo, logger log.Logger) *OperateHistoryUseCase {
	return &OperateHistoryUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (r *OperateHistoryUseCase) CreateOperateHistory(ctx context.Context, user *OperateHistory) error {
	panic("implement me")
}

func (r *OperateHistoryUseCase) GetOperateHistory(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r *OperateHistoryUseCase) UpdateOperateHistory(ctx context.Context, user *OperateHistory) error {
	panic("implement me")
}

func (r *OperateHistoryUseCase) ListOperateHistory(ctx context.Context, pageNum, pageSize int64) ([]*OperateHistory, error) {
	panic("implement me")
}

func (r *OperateHistoryUseCase) DeleteOperateHistory(ctx context.Context, id int64) error {
	panic("implement me")
}
