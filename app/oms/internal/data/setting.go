package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/oms/internal/biz"
)

type settingRepo struct {
	data *Data
	log  *log.Helper
}

func NewSettingRepo(data *Data, logger log.Logger) biz.SettingRepo {
	return &settingRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (s settingRepo) CreateSetting(ctx context.Context, setting *biz.Setting) error {
	panic("implement me")
}

func (s settingRepo) GetSetting(ctx context.Context, id int64) error {
	panic("implement me")
}

func (s settingRepo) UpdateSetting(ctx context.Context, setting *biz.Setting) error {
	panic("implement me")
}

func (s settingRepo) ListSetting(ctx context.Context, req *biz.SettingListReq) ([]*biz.Setting, error) {
	panic("implement me")
}

func (s settingRepo) DeleteSetting(ctx context.Context, id int64) error {
	panic("implement me")
}
