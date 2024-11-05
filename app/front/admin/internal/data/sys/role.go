package sys

import (
	"context"
	sysV1 "github.com/feihua/kratos-mall/api/sys/v1"
	"github.com/feihua/kratos-mall/app/front/admin/internal/biz/sys"
	"github.com/feihua/kratos-mall/app/front/admin/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
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

func (r roleRepo) ListRole(ctx context.Context, req *sys.RoleListReq) (*sys.RoleListResp, error) {
	list, _ := r.data.SysClient.RoleList(ctx, &sysV1.RoleListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*sys.Role, 0)
	copier.Copy(&orders, &list.List)

	return &sys.RoleListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (r roleRepo) DeleteRole(ctx context.Context, id int64) error {
	panic("implement me")
}
func (r roleRepo) UpdateMenuRole(ctx context.Context, id int64, menuIds []int64) error {
	r.data.SysClient.UpdateMenuRole(ctx, &sysV1.UpdateMenuRoleReq{
		RoleId:  id,
		MenuIds: menuIds,
	})

	return nil
}

func (r roleRepo) QueryMenuByRoleId(ctx context.Context, id int64) (*sys.QueryMenuByRoleIdResp, error) {
	menuByRoleId, _ := r.data.SysClient.QueryMenuByRoleId(ctx, &sysV1.QueryMenuByRoleIdReq{Id: id})

	dataResps := make([]*sys.ListMenuDataResp, 0)
	_ = copier.Copy(&dataResps, &menuByRoleId.AllData)

	return &sys.QueryMenuByRoleIdResp{
		Ids:  menuByRoleId.UserData,
		List: dataResps,
	}, nil
}
