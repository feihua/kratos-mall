package ums

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type ChangeHistoryListReq struct {
	Current  int64
	PageSize int64
}

type ChangeHistory struct {
	Id          int64
	MemberId    int64
	CreateTime  string
	ChangeType  int    // 改变类型：0->增加；1->减少
	ChangeCount int    // 积分改变数量
	OperateMan  string // 操作人员
	OperateNote string // 操作备注
	SourceType  int    // 积分来源：0->购物；1->管理员修改
}
type ChangeHistoryListResp struct {
	Total int64
	List  []*ChangeHistory
}

type ChangeHistoryRepo interface {
	CreateChangeHistory(context.Context, *ChangeHistory) error
	GetChangeHistory(ctx context.Context, id int64) error
	UpdateChangeHistory(context.Context, *ChangeHistory) error
	ListChangeHistory(ctx context.Context, req *ChangeHistoryListReq) (*ChangeHistoryListResp, error)
	DeleteChangeHistory(ctx context.Context, id int64) error
}

type ChangeHistoryUseCase struct {
	repo ChangeHistoryRepo
	log  *log.Helper
}

func NewChangeHistoryUseCase(repo ChangeHistoryRepo, logger log.Logger) *ChangeHistoryUseCase {
	return &ChangeHistoryUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (c ChangeHistoryUseCase) CreateChangeHistory(ctx context.Context, history *ChangeHistory) error {
	panic("implement me")
}

func (c ChangeHistoryUseCase) GetChangeHistory(ctx context.Context, id int64) error {
	panic("implement me")
}

func (c ChangeHistoryUseCase) UpdateChangeHistory(ctx context.Context, history *ChangeHistory) error {
	panic("implement me")
}

func (c ChangeHistoryUseCase) ListChangeHistory(ctx context.Context, req *ChangeHistoryListReq) (*ChangeHistoryListResp, error) {
	return c.repo.ListChangeHistory(ctx, req)
}

func (c ChangeHistoryUseCase) DeleteChangeHistory(ctx context.Context, id int64) error {
	panic("implement me")
}
