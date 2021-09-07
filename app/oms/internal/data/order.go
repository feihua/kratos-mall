package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/oms/internal/biz"
)

type orderRepo struct {
	data *Data
	log  *log.Helper
}

func NewOrderRepo(data *Data, logger log.Logger) biz.OrderRepo {
	return &orderRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (o orderRepo) CreateOrder(ctx context.Context, order *biz.Order) error {
	panic("implement me")
}

func (o orderRepo) GetOrder(ctx context.Context, id int64) error {
	panic("implement me")
}

func (o orderRepo) UpdateOrder(ctx context.Context, order *biz.Order) error {
	panic("implement me")
}

func (o orderRepo) ListOrder(ctx context.Context, req *biz.OrderListReq) ([]*biz.Order, error) {
	panic("implement me")
}

func (o orderRepo) DeleteOrder(ctx context.Context, id int64) error {
	panic("implement me")
}
