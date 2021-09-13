package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type IntegrationChangeHistoryListReq struct {
	Current  int64
	PageSize int64
}

type IntegrationChangeHistory struct {
	Id          int64
	MemberId    int64
	CreateTime  string
	ChangeType  int    // 改变类型：0->增加；1->减少
	ChangeCount int    // 积分改变数量
	OperateMan  string // 操作人员
	OperateNote string // 操作备注
	SourceType  int    // 积分来源：0->购物；1->管理员修改
}
type IntegrationChangeHistoryListResp struct {
	Total int64
	List  []*IntegrationChangeHistory
}

type IntegrationChangeHistoryRepo interface {
	CreateIntegrationChangeHistory(context.Context, *IntegrationChangeHistory) error
	GetIntegrationChangeHistory(ctx context.Context, id int64) error
	UpdateIntegrationChangeHistory(context.Context, *IntegrationChangeHistory) error
	ListIntegrationChangeHistory(ctx context.Context, req *IntegrationChangeHistoryListReq) (*IntegrationChangeHistoryListResp, error)
	DeleteIntegrationChangeHistory(ctx context.Context, id int64) error
}

type IntegrationChangeHistoryUseCase struct {
	repo IntegrationChangeHistoryRepo
	log  *log.Helper
}

func NewIntegrationChangeHistoryUseCase(repo IntegrationChangeHistoryRepo, logger log.Logger) *IntegrationChangeHistoryUseCase {
	return &IntegrationChangeHistoryUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (i IntegrationChangeHistoryUseCase) CreateIntegrationChangeHistory(ctx context.Context, history *IntegrationChangeHistory) error {
	panic("implement me")
}

func (i IntegrationChangeHistoryUseCase) GetIntegrationChangeHistory(ctx context.Context, id int64) error {
	panic("implement me")
}

func (i IntegrationChangeHistoryUseCase) UpdateIntegrationChangeHistory(ctx context.Context, history *IntegrationChangeHistory) error {
	panic("implement me")
}

func (i IntegrationChangeHistoryUseCase) ListIntegrationChangeHistory(ctx context.Context, req *IntegrationChangeHistoryListReq) (*IntegrationChangeHistoryListResp, error) {
	return i.repo.ListIntegrationChangeHistory(ctx, req)
}

func (i IntegrationChangeHistoryUseCase) DeleteIntegrationChangeHistory(ctx context.Context, id int64) error {
	panic("implement me")
}
