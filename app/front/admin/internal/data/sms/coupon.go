package sms

import (
	"context"
	smsV1 "github.com/feihua/kratos-mall/api/sms/v1"
	"github.com/feihua/kratos-mall/app/front/admin/internal/biz/sms"
	"github.com/feihua/kratos-mall/app/front/admin/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
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

func (c couponRepo) ListCoupon(ctx context.Context, req *sms.CouponListReq) (*sms.CouponListResp, error) {
	list, _ := c.data.SmsClient.CouponList(ctx, &smsV1.CouponListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*sms.Coupon, 0)
	copier.Copy(&orders, &list.List)

	return &sms.CouponListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (c couponRepo) DeleteCoupon(ctx context.Context, id int64) error {
	panic("implement me")
}
