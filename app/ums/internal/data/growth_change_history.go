package data

import (
	"context"
	"github.com/feihua/kratos-mall/app/ums/internal/biz"
	"github.com/feihua/kratos-mall/app/ums/internal/data/model"
	"github.com/feihua/kratos-mall/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
)

type changeHistoryRepo struct {
	data *Data
	log  *log.Helper
}

// NewChangeHistoryRepo .
func NewChangeHistoryRepo(data *Data, logger log.Logger) biz.ChangeHistoryRepo {
	return &changeHistoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c changeHistoryRepo) CreateChangeHistory(ctx context.Context, history *biz.ChangeHistory) error {
	panic("implement me")
}

func (c changeHistoryRepo) GetChangeHistory(ctx context.Context, id int64) error {
	panic("implement me")
}

func (c changeHistoryRepo) UpdateChangeHistory(ctx context.Context, history *biz.ChangeHistory) error {
	panic("implement me")
}

func (c changeHistoryRepo) ListChangeHistory(ctx context.Context, req *biz.ChangeHistoryListReq) (*biz.ChangeHistoryListResp, error) {
	var all []model.UmsGrowthChangeHistory
	result := c.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	c.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.ChangeHistory, 0)

	for _, item := range all {
		list = append(list, &biz.ChangeHistory{
			Id:          item.Id,
			MemberId:    item.MemberId,
			CreateTime:  item.CreateTime.Format("2006-01-02 15:04:05"),
			ChangeType:  item.ChangeType,
			ChangeCount: item.ChangeCount,
			OperateMan:  item.OperateMan,
			OperateNote: item.OperateNote,
			SourceType:  item.SourceType,
		})
	}

	return &biz.ChangeHistoryListResp{
		Total: count,
		List:  list,
	}, nil
}

func (c changeHistoryRepo) DeleteChangeHistory(ctx context.Context, id int64) error {
	panic("implement me")
}
