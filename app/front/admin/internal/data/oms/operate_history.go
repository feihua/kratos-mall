package oms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/front/admin/internal/biz/oms"
	"kratos-mall/app/front/admin/internal/data"
)

type operateHistoryRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewOperateHistoryRepo(logger log.Logger) oms.OperateHistoryRepo {
	return &operateHistoryRepo{
		log: log.NewHelper(logger),
	}
}

func (o operateHistoryRepo) CreateOperateHistory(ctx context.Context, history *oms.OperateHistory) error {
	panic("implement me")
}

func (o operateHistoryRepo) GetOperateHistory(ctx context.Context, id int64) error {
	panic("implement me")
}

func (o operateHistoryRepo) UpdateOperateHistory(ctx context.Context, history *oms.OperateHistory) error {
	panic("implement me")
}

func (o operateHistoryRepo) ListOperateHistory(ctx context.Context, req *oms.OperateHistoryListReq) ([]*oms.OperateHistory, error) {
	panic("implement me")
}

func (o operateHistoryRepo) DeleteOperateHistory(ctx context.Context, id int64) error {
	panic("implement me")
}
