package data

import (
	"context"
	"github.com/feihua/kratos-mall/app/sms/internal/biz"
	"github.com/feihua/kratos-mall/app/sms/internal/data/model"
	"github.com/feihua/kratos-mall/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
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

func (f flashPromotionRepo) ListFlashPromotion(ctx context.Context, req *biz.FlashPromotionListReq) (*biz.FlashPromotionListResp, error) {
	var all []model.SmsFlashPromotion
	result := f.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	f.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.FlashPromotion, 0)

	for _, item := range all {
		list = append(list, &biz.FlashPromotion{
			Id:         item.Id,
			Title:      item.Title,
			StartDate:  item.StartDate.Format("2006-01-02 15:04:05"),
			EndDate:    item.EndDate.Format("2006-01-02 15:04:05"),
			Status:     item.Status,
			CreateTime: item.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &biz.FlashPromotionListResp{
		Total: count,
		List:  list,
	}, nil
}

func (f flashPromotionRepo) DeleteFlashPromotion(ctx context.Context, id int64) error {
	panic("implement me")
}
