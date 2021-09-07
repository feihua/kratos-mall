package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/sms/internal/biz"
)

type homeNewProductRepo struct {
	data *Data
	log  *log.Helper
}

func NewHomeNewProductRepo(data *Data, logger log.Logger) biz.HomeNewProductRepo {
	return &homeNewProductRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (h homeNewProductRepo) CreateHomeNewProduct(ctx context.Context, product *biz.HomeNewProduct) error {
	panic("implement me")
}

func (h homeNewProductRepo) GetHomeNewProduct(ctx context.Context, id int64) error {
	panic("implement me")
}

func (h homeNewProductRepo) UpdateHomeNewProduct(ctx context.Context, product *biz.HomeNewProduct) error {
	panic("implement me")
}

func (h homeNewProductRepo) ListHomeNewProduct(ctx context.Context, req *biz.HomeNewProductListReq) ([]*biz.HomeNewProduct, error) {
	panic("implement me")
}

func (h homeNewProductRepo) DeleteHomeNewProduct(ctx context.Context, id int64) error {
	panic("implement me")
}
