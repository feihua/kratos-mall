package oms

import (
	"context"
	omsV1 "github.com/feihua/kratos-mall/api/oms/v1"
	"github.com/feihua/kratos-mall/app/front/admin/internal/biz/oms"
	"github.com/feihua/kratos-mall/app/front/admin/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
)

type settingRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewSettingRepo(data *data.Data, logger log.Logger) oms.SettingRepo {
	return &settingRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (s settingRepo) CreateSetting(ctx context.Context, setting *oms.Setting) error {
	panic("implement me")
}

func (s settingRepo) GetSetting(ctx context.Context, id int64) error {
	panic("implement me")
}

func (s settingRepo) UpdateSetting(ctx context.Context, setting *oms.Setting) error {
	panic("implement me")
}

func (s settingRepo) ListSetting(ctx context.Context, req *oms.SettingListReq) (*oms.SettingListResp, error) {
	list, _ := s.data.OmsClient.OrderSettingList(ctx, &omsV1.OrderSettingListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*oms.Setting, 0)
	copier.Copy(&orders, &list.List)
	return &oms.SettingListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (s settingRepo) DeleteSetting(ctx context.Context, id int64) error {
	panic("implement me")
}
