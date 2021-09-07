package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/pms/internal/biz"
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

func (b brandRepo) ListBrand(ctx context.Context, req *biz.BrandListReq) ([]*biz.Brand, error) {
	panic("implement me")
}

func (b brandRepo) DeleteBrand(ctx context.Context, id int64) error {
	panic("implement me")
}
