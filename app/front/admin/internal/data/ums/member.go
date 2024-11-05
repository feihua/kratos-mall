package ums

import (
	"context"
	umsV1 "github.com/feihua/kratos-mall/api/ums/v1"
	"github.com/feihua/kratos-mall/app/front/admin/internal/biz/ums"
	"github.com/feihua/kratos-mall/app/front/admin/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
)

type memberRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewMemberRepo(data *data.Data, logger log.Logger) ums.MemberRepo {
	return &memberRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (m memberRepo) CreateMember(ctx context.Context, member *ums.Member) error {
	panic("implement me")
}

func (m memberRepo) GetMember(ctx context.Context, id int64) error {
	panic("implement me")
}

func (m memberRepo) UpdateMember(ctx context.Context, member *ums.Member) error {
	panic("implement me")
}

func (m memberRepo) ListMember(ctx context.Context, req *ums.MemberListReq) (*ums.MemberListResp, error) {
	list, _ := m.data.UmsClient.MemberList(ctx, &umsV1.MemberListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*ums.Member, 0)
	copier.Copy(&orders, &list.List)

	return &ums.MemberListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (m memberRepo) DeleteMember(ctx context.Context, id int64) error {
	panic("implement me")
}
