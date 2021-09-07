package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/pay/internal/biz"
)

type recordRepo struct {
	data *Data
	log  *log.Helper
}

func NewRecordRepo(data *Data, logger log.Logger) biz.RecordRepo {
	return &recordRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r recordRepo) CreateRecord(ctx context.Context, record *biz.Record) error {
	panic("implement me")
}

func (r recordRepo) GetRecord(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r recordRepo) UpdateRecord(ctx context.Context, record *biz.Record) error {
	panic("implement me")
}

func (r recordRepo) ListRecord(ctx context.Context, req *biz.RecordListReq) ([]*biz.Record, error) {
	panic("implement me")
}

func (r recordRepo) DeleteRecord(ctx context.Context, id int64) error {
	panic("implement me")
}
