package oms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/front/admin/internal/biz/oms"
	"kratos-mall/app/front/admin/internal/data"
)

type returnReasonRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewReturnReasonRepo(data *data.Data, logger log.Logger) oms.ReturnReasonRepo {
	return &returnReasonRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r returnReasonRepo) CreateReturnReason(ctx context.Context, reason *oms.ReturnReason) error {
	panic("implement me")
}

func (r returnReasonRepo) GetReturnReason(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r returnReasonRepo) UpdateReturnReason(ctx context.Context, reason *oms.ReturnReason) error {
	panic("implement me")
}

func (r returnReasonRepo) ListReturnReason(ctx context.Context, req *oms.ReturnReasonListReq) ([]*oms.ReturnReason, error) {
	panic("implement me")
}

func (r returnReasonRepo) DeleteReturnReason(ctx context.Context, id int64) error {
	panic("implement me")
}
