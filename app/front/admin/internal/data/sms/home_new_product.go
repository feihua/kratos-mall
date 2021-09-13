package sms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	smsV1 "kratos-mall/api/sms/v1"
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

func (h homeNewProductRepo) ListHomeNewProduct(ctx context.Context, req *sms.HomeNewProductListReq) (*sms.HomeNewProductListResp, error) {
	list, _ := h.data.SmsClient.HomeNewProductList(ctx, &smsV1.HomeNewProductListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*sms.HomeNewProduct, 0)
	copier.Copy(&orders, &list.List)

	return &sms.HomeNewProductListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (h homeNewProductRepo) DeleteHomeNewProduct(ctx context.Context, id int64) error {
	panic("implement me")
}
