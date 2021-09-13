package pms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	pmsV1 "kratos-mall/api/pms/v1"
	"kratos-mall/app/front/admin/internal/biz/pms"
	"kratos-mall/app/front/admin/internal/data"
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
