package ums

import (
	"context"
	umsV1 "github.com/feihua/kratos-mall/api/ums/v1"
	"github.com/feihua/kratos-mall/app/front/admin/internal/biz/ums"
	"github.com/feihua/kratos-mall/app/front/admin/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
)

type changeHistoryRepo struct {
	data *data.Data
	log  *log.Helper
}

// NewChangeHistoryRepo .
func NewChangeHistoryRepo(data *data.Data, logger log.Logger) ums.ChangeHistoryRepo {
	return &changeHistoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c changeHistoryRepo) CreateChangeHistory(ctx context.Context, history *ums.ChangeHistory) error {
	panic("implement me")
}

func (c changeHistoryRepo) GetChangeHistory(ctx context.Context, id int64) error {
	panic("implement me")
}

func (c changeHistoryRepo) UpdateChangeHistory(ctx context.Context, history *ums.ChangeHistory) error {
	panic("implement me")
}

func (c changeHistoryRepo) ListChangeHistory(ctx context.Context, req *ums.ChangeHistoryListReq) (*ums.ChangeHistoryListResp, error) {
	list, _ := c.data.UmsClient.GrowthChangeHistoryList(ctx, &umsV1.GrowthChangeHistoryListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*ums.ChangeHistory, 0)
	copier.Copy(&orders, &list.List)

	return &ums.ChangeHistoryListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (c changeHistoryRepo) DeleteChangeHistory(ctx context.Context, id int64) error {
	panic("implement me")
}
