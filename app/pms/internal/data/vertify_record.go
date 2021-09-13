package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/pms/internal/biz"
	"kratos-mall/app/pms/internal/data/model"
	"kratos-mall/pkg/util/pagination"
)

type vertifyRecordRepo struct {
	data *Data
	log  *log.Helper
}

func NewVertifyRecordRepo(data *Data, logger log.Logger) biz.VertifyRecordRepo {
	return &vertifyRecordRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (v vertifyRecordRepo) CreateVertifyRecord(ctx context.Context, record *biz.VertifyRecord) error {
	panic("implement me")
}

func (v vertifyRecordRepo) GetVertifyRecord(ctx context.Context, id int64) error {
	panic("implement me")
}

func (v vertifyRecordRepo) UpdateVertifyRecord(ctx context.Context, record *biz.VertifyRecord) error {
	panic("implement me")
}

func (v vertifyRecordRepo) ListVertifyRecord(ctx context.Context, req *biz.VertifyRecordListReq) (*biz.VertifyRecordListResp, error) {
	var all []model.PmsProductVertifyRecord
	result := v.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	v.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.VertifyRecord, 0)

	for _, item := range all {
		list = append(list, &biz.VertifyRecord{
			Id:         item.Id,
			ProductId:  item.ProductId,
			CreateTime: item.CreateTime.Format("2006-01-02 15:04:05"),
			VertifyMan: item.VertifyMan,
			Status:     item.Status,
			Detail:     item.Detail,
		})
	}

	return &biz.VertifyRecordListResp{
		Total: count,
		List:  list,
	}, nil
}

func (v vertifyRecordRepo) DeleteVertifyRecord(ctx context.Context, id int64) error {
	panic("implement me")
}
