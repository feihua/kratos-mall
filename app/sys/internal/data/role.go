package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/sys/internal/biz"
)

type roleRepo struct {
	data *Data
	log  *log.Helper
}

func NewRoleRepo(data *Data, logger log.Logger) biz.RoleRepo {
	return &roleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r roleRepo) CreateRole(ctx context.Context, role *biz.Role) error {
	panic("implement me")
}

func (r roleRepo) GetRole(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r roleRepo) UpdateRole(ctx context.Context, role *biz.Role) error {
	panic("implement me")
}

func (r roleRepo) ListRole(ctx context.Context, req *biz.RoleListReq) ([]*biz.Role, error) {
	panic("implement me")
}

func (r roleRepo) DeleteRole(ctx context.Context, id int64) error {
	panic("implement me")
}
