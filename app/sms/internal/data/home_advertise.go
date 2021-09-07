package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/sms/internal/biz"
)

type homeAdvertiseRepo struct {
	data *Data
	log  *log.Helper
}

func NewHomeAdvertiseRepo(data *Data, logger log.Logger) biz.HomeAdvertiseRepo {
	return &homeAdvertiseRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (h homeAdvertiseRepo) CreateHomeAdvertise(ctx context.Context, advertise *biz.HomeAdvertise) error {
	panic("implement me")
}

func (h homeAdvertiseRepo) GetHomeAdvertise(ctx context.Context, id int64) error {
	panic("implement me")
}

func (h homeAdvertiseRepo) UpdateHomeAdvertise(ctx context.Context, advertise *biz.HomeAdvertise) error {
	panic("implement me")
}

func (h homeAdvertiseRepo) ListHomeAdvertise(ctx context.Context, req *biz.HomeAdvertiseListReq) ([]*biz.HomeAdvertise, error) {
	panic("implement me")
}

func (h homeAdvertiseRepo) DeleteHomeAdvertise(ctx context.Context, id int64) error {
	panic("implement me")
}
