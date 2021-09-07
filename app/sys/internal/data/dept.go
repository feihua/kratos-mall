package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/sys/internal/biz"
)

type deptRepo struct {
	data *Data
	log  *log.Helper
}

// NewDeptRepo .
func NewDeptRepo(data *Data, logger log.Logger) biz.DeptRepo {
	return &deptRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (d deptRepo) CreateDept(ctx context.Context, dept *biz.Dept) error {
	panic("implement me")
}

func (d deptRepo) GetDept(ctx context.Context, id int64) error {
	panic("implement me")
}

func (d deptRepo) UpdateDept(ctx context.Context, dept *biz.Dept) error {
	panic("implement me")
}

func (d deptRepo) ListDept(ctx context.Context, req *biz.DeptListReq) ([]*biz.Dept, error) {
	panic("implement me")
}

func (d deptRepo) DeleteDept(ctx context.Context, id int64) error {
	panic("implement me")
}
