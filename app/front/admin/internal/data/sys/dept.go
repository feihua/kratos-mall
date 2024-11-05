package sys

import (
	"context"
	sysV1 "github.com/feihua/kratos-mall/api/sys/v1"
	"github.com/feihua/kratos-mall/app/front/admin/internal/biz/sys"
	"github.com/feihua/kratos-mall/app/front/admin/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
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

func (d deptRepo) ListDept(ctx context.Context, req *sys.DeptListReq) (*sys.DeptListResp, error) {
	list, _ := d.data.SysClient.DeptList(ctx, &sysV1.DeptListReq{
		Name:     "",
		CreateBy: "",
	})

	orders := make([]*sys.Dept, 0)
	copier.Copy(&orders, &list.List)

	return &sys.DeptListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (d deptRepo) DeleteDept(ctx context.Context, id int64) error {
	panic("implement me")
}
