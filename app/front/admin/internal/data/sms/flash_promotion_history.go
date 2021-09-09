package sms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/front/admin/internal/biz/sms"
	"kratos-mall/app/front/admin/internal/data"
)

type flashPromotionHistoryRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewFlashPromotionHistoryRepo(data *data.Data, logger log.Logger) sms.FlashPromotionHistoryRepo {
	return &flashPromotionHistoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (f flashPromotionHistoryRepo) CreateFlashPromotionHistory(ctx context.Context, history *sms.FlashPromotionHistory) error {
	panic("implement me")
}

func (f flashPromotionHistoryRepo) GetFlashPromotionHistory(ctx context.Context, id int64) error {
	panic("implement me")
}

func (f flashPromotionHistoryRepo) UpdateFlashPromotionHistory(ctx context.Context, history *sms.FlashPromotionHistory) error {
	panic("implement me")
}

func (f flashPromotionHistoryRepo) ListFlashPromotionHistory(ctx context.Context, req *sms.FlashPromotionHistoryListReq) ([]*sms.FlashPromotionHistory, error) {
	panic("implement me")
}

func (f flashPromotionHistoryRepo) DeleteFlashPromotionHistory(ctx context.Context, id int64) error {
	panic("implement me")
}
