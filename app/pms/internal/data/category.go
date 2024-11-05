package data

import (
	"context"
	"github.com/feihua/kratos-mall/app/pms/internal/biz"
	"github.com/feihua/kratos-mall/app/pms/internal/data/model"
	"github.com/feihua/kratos-mall/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
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

func (c categoryRepo) ListCategory(ctx context.Context, req *biz.CategoryListReq) (*biz.CategoryListResp, error) {
	var all []model.PmsProductCategory
	result := c.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	c.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.Category, 0)

	for _, item := range all {
		list = append(list, &biz.Category{
			Id:           item.Id,
			ParentId:     item.ParentId,
			Name:         item.Name,
			Level:        item.Level,
			ProductCount: item.ProductCount,
			ProductUnit:  item.ProductUnit,
			NavStatus:    item.NavStatus,
			ShowStatus:   item.ShowStatus,
			Sort:         item.Sort,
			Icon:         item.Icon,
			Keywords:     item.Keywords,
			Description:  item.Description,
		})
	}

	return &biz.CategoryListResp{
		Total: count,
		List:  list,
	}, nil
}

func (c categoryRepo) DeleteCategory(ctx context.Context, id int64) error {
	panic("implement me")
}
