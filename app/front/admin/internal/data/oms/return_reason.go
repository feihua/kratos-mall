package oms

import (
	"context"
	omsV1 "github.com/feihua/kratos-mall/api/oms/v1"
	"github.com/feihua/kratos-mall/app/front/admin/internal/biz/oms"
	"github.com/feihua/kratos-mall/app/front/admin/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
)

type returnReasonRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewReturnReasonRepo(data *data.Data, logger log.Logger) oms.ReturnReasonRepo {
	return &returnReasonRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r returnReasonRepo) CreateReturnReason(ctx context.Context, reason *oms.ReturnReason) error {
	panic("implement me")
}

func (r returnReasonRepo) GetReturnReason(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r returnReasonRepo) UpdateReturnReason(ctx context.Context, reason *oms.ReturnReason) error {
	panic("implement me")
}

func (r returnReasonRepo) ListReturnReason(ctx context.Context, req *oms.ReturnReasonListReq) (*oms.ReturnReasonListResp, error) {
	list, _ := r.data.OmsClient.OrderReturnReasonList(ctx, &omsV1.OrderReturnReasonListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*oms.ReturnReason, 0)
	copier.Copy(&orders, &list.List)
	return &oms.ReturnReasonListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (r returnReasonRepo) DeleteReturnReason(ctx context.Context, id int64) error {
	panic("implement me")
}
