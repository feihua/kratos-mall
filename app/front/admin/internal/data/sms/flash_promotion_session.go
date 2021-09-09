package sms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/front/admin/internal/biz/sms"
	"kratos-mall/app/front/admin/internal/data"
)

type flashPromotionSessionRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewFlashPromotionSessionRepo(data *data.Data, logger log.Logger) sms.FlashPromotionSessionRepo {
	return &flashPromotionSessionRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (f flashPromotionSessionRepo) CreateFlashPromotionSession(ctx context.Context, session *sms.FlashPromotionSession) error {
	panic("implement me")
}

func (f flashPromotionSessionRepo) GetFlashPromotionSession(ctx context.Context, id int64) error {
	panic("implement me")
}

func (f flashPromotionSessionRepo) UpdateFlashPromotionSession(ctx context.Context, session *sms.FlashPromotionSession) error {
	panic("implement me")
}

func (f flashPromotionSessionRepo) ListFlashPromotionSession(ctx context.Context, req *sms.FlashPromotionSessionListReq) ([]*sms.FlashPromotionSession, error) {
	panic("implement me")
}

func (f flashPromotionSessionRepo) DeleteFlashPromotionSession(ctx context.Context, id int64) error {
	panic("implement me")
}
