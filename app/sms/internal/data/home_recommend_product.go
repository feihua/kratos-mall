package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/sms/internal/biz"
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

func (h homeRecommendProductRepo) ListHomeRecommendProduct(ctx context.Context, req *biz.HomeRecommendProductListReq) ([]*biz.HomeRecommendProduct, error) {
	panic("implement me")
}

func (h homeRecommendProductRepo) DeleteHomeRecommendProduct(ctx context.Context, id int64) error {
	panic("implement me")
}
