package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/sms/internal/biz"
	"kratos-mall/app/sms/internal/data/model"
	"kratos-mall/pkg/util/pagination"
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

func (h homeNewProductRepo) ListHomeNewProduct(ctx context.Context, req *biz.HomeNewProductListReq) (*biz.HomeNewProductListResp, error) {
	var all []model.SmsHomeNewProduct
	result := h.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	h.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.HomeNewProduct, 0)

	for _, item := range all {
		list = append(list, &biz.HomeNewProduct{
			Id:              item.Id,
			ProductId:       item.ProductId,
			ProductName:     item.ProductName,
			RecommendStatus: item.RecommendStatus,
			Sort:            item.Sort,
		})
	}

	return &biz.HomeNewProductListResp{
		Total: count,
		List:  list,
	}, nil
}

func (h homeNewProductRepo) DeleteHomeNewProduct(ctx context.Context, id int64) error {
	panic("implement me")
}
