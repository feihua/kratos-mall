package ums

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	umsV1 "kratos-mall/api/ums/v1"
	"kratos-mall/app/front/admin/internal/biz/ums"
	"kratos-mall/app/front/admin/internal/data"
)

type memberLevelRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewMemberLevelRepo(data *data.Data, logger log.Logger) ums.MemberLevelRepo {
	return &memberLevelRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (m memberLevelRepo) CreateMemberLevl(ctx context.Context, levl *ums.MemberLevel) error {
	panic("implement me")
}

func (m memberLevelRepo) GetMemberLevl(ctx context.Context, id int64) error {
	panic("implement me")
}

func (m memberLevelRepo) UpdateMemberLevl(ctx context.Context, levl *ums.MemberLevel) error {
	panic("implement me")
}

func (m memberLevelRepo) ListMemberLevl(ctx context.Context, req *ums.MemberLevlListReq) (*ums.MemberLevelListResp, error) {
	list, _ := m.data.UmsClient.MemberLevelList(ctx, &umsV1.MemberLevelListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*ums.MemberLevel, 0)
	copier.Copy(&orders, &list.List)

	return &ums.MemberLevelListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (m memberLevelRepo) DeleteMemberLevl(ctx context.Context, id int64) error {
	panic("implement me")
}
