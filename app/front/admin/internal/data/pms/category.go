package pms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	pmsV1 "kratos-mall/api/pms/v1"
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

func (c categoryRepo) ListCategory(ctx context.Context, req *pms.CategoryListReq) (*pms.CategoryListResp, error) {
	list, _ := c.data.PmsClient.ProductCategoryList(ctx, &pmsV1.ProductCategoryListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*pms.Category, 0)
	copier.Copy(&orders, &list.List)
	return &pms.CategoryListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (c categoryRepo) DeleteCategory(ctx context.Context, id int64) error {
	panic("implement me")
}
