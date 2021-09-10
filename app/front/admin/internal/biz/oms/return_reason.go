package oms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type ReturnReasonListReq struct {
	Current  int64
	PageSize int64
}

type ReturnReason struct {
	Id         int64
	Name       string // 退货类型
	Sort       int
	Status     int    // 状态：0->不启用；1->启用
	CreateTime string // 添加时间
}

type ReturnReasonRepo interface {
	CreateReturnReason(context.Context, *ReturnReason) error
	GetReturnReason(ctx context.Context, id int64) error
	UpdateReturnReason(context.Context, *ReturnReason) error
	ListReturnReason(ctx context.Context, req *ReturnReasonListReq) ([]*ReturnReason, error)
	DeleteReturnReason(ctx context.Context, id int64) error
}

type ReturnReasonUseCase struct {
	repo ReturnReasonRepo
	log  *log.Helper
}

func NewReturnReasonUseCase(repo ReturnReasonRepo, logger log.Logger) *ReturnReasonUseCase {
	return &ReturnReasonUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (r *ReturnReasonUseCase) CreateReturnReason(ctx context.Context, user *ReturnReason) error {
	panic("implement me")
}

func (r *ReturnReasonUseCase) GetReturnReason(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r *ReturnReasonUseCase) UpdateReturnReason(ctx context.Context, user *ReturnReason) error {
	panic("implement me")
}

func (r *ReturnReasonUseCase) ListReturnReason(ctx context.Context, pageNum, pageSize int64) ([]*ReturnReason, error) {
	panic("implement me")
}

func (r *ReturnReasonUseCase) DeleteReturnReason(ctx context.Context, id int64) error {
	panic("implement me")
}
