package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/sms/internal/biz"
)

type couponHistoryRepo struct {
	data *Data
	log  *log.Helper
}

func NewCouponHistoryRepo(data *Data, logger log.Logger) biz.CouponHistoryRepo {
	return &couponHistoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c couponHistoryRepo) CreateCouponHistory(ctx context.Context, history *biz.CouponHistory) error {
	panic("implement me")
}

func (c couponHistoryRepo) GetCouponHistory(ctx context.Context, id int64) error {
	panic("implement me")
}

func (c couponHistoryRepo) UpdateCouponHistory(ctx context.Context, history *biz.CouponHistory) error {
	panic("implement me")
}

func (c couponHistoryRepo) ListCouponHistory(ctx context.Context, req *biz.CouponHistoryListReq) ([]*biz.CouponHistory, error) {
	panic("implement me")
}

func (c couponHistoryRepo) DeleteCouponHistory(ctx context.Context, id int64) error {
	panic("implement me")
}
