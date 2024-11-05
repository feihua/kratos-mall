package data

import (
	"context"
	"github.com/feihua/kratos-mall/app/sys/internal/biz"
	"github.com/feihua/kratos-mall/app/sys/internal/data/model"
	"github.com/feihua/kratos-mall/pkg/util/pagination"
	"github.com/feihua/kratos-mall/pkg/util/time2str"
	"github.com/go-kratos/kratos/v2/log"
	"time"
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

func (r roleRepo) ListRole(ctx context.Context, req *biz.RoleListReq) (*biz.RoleListResp, error) {
	var all []model.SysRole
	result := r.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	r.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.Role, 0)

	for _, role := range all {
		list = append(list, &biz.Role{
			Id:             role.Id,
			Name:           role.Name,
			Remark:         role.Remark,
			CreateBy:       role.CreateBy,
			CreateTime:     time2str.String(role.CreateTime),
			LastUpdateBy:   role.LastUpdateBy,
			LastUpdateTime: time2str.String(role.LastUpdateTime),
			DelFlag:        role.DelFlag,
			Status:         role.Status,
		})
	}

	return &biz.RoleListResp{
		Total: count,
		List:  list,
	}, nil
}

func (r roleRepo) DeleteRole(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r roleRepo) QueryMenuByRoleId(ctx context.Context, id int64) ([]int32, error) {

	var all []model.SysRoleMenu
	r.data.db.WithContext(ctx).Where(&model.SysRoleMenu{RoleId: id}).Find(&all)

	var list []int32
	for _, user := range all {

		list = append(list, int32(user.MenuId))
	}

	return list, nil

}

func (r roleRepo) DeleteMenuRole(ctx context.Context, id int64) error {
	return r.data.db.Where(&model.SysRoleMenu{RoleId: id}).Delete(&model.SysRoleMenu{}).Error
}

func (r roleRepo) UpdateMenuRole(ctx context.Context, id int64, menuId int64) error {

	r.data.db.Create(&model.SysRoleMenu{
		RoleId:         id,
		MenuId:         menuId,
		CreateBy:       "admin",
		CreateTime:     time.Now(),
		LastUpdateBy:   "admin",
		LastUpdateTime: time.Now(),
	})

	return nil
}
