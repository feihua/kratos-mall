package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/sms/internal/biz"
)

type flashPromotionRepo struct {
	data *Data
	log  *log.Helper
}

func NewFlashPromotionRepo(data *Data, logger log.Logger) biz.FlashPromotionRepo {
	return &flashPromotionRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (f flashPromotionRepo) CreateFlashPromotion(ctx context.Context, promotion *biz.FlashPromotion) error {
	panic("implement me")
}

func (f flashPromotionRepo) GetFlashPromotion(ctx context.Context, id int64) error {
	panic("implement me")
}

func (f flashPromotionRepo) UpdateFlashPromotion(ctx context.Context, promotion *biz.FlashPromotion) error {
	panic("implement me")
}

func (f flashPromotionRepo) ListFlashPromotion(ctx context.Context, req *biz.FlashPromotionListReq) ([]*biz.FlashPromotion, error) {
	panic("implement me")
}

func (f flashPromotionRepo) DeleteFlashPromotion(ctx context.Context, id int64) error {
	panic("implement me")
}
