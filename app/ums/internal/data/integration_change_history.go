package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/ums/internal/biz"
)

type integrationChangeHistoryRepo struct {
	data *Data
	log  *log.Helper
}

func NewIntegrationChangeHistoryRepo(data *Data, logger log.Logger) biz.IntegrationChangeHistoryRepo {
	return &integrationChangeHistoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (i integrationChangeHistoryRepo) CreateIntegrationChangeHistory(ctx context.Context, history *biz.IntegrationChangeHistory) error {
	panic("implement me")
}

func (i integrationChangeHistoryRepo) GetIntegrationChangeHistory(ctx context.Context, id int64) error {
	panic("implement me")
}

func (i integrationChangeHistoryRepo) UpdateIntegrationChangeHistory(ctx context.Context, history *biz.IntegrationChangeHistory) error {
	panic("implement me")
}

func (i integrationChangeHistoryRepo) ListIntegrationChangeHistory(ctx context.Context, req *biz.IntegrationChangeHistoryListReq) ([]*biz.IntegrationChangeHistory, error) {
	panic("implement me")
}

func (i integrationChangeHistoryRepo) DeleteIntegrationChangeHistory(ctx context.Context, id int64) error {
	panic("implement me")
}
