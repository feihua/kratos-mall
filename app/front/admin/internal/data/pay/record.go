package pay

import (
	"context"
	"github.com/feihua/kratos-mall/app/front/admin/internal/biz/pay"
	"github.com/feihua/kratos-mall/app/front/admin/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

type recordRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewRecordRepo(data *data.Data, logger log.Logger) pay.RecordRepo {
	return &recordRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r recordRepo) CreateRecord(ctx context.Context, record *pay.Record) error {
	panic("implement me")
}

func (r recordRepo) GetRecord(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r recordRepo) UpdateRecord(ctx context.Context, record *pay.Record) error {
	panic("implement me")
}

func (r recordRepo) ListRecord(ctx context.Context, req *pay.RecordListReq) (*pay.RecordListResp, error) {
	panic("implement me")
}

func (r recordRepo) DeleteRecord(ctx context.Context, id int64) error {
	panic("implement me")
}
