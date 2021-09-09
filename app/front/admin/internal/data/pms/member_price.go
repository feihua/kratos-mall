package pms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
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

func (m memberPriceRepo) ListMemberPrice(ctx context.Context, req *pms.MemberPriceListReq) ([]*pms.MemberPrice, error) {
	panic("implement me")
}

func (m memberPriceRepo) DeleteMemberPrice(ctx context.Context, id int64) error {
	panic("implement me")
}
