package data

import (
	"context"
	"github.com/feihua/kratos-mall/app/ums/internal/biz"
	"github.com/feihua/kratos-mall/app/ums/internal/data/model"
	"github.com/feihua/kratos-mall/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
)

type integrationChangeHistoryRepo struct {
	data *Data
	log  *log.Helper
}

func NewIntegrationChangeHistoryRepo(data *Data, logger log.Logger) biz.IntegrationChangeHistoryRepo {
	return &integrationChangeHistoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (i integrationChangeHistoryRepo) CreateIntegrationChangeHistory(ctx context.Context, history *biz.IntegrationChangeHistory) error {
	panic("implement me")
}

func (i integrationChangeHistoryRepo) GetIntegrationChangeHistory(ctx context.Context, id int64) error {
	panic("implement me")
}

func (i integrationChangeHistoryRepo) UpdateIntegrationChangeHistory(ctx context.Context, history *biz.IntegrationChangeHistory) error {
	panic("implement me")
}

func (i integrationChangeHistoryRepo) ListIntegrationChangeHistory(ctx context.Context, req *biz.IntegrationChangeHistoryListReq) (*biz.IntegrationChangeHistoryListResp, error) {
	var all []model.UmsIntegrationChangeHistory
	result := i.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	i.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.IntegrationChangeHistory, 0)

	for _, item := range all {
		list = append(list, &biz.IntegrationChangeHistory{
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

	return &biz.IntegrationChangeHistoryListResp{
		Total: count,
		List:  list,
	}, nil
}

func (i integrationChangeHistoryRepo) DeleteIntegrationChangeHistory(ctx context.Context, id int64) error {
	panic("implement me")
}
