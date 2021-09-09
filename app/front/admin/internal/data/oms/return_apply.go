package oms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/front/admin/internal/biz/oms"
	"kratos-mall/app/front/admin/internal/data"
)

type returnApplyRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewReturnApplyRepo(data *data.Data, logger log.Logger) oms.ReturnApplyRepo {
	return &returnApplyRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r returnApplyRepo) CreateReturnApply(ctx context.Context, apply *oms.ReturnApply) error {
	panic("implement me")
}

func (r returnApplyRepo) GetReturnApply(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r returnApplyRepo) UpdateReturnApply(ctx context.Context, apply *oms.ReturnApply) error {
	panic("implement me")
}

func (r returnApplyRepo) ListReturnApply(ctx context.Context, req *oms.ReturnApplyListReq) ([]*oms.ReturnApply, error) {
	panic("implement me")
}

func (r returnApplyRepo) DeleteReturnApply(ctx context.Context, id int64) error {
	panic("implement me")
}
