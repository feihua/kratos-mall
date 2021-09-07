package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/ums/internal/biz"
)

type changeHistoryRepo struct {
	data *Data
	log  *log.Helper
}

// NewChangeHistoryRepo .
func NewChangeHistoryRepo(data *Data, logger log.Logger) biz.ChangeHistoryRepo {
	return &changeHistoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c changeHistoryRepo) CreateChangeHistory(ctx context.Context, history *biz.ChangeHistory) error {
	panic("implement me")
}

func (c changeHistoryRepo) GetChangeHistory(ctx context.Context, id int64) error {
	panic("implement me")
}

func (c changeHistoryRepo) UpdateChangeHistory(ctx context.Context, history *biz.ChangeHistory) error {
	panic("implement me")
}

func (c changeHistoryRepo) ListChangeHistory(ctx context.Context, req *biz.ChangeHistoryListReq) ([]*biz.ChangeHistory, error) {
	panic("implement me")
}

func (c changeHistoryRepo) DeleteChangeHistory(ctx context.Context, id int64) error {
	panic("implement me")
}
