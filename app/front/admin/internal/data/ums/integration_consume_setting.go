package ums

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	umsV1 "kratos-mall/api/ums/v1"
	"kratos-mall/app/front/admin/internal/biz/ums"
	"kratos-mall/app/front/admin/internal/data"
)

type consumeSettingRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewConsumeSettingRepo(data *data.Data, logger log.Logger) ums.ConsumeSettingRepo {
	return &consumeSettingRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c consumeSettingRepo) CreateConsumeSetting(ctx context.Context, setting *ums.ConsumeSetting) error {
	panic("implement me")
}

func (c consumeSettingRepo) GetConsumeSetting(ctx context.Context, id int64) error {
	panic("implement me")
}

func (c consumeSettingRepo) UpdateConsumeSetting(ctx context.Context, setting *ums.ConsumeSetting) error {
	panic("implement me")
}

func (c consumeSettingRepo) ListConsumeSetting(ctx context.Context, req *ums.ConsumeSettingListReq) (*ums.ConsumeSettingListResp, error) {
	list, _ := c.data.UmsClient.IntegrationConsumeSettingList(ctx, &umsV1.IntegrationConsumeSettingListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*ums.ConsumeSetting, 0)
	copier.Copy(&orders, &list.List)

	return &ums.ConsumeSettingListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (c consumeSettingRepo) DeleteConsumeSetting(ctx context.Context, id int64) error {
	panic("implement me")
}
