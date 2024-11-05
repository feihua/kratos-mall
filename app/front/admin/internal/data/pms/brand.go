package pms

import (
	"context"
	pmsV1 "github.com/feihua/kratos-mall/api/pms/v1"
	"github.com/feihua/kratos-mall/app/front/admin/internal/biz/pms"
	"github.com/feihua/kratos-mall/app/front/admin/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
)

type brandRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewBrandRepo(data *data.Data, logger log.Logger) pms.BrandRepo {
	return &brandRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (b brandRepo) CreateBrand(ctx context.Context, brand *pms.Brand) error {
	panic("implement me")
}

func (b brandRepo) GetBrand(ctx context.Context, id int64) error {
	panic("implement me")
}

func (b brandRepo) UpdateBrand(ctx context.Context, brand *pms.Brand) error {
	panic("implement me")
}

func (b brandRepo) ListBrand(ctx context.Context, req *pms.BrandListReq) (*pms.BrandListResp, error) {
	list, _ := b.data.PmsClient.BrandList(ctx, &pmsV1.BrandListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*pms.Brand, 0)
	copier.Copy(&orders, &list.List)
	return &pms.BrandListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (b brandRepo) DeleteBrand(ctx context.Context, id int64) error {
	panic("implement me")
}
