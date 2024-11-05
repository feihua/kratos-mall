package sms

import (
	"context"
	smsV1 "github.com/feihua/kratos-mall/api/sms/v1"
	"github.com/feihua/kratos-mall/app/front/admin/internal/biz/sms"
	"github.com/feihua/kratos-mall/app/front/admin/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
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

func (h homeRecommendProductRepo) ListHomeRecommendProduct(ctx context.Context, req *sms.HomeRecommendProductListReq) (*sms.HomeRecommendProductListResp, error) {
	list, _ := h.data.SmsClient.HomeRecommendProductList(ctx, &smsV1.HomeRecommendProductListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*sms.HomeRecommendProduct, 0)
	copier.Copy(&orders, &list.List)

	return &sms.HomeRecommendProductListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (h homeRecommendProductRepo) DeleteHomeRecommendProduct(ctx context.Context, id int64) error {
	panic("implement me")
}
