package ums

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/front/admin/internal/biz/ums"
	"kratos-mall/app/front/admin/internal/data"
)

type loginLogRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewLoginLogRepo(data *data.Data, logger log.Logger) ums.LoginLogRepo {
	return &loginLogRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (l loginLogRepo) CreateLoginLog(ctx context.Context, loginLog *ums.LoginLog) error {
	panic("implement me")
}

func (l loginLogRepo) GetLoginLog(ctx context.Context, id int64) error {
	panic("implement me")
}

func (l loginLogRepo) UpdateLoginLog(ctx context.Context, loginLog *ums.LoginLog) error {
	panic("implement me")
}

func (l loginLogRepo) ListLoginLog(ctx context.Context, req *ums.LoginLogListReq) ([]*ums.LoginLog, error) {
	panic("implement me")
}

func (l loginLogRepo) DeleteLoginLog(ctx context.Context, id int64) error {
	panic("implement me")
}
