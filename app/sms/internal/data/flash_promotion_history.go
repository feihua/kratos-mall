package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/sms/internal/biz"
	"kratos-mall/app/sms/internal/data/model"
	"kratos-mall/pkg/util/pagination"
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

func (f flashPromotionHistoryRepo) ListFlashPromotionHistory(ctx context.Context, req *biz.FlashPromotionHistoryListReq) (*biz.FlashPromotionHistoryListResp, error) {
	var all []model.SmsFlashPromotionLog
	result := f.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	f.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.FlashPromotionHistory, 0)

	for _, item := range all {
		list = append(list, &biz.FlashPromotionHistory{
			Id:            item.Id,
			MemberId:      item.MemberId,
			ProductId:     item.ProductId,
			MemberPhone:   item.MemberPhone,
			ProductName:   item.ProductName,
			SubscribeTime: item.SubscribeTime.Format("2006-01-02 15:04:05"),
			SendTime:      item.SendTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &biz.FlashPromotionHistoryListResp{
		Total: count,
		List:  list,
	}, nil
}

func (f flashPromotionHistoryRepo) DeleteFlashPromotionHistory(ctx context.Context, id int64) error {
	panic("implement me")
}
