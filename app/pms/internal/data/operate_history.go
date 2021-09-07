package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/pms/internal/biz"
)

type operateHistoryRepo struct {
	data *Data
	log  *log.Helper
}

func NewOperateHistoryRepo(data *Data, logger log.Logger) biz.OperateHistoryRepo {
	return &operateHistoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (o operateHistoryRepo) CreateOperateHistory(ctx context.Context, history *biz.OperateHistory) error {
	panic("implement me")
}

func (o operateHistoryRepo) GetOperateHistory(ctx context.Context, id int64) error {
	panic("implement me")
}

func (o operateHistoryRepo) UpdateOperateHistory(ctx context.Context, history *biz.OperateHistory) error {
	panic("implement me")
}

func (o operateHistoryRepo) ListOperateHistory(ctx context.Context, req *biz.OperateHistoryListReq) ([]*biz.OperateHistory, error) {
	panic("implement me")
}

func (o operateHistoryRepo) DeleteOperateHistory(ctx context.Context, id int64) error {
	panic("implement me")
}
