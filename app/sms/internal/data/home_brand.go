package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/sms/internal/biz"
)

type homeBrandRepo struct {
	data *Data
	log  *log.Helper
}

func NewHomeBrandRepo(data *Data, logger log.Logger) biz.HomeBrandRepo {
	return &homeBrandRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (h homeBrandRepo) CreateHomeBrand(ctx context.Context, brand *biz.HomeBrand) error {
	panic("implement me")
}

func (h homeBrandRepo) GetHomeBrand(ctx context.Context, id int64) error {
	panic("implement me")
}

func (h homeBrandRepo) UpdateHomeBrand(ctx context.Context, brand *biz.HomeBrand) error {
	panic("implement me")
}

func (h homeBrandRepo) ListHomeBrand(ctx context.Context, req *biz.HomeBrandListReq) ([]*biz.HomeBrand, error) {
	panic("implement me")
}

func (h homeBrandRepo) DeleteHomeBrand(ctx context.Context, id int64) error {
	panic("implement me")
}
