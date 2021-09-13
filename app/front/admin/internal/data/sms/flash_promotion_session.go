package sms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	smsV1 "kratos-mall/api/sms/v1"
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

func (f flashPromotionSessionRepo) ListFlashPromotionSession(ctx context.Context, req *sms.FlashPromotionSessionListReq) (*sms.FlashPromotionSessionListResp, error) {
	list, _ := f.data.SmsClient.FlashPromotionSessionList(ctx, &smsV1.FlashPromotionSessionListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*sms.FlashPromotionSession, 0)
	copier.Copy(&orders, &list.List)

	return &sms.FlashPromotionSessionListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (f flashPromotionSessionRepo) DeleteFlashPromotionSession(ctx context.Context, id int64) error {
	panic("implement me")
}
