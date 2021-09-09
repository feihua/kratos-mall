package oms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
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

func (c cartItemRepo) ListCartItem(ctx context.Context, req *oms.CartItemListReq) ([]*oms.CartItem, error) {
	panic("implement me")
}

func (c cartItemRepo) DeleteCartItem(ctx context.Context, id int64) error {
	panic("implement me")
}
