package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/ums/internal/biz"
)

type memberAddressRepo struct {
	data *Data
	log  *log.Helper
}

func NewMemberAddressRepo(data *Data, logger log.Logger) biz.MemberAddressRepo {
	return &memberAddressRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (m memberAddressRepo) CreateMemberAddress(ctx context.Context, address *biz.MemberAddress) error {
	panic("implement me")
}

func (m memberAddressRepo) GetMemberAddress(ctx context.Context, id int64) error {
	panic("implement me")
}

func (m memberAddressRepo) UpdateMemberAddress(ctx context.Context, address *biz.MemberAddress) error {
	panic("implement me")
}

func (m memberAddressRepo) ListMemberAddress(ctx context.Context, req *biz.MemberAddressListReq) ([]*biz.MemberAddress, error) {
	panic("implement me")
}

func (m memberAddressRepo) DeleteMemberAddress(ctx context.Context, id int64) error {
	panic("implement me")
}
