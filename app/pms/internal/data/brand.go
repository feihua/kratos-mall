package data

import (
	"context"
	"github.com/feihua/kratos-mall/app/pms/internal/biz"
	"github.com/feihua/kratos-mall/app/pms/internal/data/model"
	"github.com/feihua/kratos-mall/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
)

type brandRepo struct {
	data *Data
	log  *log.Helper
}

func NewBrandRepo(data *Data, logger log.Logger) biz.BrandRepo {
	return &brandRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (b brandRepo) CreateBrand(ctx context.Context, brand *biz.Brand) error {
	panic("implement me")
}

func (b brandRepo) GetBrand(ctx context.Context, id int64) error {
	panic("implement me")
}

func (b brandRepo) UpdateBrand(ctx context.Context, brand *biz.Brand) error {
	panic("implement me")
}

func (b brandRepo) ListBrand(ctx context.Context, req *biz.BrandListReq) (*biz.BrandListResp, error) {
	var all []model.PmsBrand
	result := b.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	b.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.Brand, 0)

	for _, item := range all {
		list = append(list, &biz.Brand{
			Id:                  item.Id,
			Name:                item.Name,
			FirstLetter:         item.FirstLetter,
			Sort:                item.Sort,
			FactoryStatus:       item.FactoryStatus,
			ShowStatus:          item.ShowStatus,
			ProductCount:        item.ProductCount,
			ProductCommentCount: item.ProductCommentCount,
			Logo:                item.Logo,
			BigPic:              item.BigPic,
			BrandStory:          item.BrandStory,
		})
	}

	return &biz.BrandListResp{
		Total: count,
		List:  list,
	}, nil
}

func (b brandRepo) DeleteBrand(ctx context.Context, id int64) error {
	panic("implement me")
}
