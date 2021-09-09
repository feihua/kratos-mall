package sms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/front/admin/internal/biz/sms"
	"kratos-mall/app/front/admin/internal/data"
)

type couponHistoryRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewCouponHistoryRepo(data *data.Data, logger log.Logger) sms.CouponHistoryRepo {
	return &couponHistoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c couponHistoryRepo) CreateCouponHistory(ctx context.Context, history *sms.CouponHistory) error {
	panic("implement me")
}

func (c couponHistoryRepo) GetCouponHistory(ctx context.Context, id int64) error {
	panic("implement me")
}

func (c couponHistoryRepo) UpdateCouponHistory(ctx context.Context, history *sms.CouponHistory) error {
	panic("implement me")
}

func (c couponHistoryRepo) ListCouponHistory(ctx context.Context, req *sms.CouponHistoryListReq) ([]*sms.CouponHistory, error) {
	panic("implement me")
}

func (c couponHistoryRepo) DeleteCouponHistory(ctx context.Context, id int64) error {
	panic("implement me")
}
