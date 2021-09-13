package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/sms/internal/biz"
	"kratos-mall/app/sms/internal/data/model"
	"kratos-mall/pkg/util/pagination"
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

func (f flashPromotionSessionRepo) ListFlashPromotionSession(ctx context.Context, req *biz.FlashPromotionSessionListReq) (*biz.FlashPromotionSessionListResp, error) {
	var all []model.SmsFlashPromotionSession
	result := f.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	f.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.FlashPromotionSession, 0)

	for _, item := range all {
		list = append(list, &biz.FlashPromotionSession{
			Id:         item.Id,
			Name:       item.Name,
			StartTime:  item.StartTime.Format("2006-01-02 15:04:05"),
			EndTime:    item.EndTime.Format("2006-01-02 15:04:05"),
			Status:     item.Status,
			CreateTime: item.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &biz.FlashPromotionSessionListResp{
		Total: count,
		List:  list,
	}, nil
}

func (f flashPromotionSessionRepo) DeleteFlashPromotionSession(ctx context.Context, id int64) error {
	panic("implement me")
}
