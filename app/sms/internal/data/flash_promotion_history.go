package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/sms/internal/biz"
)

type flashPromotionHistoryRepo struct {
	data *Data
	log  *log.Helper
}

func NewFlashPromotionHistoryRepo(data *Data, logger log.Logger) biz.FlashPromotionHistoryRepo {
	return &flashPromotionHistoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (f flashPromotionHistoryRepo) CreateFlashPromotionHistory(ctx context.Context, history *biz.FlashPromotionHistory) error {
	panic("implement me")
}

func (f flashPromotionHistoryRepo) GetFlashPromotionHistory(ctx context.Context, id int64) error {
	panic("implement me")
}

func (f flashPromotionHistoryRepo) UpdateFlashPromotionHistory(ctx context.Context, history *biz.FlashPromotionHistory) error {
	panic("implement me")
}

func (f flashPromotionHistoryRepo) ListFlashPromotionHistory(ctx context.Context, req *biz.FlashPromotionHistoryListReq) ([]*biz.FlashPromotionHistory, error) {
	panic("implement me")
}

func (f flashPromotionHistoryRepo) DeleteFlashPromotionHistory(ctx context.Context, id int64) error {
	panic("implement me")
}
