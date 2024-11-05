package pms

import (
	"context"
	pmsV1 "github.com/feihua/kratos-mall/api/pms/v1"
	"github.com/feihua/kratos-mall/app/front/admin/internal/biz/pms"
	"github.com/feihua/kratos-mall/app/front/admin/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
)

type operateHistoryRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewOperateHistoryRepo(data *data.Data, logger log.Logger) pms.OperateHistoryRepo {
	return &operateHistoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (o operateHistoryRepo) CreateOperateHistory(ctx context.Context, history *pms.OperateHistory) error {
	panic("implement me")
}

func (o operateHistoryRepo) GetOperateHistory(ctx context.Context, id int64) error {
	panic("implement me")
}

func (o operateHistoryRepo) UpdateOperateHistory(ctx context.Context, history *pms.OperateHistory) error {
	panic("implement me")
}

func (o operateHistoryRepo) ListOperateHistory(ctx context.Context, req *pms.OperateHistoryListReq) (*pms.OperateHistoryListResp, error) {
	list, _ := o.data.PmsClient.ProductOperateLogList(ctx, &pmsV1.ProductOperateLogListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*pms.OperateHistory, 0)
	copier.Copy(&orders, &list.List)

	return &pms.OperateHistoryListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (o operateHistoryRepo) DeleteOperateHistory(ctx context.Context, id int64) error {
	panic("implement me")
}
