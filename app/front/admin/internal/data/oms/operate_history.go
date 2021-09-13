package oms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	omsV1 "kratos-mall/api/oms/v1"
	"kratos-mall/app/front/admin/internal/biz/oms"
	"kratos-mall/app/front/admin/internal/data"
)

type operateHistoryRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewOperateHistoryRepo(logger log.Logger) oms.OperateHistoryRepo {
	return &operateHistoryRepo{
		log: log.NewHelper(logger),
	}
}

func (o operateHistoryRepo) CreateOperateHistory(ctx context.Context, history *oms.OperateHistory) error {
	panic("implement me")
}

func (o operateHistoryRepo) GetOperateHistory(ctx context.Context, id int64) error {
	panic("implement me")
}

func (o operateHistoryRepo) UpdateOperateHistory(ctx context.Context, history *oms.OperateHistory) error {
	panic("implement me")
}

func (o operateHistoryRepo) ListOperateHistory(ctx context.Context, req *oms.OperateHistoryListReq) (*oms.OperateHistoryListResp, error) {
	list, _ := o.data.OmsClient.OrderOperateHistoryList(ctx, &omsV1.OrderOperateHistoryListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*oms.OperateHistory, 0)
	copier.Copy(&orders, &list.List)
	return &oms.OperateHistoryListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (o operateHistoryRepo) DeleteOperateHistory(ctx context.Context, id int64) error {
	panic("implement me")
}
