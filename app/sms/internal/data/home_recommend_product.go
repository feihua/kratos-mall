package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/sms/internal/biz"
	"kratos-mall/app/sms/internal/data/model"
	"kratos-mall/pkg/util/pagination"
)

type homeRecommendProductRepo struct {
	data *Data
	log  *log.Helper
}

func NewHomeRecommendProductRepo(data *Data, logger log.Logger) biz.HomeRecommendProductRepo {
	return &homeRecommendProductRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (h homeRecommendProductRepo) CreateHomeRecommendProduct(ctx context.Context, product *biz.HomeRecommendProduct) error {
	panic("implement me")
}

func (h homeRecommendProductRepo) GetHomeRecommendProduct(ctx context.Context, id int64) error {
	panic("implement me")
}

func (h homeRecommendProductRepo) UpdateHomeRecommendProduct(ctx context.Context, product *biz.HomeRecommendProduct) error {
	panic("implement me")
}

func (h homeRecommendProductRepo) ListHomeRecommendProduct(ctx context.Context, req *biz.HomeRecommendProductListReq) (*biz.HomeRecommendProductListResp, error) {
	var all []model.SmsHomeRecommendProduct
	result := h.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	h.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.HomeRecommendProduct, 0)

	for _, item := range all {
		list = append(list, &biz.HomeRecommendProduct{
			Id:              item.Id,
			ProductId:       item.ProductId,
			ProductName:     item.ProductName,
			RecommendStatus: item.RecommendStatus,
			Sort:            item.Sort,
		})
	}

	return &biz.HomeRecommendProductListResp{
		Total: count,
		List:  list,
	}, nil
}

func (h homeRecommendProductRepo) DeleteHomeRecommendProduct(ctx context.Context, id int64) error {
	panic("implement me")
}
