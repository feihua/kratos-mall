package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/pms/internal/biz"
)

type categoryRepo struct {
	data *Data
	log  *log.Helper
}

func NewCategoryRepo(data *Data, logger log.Logger) biz.CategoryRepo {
	return &categoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c categoryRepo) CreateCategory(ctx context.Context, category *biz.Category) error {
	panic("implement me")
}

func (c categoryRepo) GetCategory(ctx context.Context, id int64) error {
	panic("implement me")
}

func (c categoryRepo) UpdateCategory(ctx context.Context, category *biz.Category) error {
	panic("implement me")
}

func (c categoryRepo) ListCategory(ctx context.Context, req *biz.CategoryListReq) ([]*biz.Category, error) {
	panic("implement me")
}

func (c categoryRepo) DeleteCategory(ctx context.Context, id int64) error {
	panic("implement me")
}
