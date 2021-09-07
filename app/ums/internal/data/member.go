package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/ums/internal/biz"
)

type memberRepo struct {
	data *Data
	log  *log.Helper
}

func NewMemberRepo(data *Data, logger log.Logger) biz.MemberRepo {
	return &memberRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (m memberRepo) CreateMember(ctx context.Context, member *biz.Member) error {
	panic("implement me")
}

func (m memberRepo) GetMember(ctx context.Context, id int64) error {
	panic("implement me")
}

func (m memberRepo) UpdateMember(ctx context.Context, member *biz.Member) error {
	panic("implement me")
}

func (m memberRepo) ListMember(ctx context.Context, req *biz.MemberListReq) ([]*biz.Member, error) {
	panic("implement me")
}

func (m memberRepo) DeleteMember(ctx context.Context, id int64) error {
	panic("implement me")
}
