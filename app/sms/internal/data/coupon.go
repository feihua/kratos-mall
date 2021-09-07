package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/sms/internal/biz"
)

type couponRepo struct {
	data *Data
	log  *log.Helper
}

func NewCouponRepo(data *Data, logger log.Logger) biz.CouponRepo {
	return &couponRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c couponRepo) CreateCoupon(ctx context.Context, coupon *biz.Coupon) error {
	panic("implement me")
}

func (c couponRepo) GetCoupon(ctx context.Context, id int64) error {
	panic("implement me")
}

func (c couponRepo) UpdateCoupon(ctx context.Context, coupon *biz.Coupon) error {
	panic("implement me")
}

func (c couponRepo) ListCoupon(ctx context.Context, req *biz.CouponListReq) ([]*biz.Coupon, error) {
	panic("implement me")
}

func (c couponRepo) DeleteCoupon(ctx context.Context, id int64) error {
	panic("implement me")
}
