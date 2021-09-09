package sys

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/front/admin/internal/biz/sys"
	"kratos-mall/app/front/admin/internal/data"
)

type deptRepo struct {
	data *data.Data
	log  *log.Helper
}

// NewDeptRepo .
func NewDeptRepo(data *data.Data, logger log.Logger) sys.DeptRepo {
	return &deptRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (d deptRepo) CreateDept(ctx context.Context, dept *sys.Dept) error {
	panic("implement me")
}

func (d deptRepo) GetDept(ctx context.Context, id int64) error {
	panic("implement me")
}

func (d deptRepo) UpdateDept(ctx context.Context, dept *sys.Dept) error {
	panic("implement me")
}

func (d deptRepo) ListDept(ctx context.Context, req *sys.DeptListReq) ([]*sys.Dept, error) {
	panic("implement me")
}

func (d deptRepo) DeleteDept(ctx context.Context, id int64) error {
	panic("implement me")
}
