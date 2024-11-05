package pms

import (
	"context"
	pmsV1 "github.com/feihua/kratos-mall/api/pms/v1"
	"github.com/feihua/kratos-mall/app/front/admin/internal/biz/pms"
	"github.com/feihua/kratos-mall/app/front/admin/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
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

func (s skuStockRepo) ListSkuStock(ctx context.Context, req *pms.SkuStockListReq) (*pms.SkuStockListResp, error) {
	list, _ := s.data.PmsClient.SkuStockList(ctx, &pmsV1.SkuStockListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*pms.SkuStock, 0)
	copier.Copy(&orders, &list.List)

	return &pms.SkuStockListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (s skuStockRepo) DeleteSkuStock(ctx context.Context, id int64) error {
	panic("implement me")
}
