package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/oms/internal/biz"
)

type returnApplyRepo struct {
	data *Data
	log  *log.Helper
}

func NewReturnApplyRepo(data *Data, logger log.Logger) biz.ReturnApplyRepo {
	return &returnApplyRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r returnApplyRepo) CreateReturnApply(ctx context.Context, apply *biz.ReturnApply) error {
	panic("implement me")
}

func (r returnApplyRepo) GetReturnApply(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r returnApplyRepo) UpdateReturnApply(ctx context.Context, apply *biz.ReturnApply) error {
	panic("implement me")
}

func (r returnApplyRepo) ListReturnApply(ctx context.Context, req *biz.ReturnApplyListReq) ([]*biz.ReturnApply, error) {
	panic("implement me")
}

func (r returnApplyRepo) DeleteReturnApply(ctx context.Context, id int64) error {
	panic("implement me")
}
