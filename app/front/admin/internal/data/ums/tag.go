package ums

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/front/admin/internal/biz/ums"
	"kratos-mall/app/front/admin/internal/data"
)

type tagRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewTagRepo(data *data.Data, logger log.Logger) ums.TagRepo {
	return &tagRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (t tagRepo) CreateTag(ctx context.Context, tag *ums.Tag) error {
	panic("implement me")
}

func (t tagRepo) GetTag(ctx context.Context, id int64) error {
	panic("implement me")
}

func (t tagRepo) UpdateTag(ctx context.Context, tag *ums.Tag) error {
	panic("implement me")
}

func (t tagRepo) ListTag(ctx context.Context, req *ums.TagListReq) ([]*ums.Tag, error) {
	panic("implement me")
}

func (t tagRepo) DeleteTag(ctx context.Context, id int64) error {
	panic("implement me")
}
