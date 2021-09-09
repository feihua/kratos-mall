package pms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/front/admin/internal/biz/pms"
	"kratos-mall/app/front/admin/internal/data"
)

type skuStockRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewSkuStockRepo(data *data.Data, logger log.Logger) pms.SkuStockRepo {
	return &skuStockRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (s skuStockRepo) CreateSkuStock(ctx context.Context, stock *pms.SkuStock) error {
	panic("implement me")
}

func (s skuStockRepo) GetSkuStock(ctx context.Context, id int64) error {
	panic("implement me")
}

func (s skuStockRepo) UpdateSkuStock(ctx context.Context, stock *pms.SkuStock) error {
	panic("implement me")
}

func (s skuStockRepo) ListSkuStock(ctx context.Context, req *pms.SkuStockListReq) ([]*pms.SkuStock, error) {
	panic("implement me")
}

func (s skuStockRepo) DeleteSkuStock(ctx context.Context, id int64) error {
	panic("implement me")
}
