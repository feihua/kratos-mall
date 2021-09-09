package sms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/front/admin/internal/biz/sms"
	"kratos-mall/app/front/admin/internal/data"
)

type homeRecommendProductRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewHomeRecommendProductRepo(data *data.Data, logger log.Logger) sms.HomeRecommendProductRepo {
	return &homeRecommendProductRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (h homeRecommendProductRepo) CreateHomeRecommendProduct(ctx context.Context, product *sms.HomeRecommendProduct) error {
	panic("implement me")
}

func (h homeRecommendProductRepo) GetHomeRecommendProduct(ctx context.Context, id int64) error {
	panic("implement me")
}

func (h homeRecommendProductRepo) UpdateHomeRecommendProduct(ctx context.Context, product *sms.HomeRecommendProduct) error {
	panic("implement me")
}

func (h homeRecommendProductRepo) ListHomeRecommendProduct(ctx context.Context, req *sms.HomeRecommendProductListReq) ([]*sms.HomeRecommendProduct, error) {
	panic("implement me")
}

func (h homeRecommendProductRepo) DeleteHomeRecommendProduct(ctx context.Context, id int64) error {
	panic("implement me")
}
