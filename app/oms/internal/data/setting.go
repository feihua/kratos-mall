package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/oms/internal/biz"
	"kratos-mall/app/oms/internal/data/model"
	"kratos-mall/pkg/util/pagination"
)

type settingRepo struct {
	data *Data
	log  *log.Helper
}

func NewSettingRepo(data *Data, logger log.Logger) biz.SettingRepo {
	return &settingRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (s settingRepo) CreateSetting(ctx context.Context, setting *biz.Setting) error {
	panic("implement me")
}

func (s settingRepo) GetSetting(ctx context.Context, id int64) error {
	panic("implement me")
}

func (s settingRepo) UpdateSetting(ctx context.Context, setting *biz.Setting) error {
	panic("implement me")
}

func (s settingRepo) ListSetting(ctx context.Context, req *biz.SettingListReq) (*biz.SettingListResp, error) {
	var all []model.OmsOrderSetting
	result := s.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	s.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.Setting, 0)

	for _, item := range all {
		list = append(list, &biz.Setting{
			Id:                  item.Id,
			FlashOrderOvertime:  item.FinishOvertime,
			NormalOrderOvertime: item.NormalOrderOvertime,
			ConfirmOvertime:     item.ConfirmOvertime,
			FinishOvertime:      item.FinishOvertime,
			CommentOvertime:     item.CommentOvertime,
		})
	}

	return &biz.SettingListResp{
		Total: count,
		List:  list,
	}, nil
}

func (s settingRepo) DeleteSetting(ctx context.Context, id int64) error {
	panic("implement me")
}
