package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/pms/internal/biz"
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

func (v vertifyRecordRepo) ListVertifyRecord(ctx context.Context, req *biz.VertifyRecordListReq) ([]*biz.VertifyRecord, error) {
	panic("implement me")
}

func (v vertifyRecordRepo) DeleteVertifyRecord(ctx context.Context, id int64) error {
	panic("implement me")
}
