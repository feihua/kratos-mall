package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/sms/internal/biz"
)

type flashPromotionSessionRepo struct {
	data *Data
	log  *log.Helper
}

func NewFlashPromotionSessionRepo(data *Data, logger log.Logger) biz.FlashPromotionSessionRepo {
	return &flashPromotionSessionRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (f flashPromotionSessionRepo) CreateFlashPromotionSession(ctx context.Context, session *biz.FlashPromotionSession) error {
	panic("implement me")
}

func (f flashPromotionSessionRepo) GetFlashPromotionSession(ctx context.Context, id int64) error {
	panic("implement me")
}

func (f flashPromotionSessionRepo) UpdateFlashPromotionSession(ctx context.Context, session *biz.FlashPromotionSession) error {
	panic("implement me")
}

func (f flashPromotionSessionRepo) ListFlashPromotionSession(ctx context.Context, req *biz.FlashPromotionSessionListReq) ([]*biz.FlashPromotionSession, error) {
	panic("implement me")
}

func (f flashPromotionSessionRepo) DeleteFlashPromotionSession(ctx context.Context, id int64) error {
	panic("implement me")
}
