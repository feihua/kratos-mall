package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/pms/internal/biz"
)

type memberPriceRepo struct {
	data *Data
	log  *log.Helper
}

func NewMemberPriceRepo(data *Data, logger log.Logger) biz.MemberPriceRepo {
	return &memberPriceRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (m memberPriceRepo) CreateMemberPrice(ctx context.Context, price *biz.MemberPrice) error {
	panic("implement me")
}

func (m memberPriceRepo) GetMemberPrice(ctx context.Context, id int64) error {
	panic("implement me")
}

func (m memberPriceRepo) UpdateMemberPrice(ctx context.Context, price *biz.MemberPrice) error {
	panic("implement me")
}

func (m memberPriceRepo) ListMemberPrice(ctx context.Context, req *biz.MemberPriceListReq) ([]*biz.MemberPrice, error) {
	panic("implement me")
}

func (m memberPriceRepo) DeleteMemberPrice(ctx context.Context, id int64) error {
	panic("implement me")
}
