package ums

import (
	"context"
	umsV1 "github.com/feihua/kratos-mall/api/ums/v1"
	"github.com/feihua/kratos-mall/app/front/admin/internal/biz/ums"
	"github.com/feihua/kratos-mall/app/front/admin/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
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

func (l loginLogRepo) ListLoginLog(ctx context.Context, req *ums.LoginLogListReq) (*ums.LoginLogListResp, error) {
	list, _ := l.data.UmsClient.MemberLoginLogList(ctx, &umsV1.MemberLoginLogListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*ums.LoginLog, 0)
	copier.Copy(&orders, &list.List)

	return &ums.LoginLogListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (l loginLogRepo) DeleteLoginLog(ctx context.Context, id int64) error {
	panic("implement me")
}
