package data

import (
	"context"
	"github.com/feihua/kratos-mall/app/sms/internal/biz"
	"github.com/feihua/kratos-mall/app/sms/internal/data/model"
	"github.com/feihua/kratos-mall/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
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

func (c couponRepo) ListCoupon(ctx context.Context, req *biz.CouponListReq) (*biz.CouponListResp, error) {
	var all []model.SmsCoupon
	result := c.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	c.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.Coupon, 0)

	for _, coupon := range all {
		list = append(list, &biz.Coupon{
			Id:           coupon.Id,
			Type:         coupon.Type,
			Name:         coupon.Name,
			Platform:     coupon.Platform,
			Count:        coupon.Count,
			Amount:       coupon.Amount,
			PerLimit:     coupon.PerLimit,
			MinPoint:     coupon.MinPoint,
			StartTime:    coupon.StartTime.Format("2006-01-02 15:04:05"),
			EndTime:      coupon.EndTime.Format("2006-01-02 15:04:05"),
			UseType:      coupon.UseType,
			Note:         coupon.Note,
			PublishCount: coupon.PublishCount,
			UseCount:     coupon.UseCount,
			ReceiveCount: coupon.ReceiveCount,
			EnableTime:   coupon.EnableTime.Format("2006-01-02 15:04:05"),
			Code:         coupon.Code,
			MemberLevel:  coupon.MemberLevel,
		})
	}

	return &biz.CouponListResp{
		Total: count,
		List:  list,
	}, nil
}

func (c couponRepo) DeleteCoupon(ctx context.Context, id int64) error {
	panic("implement me")
}
