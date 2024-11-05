package sms

import (
	"context"
	smsV1 "github.com/feihua/kratos-mall/api/sms/v1"
	"github.com/feihua/kratos-mall/app/front/admin/internal/biz/sms"
	"github.com/feihua/kratos-mall/app/front/admin/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
)

type flashPromotionRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewFlashPromotionRepo(data *data.Data, logger log.Logger) sms.FlashPromotionRepo {
	return &flashPromotionRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (f flashPromotionRepo) CreateFlashPromotion(ctx context.Context, promotion *sms.FlashPromotion) error {
	panic("implement me")
}

func (f flashPromotionRepo) GetFlashPromotion(ctx context.Context, id int64) error {
	panic("implement me")
}

func (f flashPromotionRepo) UpdateFlashPromotion(ctx context.Context, promotion *sms.FlashPromotion) error {
	panic("implement me")
}

func (f flashPromotionRepo) ListFlashPromotion(ctx context.Context, req *sms.FlashPromotionListReq) (*sms.FlashPromotionListResp, error) {
	list, _ := f.data.SmsClient.FlashPromotionList(ctx, &smsV1.FlashPromotionListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*sms.FlashPromotion, 0)
	copier.Copy(&orders, &list.List)

	return &sms.FlashPromotionListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (f flashPromotionRepo) DeleteFlashPromotion(ctx context.Context, id int64) error {
	panic("implement me")
}
