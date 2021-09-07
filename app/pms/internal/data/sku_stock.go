package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/pms/internal/biz"
)

type skuStockRepo struct {
	data *Data
	log  *log.Helper
}

func NewSkuStockRepo(data *Data, logger log.Logger) biz.SkuStockRepo {
	return &skuStockRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (s skuStockRepo) CreateSkuStock(ctx context.Context, stock *biz.SkuStock) error {
	panic("implement me")
}

func (s skuStockRepo) GetSkuStock(ctx context.Context, id int64) error {
	panic("implement me")
}

func (s skuStockRepo) UpdateSkuStock(ctx context.Context, stock *biz.SkuStock) error {
	panic("implement me")
}

func (s skuStockRepo) ListSkuStock(ctx context.Context, req *biz.SkuStockListReq) ([]*biz.SkuStock, error) {
	panic("implement me")
}

func (s skuStockRepo) DeleteSkuStock(ctx context.Context, id int64) error {
	panic("implement me")
}
