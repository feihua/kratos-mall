package oms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/front/admin/internal/biz/oms"
	"kratos-mall/app/front/admin/internal/data"
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

func (s settingRepo) ListSetting(ctx context.Context, req *oms.SettingListReq) ([]*oms.Setting, error) {
	panic("implement me")
}

func (s settingRepo) DeleteSetting(ctx context.Context, id int64) error {
	panic("implement me")
}
