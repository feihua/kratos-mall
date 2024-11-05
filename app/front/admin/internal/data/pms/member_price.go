package pms

import (
	"context"
	pmsV1 "github.com/feihua/kratos-mall/api/pms/v1"
	"github.com/feihua/kratos-mall/app/front/admin/internal/biz/pms"
	"github.com/feihua/kratos-mall/app/front/admin/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
)

type memberPriceRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewMemberPriceRepo(data *data.Data, logger log.Logger) pms.MemberPriceRepo {
	return &memberPriceRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (m memberPriceRepo) CreateMemberPrice(ctx context.Context, price *pms.MemberPrice) error {
	panic("implement me")
}

func (m memberPriceRepo) GetMemberPrice(ctx context.Context, id int64) error {
	panic("implement me")
}

func (m memberPriceRepo) UpdateMemberPrice(ctx context.Context, price *pms.MemberPrice) error {
	panic("implement me")
}

func (m memberPriceRepo) ListMemberPrice(ctx context.Context, req *pms.MemberPriceListReq) (*pms.MemberPriceListResp, error) {
	list, _ := m.data.PmsClient.MemberPriceList(ctx, &pmsV1.MemberPriceListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*pms.MemberPrice, 0)
	copier.Copy(&orders, &list.List)

	return &pms.MemberPriceListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (m memberPriceRepo) DeleteMemberPrice(ctx context.Context, id int64) error {
	panic("implement me")
}
