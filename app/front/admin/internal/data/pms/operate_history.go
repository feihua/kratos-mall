package pms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/front/admin/internal/biz/pms"
	"kratos-mall/app/front/admin/internal/data"
)

type operateHistoryRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewOperateHistoryRepo(data *data.Data, logger log.Logger) pms.OperateHistoryRepo {
	return &operateHistoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (o operateHistoryRepo) CreateOperateHistory(ctx context.Context, history *pms.OperateHistory) error {
	panic("implement me")
}

func (o operateHistoryRepo) GetOperateHistory(ctx context.Context, id int64) error {
	panic("implement me")
}

func (o operateHistoryRepo) UpdateOperateHistory(ctx context.Context, history *pms.OperateHistory) error {
	panic("implement me")
}

func (o operateHistoryRepo) ListOperateHistory(ctx context.Context, req *pms.OperateHistoryListReq) ([]*pms.OperateHistory, error) {
	panic("implement me")
}

func (o operateHistoryRepo) DeleteOperateHistory(ctx context.Context, id int64) error {
	panic("implement me")
}
