package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/pay/internal/biz"
)

type merchantRepo struct {
	data *Data
	log  *log.Helper
}

func NewMerchantRepo(data *Data, logger log.Logger) biz.MerchantRepo {
	return &merchantRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (m merchantRepo) CreateMerchant(ctx context.Context, merchant *biz.Merchant) error {
	panic("implement me")
}

func (m merchantRepo) GetMerchant(ctx context.Context, id int64) error {
	panic("implement me")
}

func (m merchantRepo) UpdateMerchant(ctx context.Context, merchant *biz.Merchant) error {
	panic("implement me")
}

func (m merchantRepo) ListMerchant(ctx context.Context, req *biz.MerchantListReq) ([]*biz.Merchant, error) {
	panic("implement me")
}

func (m merchantRepo) DeleteMerchant(ctx context.Context, id int64) error {
	panic("implement me")
}
