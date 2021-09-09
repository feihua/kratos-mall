package ums

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/front/admin/internal/biz/ums"
	"kratos-mall/app/front/admin/internal/data"
)

type integrationChangeHistoryRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewIntegrationChangeHistoryRepo(data *data.Data, logger log.Logger) ums.IntegrationChangeHistoryRepo {
	return &integrationChangeHistoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (i integrationChangeHistoryRepo) CreateIntegrationChangeHistory(ctx context.Context, history *ums.IntegrationChangeHistory) error {
	panic("implement me")
}

func (i integrationChangeHistoryRepo) GetIntegrationChangeHistory(ctx context.Context, id int64) error {
	panic("implement me")
}

func (i integrationChangeHistoryRepo) UpdateIntegrationChangeHistory(ctx context.Context, history *ums.IntegrationChangeHistory) error {
	panic("implement me")
}

func (i integrationChangeHistoryRepo) ListIntegrationChangeHistory(ctx context.Context, req *ums.IntegrationChangeHistoryListReq) ([]*ums.IntegrationChangeHistory, error) {
	panic("implement me")
}

func (i integrationChangeHistoryRepo) DeleteIntegrationChangeHistory(ctx context.Context, id int64) error {
	panic("implement me")
}
