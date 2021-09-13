package pay

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/front/admin/internal/biz/pay"
	"kratos-mall/app/front/admin/internal/data"
)

type merchantRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewMerchantRepo(data *data.Data, logger log.Logger) pay.MerchantRepo {
	return &merchantRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (m merchantRepo) CreateMerchant(ctx context.Context, merchant *pay.Merchant) error {
	panic("implement me")
}

func (m merchantRepo) GetMerchant(ctx context.Context, id int64) error {
	panic("implement me")
}

func (m merchantRepo) UpdateMerchant(ctx context.Context, merchant *pay.Merchant) error {
	panic("implement me")
}

func (m merchantRepo) ListMerchant(ctx context.Context, req *pay.MerchantListReq) (*pay.MerchantListResp, error) {
	panic("implement me")
}

func (m merchantRepo) DeleteMerchant(ctx context.Context, id int64) error {
	panic("implement me")
}
