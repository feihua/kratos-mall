package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/ums/internal/biz"
)

type memberLevelRepo struct {
	data *Data
	log  *log.Helper
}

func NewMemberLevelRepo(data *Data, logger log.Logger) biz.MemberLevelRepo {
	return &memberLevelRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (m memberLevelRepo) CreateMemberLevl(ctx context.Context, levl *biz.MemberLevel) error {
	panic("implement me")
}

func (m memberLevelRepo) GetMemberLevl(ctx context.Context, id int64) error {
	panic("implement me")
}

func (m memberLevelRepo) UpdateMemberLevl(ctx context.Context, levl *biz.MemberLevel) error {
	panic("implement me")
}

func (m memberLevelRepo) ListMemberLevl(ctx context.Context, req *biz.MemberLevlListReq) ([]*biz.MemberLevel, error) {
	panic("implement me")
}

func (m memberLevelRepo) DeleteMemberLevl(ctx context.Context, id int64) error {
	panic("implement me")
}
