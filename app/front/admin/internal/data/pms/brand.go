package pms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/front/admin/internal/biz/pms"
	"kratos-mall/app/front/admin/internal/data"
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

func (b brandRepo) ListBrand(ctx context.Context, req *pms.BrandListReq) ([]*pms.Brand, error) {
	panic("implement me")
}

func (b brandRepo) DeleteBrand(ctx context.Context, id int64) error {
	panic("implement me")
}
