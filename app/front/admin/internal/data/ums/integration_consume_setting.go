package ums

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
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

func (c consumeSettingRepo) ListConsumeSetting(ctx context.Context, req *ums.ConsumeSettingListReq) ([]*ums.ConsumeSetting, error) {
	panic("implement me")
}

func (c consumeSettingRepo) DeleteConsumeSetting(ctx context.Context, id int64) error {
	panic("implement me")
}
