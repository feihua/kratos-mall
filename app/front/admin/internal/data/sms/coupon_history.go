package sms

import (
	"context"
	smsV1 "github.com/feihua/kratos-mall/api/sms/v1"
	"github.com/feihua/kratos-mall/app/front/admin/internal/biz/sms"
	"github.com/feihua/kratos-mall/app/front/admin/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
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

func (c couponHistoryRepo) ListCouponHistory(ctx context.Context, req *sms.CouponHistoryListReq) (*sms.CouponHistoryListResp, error) {
	list, _ := c.data.SmsClient.CouponHistoryList(ctx, &smsV1.CouponHistoryListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*sms.CouponHistory, 0)
	copier.Copy(&orders, &list.List)

	return &sms.CouponHistoryListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (c couponHistoryRepo) DeleteCouponHistory(ctx context.Context, id int64) error {
	panic("implement me")
}
