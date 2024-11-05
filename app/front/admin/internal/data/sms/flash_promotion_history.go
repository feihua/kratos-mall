package sms

import (
	"context"
	smsV1 "github.com/feihua/kratos-mall/api/sms/v1"
	"github.com/feihua/kratos-mall/app/front/admin/internal/biz/sms"
	"github.com/feihua/kratos-mall/app/front/admin/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
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

func (f flashPromotionHistoryRepo) ListFlashPromotionHistory(ctx context.Context, req *sms.FlashPromotionHistoryListReq) (*sms.FlashPromotionHistoryListResp, error) {
	list, _ := f.data.SmsClient.FlashPromotionLogList(ctx, &smsV1.FlashPromotionLogListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*sms.FlashPromotionHistory, 0)
	copier.Copy(&orders, &list.List)

	return &sms.FlashPromotionHistoryListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (f flashPromotionHistoryRepo) DeleteFlashPromotionHistory(ctx context.Context, id int64) error {
	panic("implement me")
}
