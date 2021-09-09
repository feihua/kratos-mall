package ums

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
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

func (m memberLevelRepo) ListMemberLevl(ctx context.Context, req *ums.MemberLevlListReq) ([]*ums.MemberLevel, error) {
	panic("implement me")
}

func (m memberLevelRepo) DeleteMemberLevl(ctx context.Context, id int64) error {
	panic("implement me")
}
