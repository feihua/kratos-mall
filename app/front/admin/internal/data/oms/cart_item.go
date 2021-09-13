package oms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	omsV1 "kratos-mall/api/oms/v1"
	"kratos-mall/app/front/admin/internal/biz/oms"
	"kratos-mall/app/front/admin/internal/data"
)

type cartItemRepo struct {
	data *data.Data
	log  *log.Helper
}

// NewCartItemRepo .
func NewCartItemRepo(data *data.Data, logger log.Logger) oms.CartItemRepo {
	return &cartItemRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c cartItemRepo) CreateCartItem(ctx context.Context, item *oms.CartItem) error {
	panic("implement me")
}

func (c cartItemRepo) GetCartItem(ctx context.Context, id int64) error {
	panic("implement me")
}

func (c cartItemRepo) UpdateCartItem(ctx context.Context, item *oms.CartItem) error {
	panic("implement me")
}

func (c cartItemRepo) ListCartItem(ctx context.Context, req *oms.CartItemListReq) (*oms.CartItemListResp, error) {
	list, _ := c.data.OmsClient.CartItemList(ctx, &omsV1.CartItemListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*oms.CartItem, 0)
	copier.Copy(&orders, &list.List)
	return &oms.CartItemListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (c cartItemRepo) DeleteCartItem(ctx context.Context, id int64) error {
	panic("implement me")
}
