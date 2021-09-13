package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/ums/internal/biz"
	"kratos-mall/app/ums/internal/data/model"
	"kratos-mall/pkg/util/pagination"
)

type consumeSettingRepo struct {
	data *Data
	log  *log.Helper
}

func NewConsumeSettingRepo(data *Data, logger log.Logger) biz.ConsumeSettingRepo {
	return &consumeSettingRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c consumeSettingRepo) CreateConsumeSetting(ctx context.Context, setting *biz.ConsumeSetting) error {
	panic("implement me")
}

func (c consumeSettingRepo) GetConsumeSetting(ctx context.Context, id int64) error {
	panic("implement me")
}

func (c consumeSettingRepo) UpdateConsumeSetting(ctx context.Context, setting *biz.ConsumeSetting) error {
	panic("implement me")
}

func (c consumeSettingRepo) ListConsumeSetting(ctx context.Context, req *biz.ConsumeSettingListReq) (*biz.ConsumeSettingListResp, error) {
	var all []model.UmsIntegrationConsumeSetting
	result := c.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	c.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.ConsumeSetting, 0)

	for _, item := range all {
		list = append(list, &biz.ConsumeSetting{
			Id:                 item.Id,
			DeductionPerAmount: item.DeductionPerAmount,
			MaxPercentPerOrder: item.MaxPercentPerOrder,
			UseUnit:            item.UseUnit,
			CouponStatus:       item.CouponStatus,
		})
	}

	return &biz.ConsumeSettingListResp{
		Total: count,
		List:  list,
	}, nil
}

func (c consumeSettingRepo) DeleteConsumeSetting(ctx context.Context, id int64) error {
	panic("implement me")
}
