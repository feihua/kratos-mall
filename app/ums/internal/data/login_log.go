package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/ums/internal/biz"
)

type loginLogRepo struct {
	data *Data
	log  *log.Helper
}

func NewLoginLogRepo(data *Data, logger log.Logger) biz.LoginLogRepo {
	return &loginLogRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (l loginLogRepo) CreateLoginLog(ctx context.Context, loginLog *biz.LoginLog) error {
	panic("implement me")
}

func (l loginLogRepo) GetLoginLog(ctx context.Context, id int64) error {
	panic("implement me")
}

func (l loginLogRepo) UpdateLoginLog(ctx context.Context, loginLog *biz.LoginLog) error {
	panic("implement me")
}

func (l loginLogRepo) ListLoginLog(ctx context.Context, req *biz.LoginLogListReq) ([]*biz.LoginLog, error) {
	panic("implement me")
}

func (l loginLogRepo) DeleteLoginLog(ctx context.Context, id int64) error {
	panic("implement me")
}
