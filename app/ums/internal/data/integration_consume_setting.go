package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/ums/internal/biz"
)

type consumeSettingRepo struct {
	data *Data
	log  *log.Helper
}

func NewConsumeSettingRepo(data *Data, logger log.Logger) biz.ConsumeSettingRepo {
	return &consumeSettingRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c consumeSettingRepo) CreateConsumeSetting(ctx context.Context, setting *biz.ConsumeSetting) error {
	panic("implement me")
}

func (c consumeSettingRepo) GetConsumeSetting(ctx context.Context, id int64) error {
	panic("implement me")
}

func (c consumeSettingRepo) UpdateConsumeSetting(ctx context.Context, setting *biz.ConsumeSetting) error {
	panic("implement me")
}

func (c consumeSettingRepo) ListConsumeSetting(ctx context.Context, req *biz.ConsumeSettingListReq) ([]*biz.ConsumeSetting, error) {
	panic("implement me")
}

func (c consumeSettingRepo) DeleteConsumeSetting(ctx context.Context, id int64) error {
	panic("implement me")
}
