package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
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
	CreateTime       string
}
type OperateHistoryListResp struct {
	Total int64
	List  []*OperateHistory
}

type OperateHistoryRepo interface {
	CreateOperateHistory(context.Context, *OperateHistory) error
	GetOperateHistory(ctx context.Context, id int64) error
	UpdateOperateHistory(context.Context, *OperateHistory) error
	ListOperateHistory(ctx context.Context, req *OperateHistoryListReq) (*OperateHistoryListResp, error)
	DeleteOperateHistory(ctx context.Context, id int64) error
}

type OperateHistoryUseCase struct {
	repo OperateHistoryRepo
	log  *log.Helper
}

func NewOperateHistoryUseCase(repo OperateHistoryRepo, logger log.Logger) *OperateHistoryUseCase {
	return &OperateHistoryUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (o OperateHistoryUseCase) CreateOperateHistory(ctx context.Context, history *OperateHistory) error {
	panic("implement me")
}

func (o OperateHistoryUseCase) GetOperateHistory(ctx context.Context, id int64) error {
	panic("implement me")
}

func (o OperateHistoryUseCase) UpdateOperateHistory(ctx context.Context, history *OperateHistory) error {
	panic("implement me")
}

func (o OperateHistoryUseCase) ListOperateHistory(ctx context.Context, req *OperateHistoryListReq) (*OperateHistoryListResp, error) {
	return o.repo.ListOperateHistory(ctx, req)
}

func (o OperateHistoryUseCase) DeleteOperateHistory(ctx context.Context, id int64) error {
	panic("implement me")
}
