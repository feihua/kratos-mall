package ums

import (
	"context"
	umsV1 "github.com/feihua/kratos-mall/api/ums/v1"
	"github.com/feihua/kratos-mall/app/front/admin/internal/biz/ums"
	"github.com/feihua/kratos-mall/app/front/admin/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
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

func (i integrationChangeHistoryRepo) ListIntegrationChangeHistory(ctx context.Context, req *ums.IntegrationChangeHistoryListReq) (*ums.IntegrationChangeHistoryListResp, error) {
	list, _ := i.data.UmsClient.IntegrationChangeHistoryList(ctx, &umsV1.IntegrationChangeHistoryListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*ums.IntegrationChangeHistory, 0)
	copier.Copy(&orders, &list.List)

	return &ums.IntegrationChangeHistoryListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (i integrationChangeHistoryRepo) DeleteIntegrationChangeHistory(ctx context.Context, id int64) error {
	panic("implement me")
}
