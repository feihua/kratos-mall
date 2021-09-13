package biz

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
type ReturnReasonListResp struct {
	Total int64
	List  []*ReturnReason
}

type ReturnReasonRepo interface {
	CreateReturnReason(context.Context, *ReturnReason) error
	GetReturnReason(ctx context.Context, id int64) error
	UpdateReturnReason(context.Context, *ReturnReason) error
	ListReturnReason(ctx context.Context, req *ReturnReasonListReq) (*ReturnReasonListResp, error)
	DeleteReturnReason(ctx context.Context, id int64) error
}

type ReturnReasonUseCase struct {
	repo ReturnReasonRepo
	log  *log.Helper
}

func NewReturnReasonUseCase(repo ReturnReasonRepo, logger log.Logger) *ReturnReasonUseCase {
	return &ReturnReasonUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (r ReturnReasonUseCase) CreateReturnReason(ctx context.Context, reason *ReturnReason) error {
	panic("implement me")
}

func (r ReturnReasonUseCase) GetReturnReason(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r ReturnReasonUseCase) UpdateReturnReason(ctx context.Context, reason *ReturnReason) error {
	panic("implement me")
}

func (r ReturnReasonUseCase) ListReturnReason(ctx context.Context, req *ReturnReasonListReq) (*ReturnReasonListResp, error) {
	return r.repo.ListReturnReason(ctx, req)
}

func (r ReturnReasonUseCase) DeleteReturnReason(ctx context.Context, id int64) error {
	panic("implement me")
}
