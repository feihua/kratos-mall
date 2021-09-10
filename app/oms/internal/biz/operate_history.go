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
	Id          int64
	OrderId     int64  // 订单id
	OperateMan  string // 操作人：用户；系统；后台管理员
	CreateTime  string // 操作时间
	OrderStatus int    // 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
	Note        string // 备注
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
