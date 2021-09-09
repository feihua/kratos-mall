package sys

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/front/admin/internal/biz/sys"
	"kratos-mall/app/front/admin/internal/data"
)

type roleRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewRoleRepo(data *data.Data, logger log.Logger) sys.RoleRepo {
	return &roleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r roleRepo) CreateRole(ctx context.Context, role *sys.Role) error {
	panic("implement me")
}

func (r roleRepo) GetRole(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r roleRepo) UpdateRole(ctx context.Context, role *sys.Role) error {
	panic("implement me")
}

func (r roleRepo) ListRole(ctx context.Context, req *sys.RoleListReq) ([]*sys.Role, error) {
	panic("implement me")
}

func (r roleRepo) DeleteRole(ctx context.Context, id int64) error {
	panic("implement me")
}
