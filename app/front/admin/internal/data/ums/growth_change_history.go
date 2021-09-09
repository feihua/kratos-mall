package ums

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/front/admin/internal/biz/ums"
	"kratos-mall/app/front/admin/internal/data"
)

type changeHistoryRepo struct {
	data *data.Data
	log  *log.Helper
}

// NewChangeHistoryRepo .
func NewChangeHistoryRepo(data *data.Data, logger log.Logger) ums.ChangeHistoryRepo {
	return &changeHistoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c changeHistoryRepo) CreateChangeHistory(ctx context.Context, history *ums.ChangeHistory) error {
	panic("implement me")
}

func (c changeHistoryRepo) GetChangeHistory(ctx context.Context, id int64) error {
	panic("implement me")
}

func (c changeHistoryRepo) UpdateChangeHistory(ctx context.Context, history *ums.ChangeHistory) error {
	panic("implement me")
}

func (c changeHistoryRepo) ListChangeHistory(ctx context.Context, req *ums.ChangeHistoryListReq) ([]*ums.ChangeHistory, error) {
	panic("implement me")
}

func (c changeHistoryRepo) DeleteChangeHistory(ctx context.Context, id int64) error {
	panic("implement me")
}
