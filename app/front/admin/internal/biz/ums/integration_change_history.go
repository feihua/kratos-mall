package ums

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type IntegrationChangeHistoryListReq struct {
	Current  int64
	PageSize int64
}

type IntegrationChangeHistory struct {
	Id          int64
	MemberId    int64
	CreateTime  time.Time
	ChangeType  int    // 改变类型：0->增加；1->减少
	ChangeCount int    // 积分改变数量
	OperateMan  string // 操作人员
	OperateNote string // 操作备注
	SourceType  int    // 积分来源：0->购物；1->管理员修改
}

type IntegrationChangeHistoryRepo interface {
	CreateIntegrationChangeHistory(context.Context, *IntegrationChangeHistory) error
	GetIntegrationChangeHistory(ctx context.Context, id int64) error
	UpdateIntegrationChangeHistory(context.Context, *IntegrationChangeHistory) error
	ListIntegrationChangeHistory(ctx context.Context, req *IntegrationChangeHistoryListReq) ([]*IntegrationChangeHistory, error)
	DeleteIntegrationChangeHistory(ctx context.Context, id int64) error
}

type IntegrationChangeHistoryUseCase struct {
	repo IntegrationChangeHistoryRepo
	log  *log.Helper
}

func NewIntegrationChangeHistoryUseCase(repo IntegrationChangeHistoryRepo, logger log.Logger) *IntegrationChangeHistoryUseCase {
	return &IntegrationChangeHistoryUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (r *IntegrationChangeHistoryUseCase) CreateIntegrationChangeHistory(ctx context.Context, user *IntegrationChangeHistory) error {
	panic("implement me")
}

func (r *IntegrationChangeHistoryUseCase) GetIntegrationChangeHistory(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r *IntegrationChangeHistoryUseCase) UpdateIntegrationChangeHistory(ctx context.Context, user *IntegrationChangeHistory) error {
	panic("implement me")
}

func (r *IntegrationChangeHistoryUseCase) ListIntegrationChangeHistory(ctx context.Context, pageNum, pageSize int64) ([]*IntegrationChangeHistory, error) {
	panic("implement me")
}

func (r *IntegrationChangeHistoryUseCase) DeleteIntegrationChangeHistory(ctx context.Context, id int64) error {
	panic("implement me")
}
