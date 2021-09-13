package sms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	smsV1 "kratos-mall/api/sms/v1"
	"kratos-mall/app/front/admin/internal/biz/sms"
	"kratos-mall/app/front/admin/internal/data"
)

type homeBrandRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewHomeBrandRepo(data *data.Data, logger log.Logger) sms.HomeBrandRepo {
	return &homeBrandRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (h homeBrandRepo) CreateHomeBrand(ctx context.Context, brand *sms.HomeBrand) error {
	panic("implement me")
}

func (h homeBrandRepo) GetHomeBrand(ctx context.Context, id int64) error {
	panic("implement me")
}

func (h homeBrandRepo) UpdateHomeBrand(ctx context.Context, brand *sms.HomeBrand) error {
	panic("implement me")
}

func (h homeBrandRepo) ListHomeBrand(ctx context.Context, req *sms.HomeBrandListReq) (*sms.HomeBrandListResp, error) {
	list, _ := h.data.SmsClient.HomeBrandList(ctx, &smsV1.HomeBrandListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*sms.HomeBrand, 0)
	copier.Copy(&orders, &list.List)

	return &sms.HomeBrandListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (h homeBrandRepo) DeleteHomeBrand(ctx context.Context, id int64) error {
	panic("implement me")
}
