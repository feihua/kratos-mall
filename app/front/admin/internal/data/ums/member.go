package ums

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/front/admin/internal/biz/ums"
	"kratos-mall/app/front/admin/internal/data"
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

func (m memberRepo) ListMember(ctx context.Context, req *ums.MemberListReq) ([]*ums.Member, error) {
	panic("implement me")
}

func (m memberRepo) DeleteMember(ctx context.Context, id int64) error {
	panic("implement me")
}
