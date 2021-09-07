package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/oms/internal/biz"
)

type returnReasonRepo struct {
	data *Data
	log  *log.Helper
}

func NewReturnReasonRepo(data *Data, logger log.Logger) biz.ReturnReasonRepo {
	return &returnReasonRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r returnReasonRepo) CreateReturnReason(ctx context.Context, reason *biz.ReturnReason) error {
	panic("implement me")
}

func (r returnReasonRepo) GetReturnReason(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r returnReasonRepo) UpdateReturnReason(ctx context.Context, reason *biz.ReturnReason) error {
	panic("implement me")
}

func (r returnReasonRepo) ListReturnReason(ctx context.Context, req *biz.ReturnReasonListReq) ([]*biz.ReturnReason, error) {
	panic("implement me")
}

func (r returnReasonRepo) DeleteReturnReason(ctx context.Context, id int64) error {
	panic("implement me")
}
