package data

import (
	"context"
	"github.com/feihua/kratos-mall/app/pms/internal/biz"
	"github.com/feihua/kratos-mall/app/pms/internal/data/model"
	"github.com/feihua/kratos-mall/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
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

func (s skuStockRepo) ListSkuStock(ctx context.Context, req *biz.SkuStockListReq) (*biz.SkuStockListResp, error) {
	var all []model.PmsSkuStock
	result := s.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	s.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.SkuStock, 0)

	for _, item := range all {
		list = append(list, &biz.SkuStock{
			Id:             item.Id,
			ProductId:      item.ProductId,
			SkuCode:        item.SkuCode,
			Price:          item.Price,
			Stock:          item.Stock,
			LowStock:       item.LowStock,
			Pic:            item.Pic,
			Sale:           item.Sale,
			PromotionPrice: item.PromotionPrice,
			LockStock:      item.LockStock,
			SpData:         item.SpData,
		})
	}

	return &biz.SkuStockListResp{
		Total: count,
		List:  list,
	}, nil
}

func (s skuStockRepo) DeleteSkuStock(ctx context.Context, id int64) error {
	panic("implement me")
}
