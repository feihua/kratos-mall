package data

import (
	"context"
	"github.com/feihua/kratos-mall/app/sms/internal/biz"
	"github.com/feihua/kratos-mall/app/sms/internal/data/model"
	"github.com/feihua/kratos-mall/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
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

func (c couponHistoryRepo) ListCouponHistory(ctx context.Context, req *biz.CouponHistoryListReq) (*biz.CouponHistoryListResp, error) {
	var all []model.SmsCouponHistory
	result := c.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	c.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.CouponHistory, 0)

	for _, item := range all {
		list = append(list, &biz.CouponHistory{
			Id:             item.Id,
			CouponId:       item.CouponId,
			MemberId:       item.MemberId,
			CouponCode:     item.CouponCode,
			MemberNickname: item.MemberNickname,
			GetType:        item.GetType,
			CreateTime:     item.CreateTime.Format("2006-01-02 15:04:05"),
			UseStatus:      item.UseStatus,
			UseTime:        item.UseTime.Format("2006-01-02 15:04:05"),
			OrderId:        item.OrderId,
			OrderSn:        item.OrderSn,
		})
	}

	return &biz.CouponHistoryListResp{
		Total: count,
		List:  list,
	}, nil
}

func (c couponHistoryRepo) DeleteCouponHistory(ctx context.Context, id int64) error {
	panic("implement me")
}
