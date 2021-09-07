package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/pms/internal/biz"
)

type productRepo struct {
	data *Data
	log  *log.Helper
}

func NewProductRepo(data *Data, logger log.Logger) biz.ProductRepo {
	return &productRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (p productRepo) CreateProduct(ctx context.Context, product *biz.Product) error {
	panic("implement me")
}

func (p productRepo) GetProduct(ctx context.Context, id int64) error {
	panic("implement me")
}

func (p productRepo) UpdateProduct(ctx context.Context, product *biz.Product) error {
	panic("implement me")
}

func (p productRepo) ListProduct(ctx context.Context, req *biz.ProductListReq) ([]*biz.Product, error) {
	panic("implement me")
}

func (p productRepo) DeleteProduct(ctx context.Context, id int64) error {
	panic("implement me")
}
