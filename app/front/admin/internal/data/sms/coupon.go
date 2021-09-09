package sms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/front/admin/internal/biz/sms"
	"kratos-mall/app/front/admin/internal/data"
)

type couponRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewCouponRepo(data *data.Data, logger log.Logger) sms.CouponRepo {
	return &couponRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c couponRepo) CreateCoupon(ctx context.Context, coupon *sms.Coupon) error {
	panic("implement me")
}

func (c couponRepo) GetCoupon(ctx context.Context, id int64) error {
	panic("implement me")
}

func (c couponRepo) UpdateCoupon(ctx context.Context, coupon *sms.Coupon) error {
	panic("implement me")
}

func (c couponRepo) ListCoupon(ctx context.Context, req *sms.CouponListReq) ([]*sms.Coupon, error) {
	panic("implement me")
}

func (c couponRepo) DeleteCoupon(ctx context.Context, id int64) error {
	panic("implement me")
}
