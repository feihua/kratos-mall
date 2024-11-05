package sms

import (
	"context"
	smsV1 "github.com/feihua/kratos-mall/api/sms/v1"
	"github.com/feihua/kratos-mall/app/front/admin/internal/biz/sms"
	"github.com/feihua/kratos-mall/app/front/admin/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
)

type homeAdvertiseRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewHomeAdvertiseRepo(data *data.Data, logger log.Logger) sms.HomeAdvertiseRepo {
	return &homeAdvertiseRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (h homeAdvertiseRepo) CreateHomeAdvertise(ctx context.Context, advertise *sms.HomeAdvertise) error {
	panic("implement me")
}

func (h homeAdvertiseRepo) GetHomeAdvertise(ctx context.Context, id int64) error {
	panic("implement me")
}

func (h homeAdvertiseRepo) UpdateHomeAdvertise(ctx context.Context, advertise *sms.HomeAdvertise) error {
	panic("implement me")
}

func (h homeAdvertiseRepo) ListHomeAdvertise(ctx context.Context, req *sms.HomeAdvertiseListReq) (*sms.HomeAdvertiseListResp, error) {
	list, _ := h.data.SmsClient.HomeAdvertiseList(ctx, &smsV1.HomeAdvertiseListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*sms.HomeAdvertise, 0)
	copier.Copy(&orders, &list.List)

	return &sms.HomeAdvertiseListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (h homeAdvertiseRepo) DeleteHomeAdvertise(ctx context.Context, id int64) error {
	panic("implement me")
}
