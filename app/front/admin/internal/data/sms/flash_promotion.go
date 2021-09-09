package sms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/front/admin/internal/biz/sms"
	"kratos-mall/app/front/admin/internal/data"
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

func (f flashPromotionRepo) ListFlashPromotion(ctx context.Context, req *sms.FlashPromotionListReq) ([]*sms.FlashPromotion, error) {
	panic("implement me")
}

func (f flashPromotionRepo) DeleteFlashPromotion(ctx context.Context, id int64) error {
	panic("implement me")
}
