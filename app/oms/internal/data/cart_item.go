package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/oms/internal/biz"
)

type cartItemRepo struct {
	data *Data
	log  *log.Helper
}

// NewCartItemRepo .
func NewCartItemRepo(data *Data, logger log.Logger) biz.CartItemRepo {
	return &cartItemRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c cartItemRepo) CreateCartItem(ctx context.Context, item *biz.CartItem) error {
	panic("implement me")
}

func (c cartItemRepo) GetCartItem(ctx context.Context, id int64) error {
	panic("implement me")
}

func (c cartItemRepo) UpdateCartItem(ctx context.Context, item *biz.CartItem) error {
	panic("implement me")
}

func (c cartItemRepo) ListCartItem(ctx context.Context, req *biz.CartItemListReq) ([]*biz.CartItem, error) {
	panic("implement me")
}

func (c cartItemRepo) DeleteCartItem(ctx context.Context, id int64) error {
	panic("implement me")
}
