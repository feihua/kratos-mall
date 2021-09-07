package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/ums/internal/biz"
)

type tagRepo struct {
	data *Data
	log  *log.Helper
}

func NewTagRepo(data *Data, logger log.Logger) biz.TagRepo {
	return &tagRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (t tagRepo) CreateTag(ctx context.Context, tag *biz.Tag) error {
	panic("implement me")
}

func (t tagRepo) GetTag(ctx context.Context, id int64) error {
	panic("implement me")
}

func (t tagRepo) UpdateTag(ctx context.Context, tag *biz.Tag) error {
	panic("implement me")
}

func (t tagRepo) ListTag(ctx context.Context, req *biz.TagListReq) ([]*biz.Tag, error) {
	panic("implement me")
}

func (t tagRepo) DeleteTag(ctx context.Context, id int64) error {
	panic("implement me")
}
