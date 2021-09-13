package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/sys/internal/biz"
	"kratos-mall/app/sys/internal/data/model"
	"kratos-mall/pkg/util/pagination"
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
			CreateTime:     role.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   role.LastUpdateBy,
			LastUpdateTime: role.LastUpdateTime.Format("2006-01-02 15:04:05"),
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
