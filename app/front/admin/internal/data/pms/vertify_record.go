package pms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	pmsV1 "kratos-mall/api/pms/v1"
	"kratos-mall/app/front/admin/internal/biz/pms"
	"kratos-mall/app/front/admin/internal/data"
)

type vertifyRecordRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewVertifyRecordRepo(data *data.Data, logger log.Logger) pms.VertifyRecordRepo {
	return &vertifyRecordRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (v vertifyRecordRepo) CreateVertifyRecord(ctx context.Context, record *pms.VertifyRecord) error {
	panic("implement me")
}

func (v vertifyRecordRepo) GetVertifyRecord(ctx context.Context, id int64) error {
	panic("implement me")
}

func (v vertifyRecordRepo) UpdateVertifyRecord(ctx context.Context, record *pms.VertifyRecord) error {
	panic("implement me")
}

func (v vertifyRecordRepo) ListVertifyRecord(ctx context.Context, req *pms.VertifyRecordListReq) (*pms.VertifyRecordListResp, error) {
	list, _ := v.data.PmsClient.ProductVertifyRecordList(ctx, &pmsV1.ProductVertifyRecordListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*pms.VertifyRecord, 0)
	copier.Copy(&orders, &list.List)

	return &pms.VertifyRecordListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (v vertifyRecordRepo) DeleteVertifyRecord(ctx context.Context, id int64) error {
	panic("implement me")
}
