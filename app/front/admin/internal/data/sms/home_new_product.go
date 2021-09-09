package sms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/front/admin/internal/biz/sms"
	"kratos-mall/app/front/admin/internal/data"
)

type homeNewProductRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewHomeNewProductRepo(data *data.Data, logger log.Logger) sms.HomeNewProductRepo {
	return &homeNewProductRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (h homeNewProductRepo) CreateHomeNewProduct(ctx context.Context, product *sms.HomeNewProduct) error {
	panic("implement me")
}

func (h homeNewProductRepo) GetHomeNewProduct(ctx context.Context, id int64) error {
	panic("implement me")
}

func (h homeNewProductRepo) UpdateHomeNewProduct(ctx context.Context, product *sms.HomeNewProduct) error {
	panic("implement me")
}

func (h homeNewProductRepo) ListHomeNewProduct(ctx context.Context, req *sms.HomeNewProductListReq) ([]*sms.HomeNewProduct, error) {
	panic("implement me")
}

func (h homeNewProductRepo) DeleteHomeNewProduct(ctx context.Context, id int64) error {
	panic("implement me")
}
