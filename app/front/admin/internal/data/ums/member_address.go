package ums

import (
	"context"
	umsV1 "github.com/feihua/kratos-mall/api/ums/v1"
	"github.com/feihua/kratos-mall/app/front/admin/internal/biz/ums"
	"github.com/feihua/kratos-mall/app/front/admin/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
)

type memberAddressRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewMemberAddressRepo(data *data.Data, logger log.Logger) ums.MemberAddressRepo {
	return &memberAddressRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (m memberAddressRepo) CreateMemberAddress(ctx context.Context, address *ums.MemberAddress) error {
	panic("implement me")
}

func (m memberAddressRepo) GetMemberAddress(ctx context.Context, id int64) error {
	panic("implement me")
}

func (m memberAddressRepo) UpdateMemberAddress(ctx context.Context, address *ums.MemberAddress) error {
	panic("implement me")
}

func (m memberAddressRepo) ListMemberAddress(ctx context.Context, req *ums.MemberAddressListReq) (*ums.MemberAddressListResp, error) {
	list, _ := m.data.UmsClient.MemberReceiveAddressList(ctx, &umsV1.MemberReceiveAddressListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*ums.MemberAddress, 0)
	copier.Copy(&orders, &list.List)

	return &ums.MemberAddressListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (m memberAddressRepo) DeleteMemberAddress(ctx context.Context, id int64) error {
	panic("implement me")
}
