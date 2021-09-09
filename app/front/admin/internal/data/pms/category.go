package pms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/front/admin/internal/biz/pms"
	"kratos-mall/app/front/admin/internal/data"
)

type categoryRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewCategoryRepo(data *data.Data, logger log.Logger) pms.CategoryRepo {
	return &categoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c categoryRepo) CreateCategory(ctx context.Context, category *pms.Category) error {
	panic("implement me")
}

func (c categoryRepo) GetCategory(ctx context.Context, id int64) error {
	panic("implement me")
}

func (c categoryRepo) UpdateCategory(ctx context.Context, category *pms.Category) error {
	panic("implement me")
}

func (c categoryRepo) ListCategory(ctx context.Context, req *pms.CategoryListReq) ([]*pms.Category, error) {
	panic("implement me")
}

func (c categoryRepo) DeleteCategory(ctx context.Context, id int64) error {
	panic("implement me")
}
