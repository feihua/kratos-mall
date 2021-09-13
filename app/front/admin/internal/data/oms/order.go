package oms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	omsV1 "kratos-mall/api/oms/v1"
	"kratos-mall/app/front/admin/internal/biz/oms"
	"kratos-mall/app/front/admin/internal/data"
)

type orderRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewOrderRepo(data *data.Data, logger log.Logger) oms.OrderRepo {
	return &orderRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (o orderRepo) CreateOrder(ctx context.Context, order *oms.Order) error {
	panic("implement me")
}

func (o orderRepo) GetOrder(ctx context.Context, id int64) error {
	panic("implement me")
}

func (o orderRepo) UpdateOrder(ctx context.Context, order *oms.Order) error {
	panic("implement me")
}

func (o orderRepo) ListOrder(ctx context.Context, req *oms.OrderListReq) (*oms.OrderListResp, error) {
	list, _ := o.data.OmsClient.OrderList(ctx, &omsV1.OrderListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*oms.Order, 0)
	copier.Copy(&orders, &list.List)

	return &oms.OrderListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (o orderRepo) DeleteOrder(ctx context.Context, id int64) error {
	panic("implement me")
}
