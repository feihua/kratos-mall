package sys

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	sysV1 "kratos-mall/api/sys/v1"
	"kratos-mall/app/front/admin/internal/biz/sys"
	"kratos-mall/app/front/admin/internal/data"
)

type menuRepo struct {
	data *data.Data
	log  *log.Helper
}

// NewMenuRepo .
func NewMenuRepo(data *data.Data, logger log.Logger) sys.MenuRepo {
	return &menuRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u menuRepo) CreateMenu(ctx context.Context, menu *sys.Menu) error {
	panic("implement me")
}

func (u menuRepo) GetMenu(ctx context.Context, id int64) error {
	panic("implement me")
}

func (u menuRepo) UpdateMenu(ctx context.Context, menu *sys.Menu) error {
	panic("implement me")
}

func (m menuRepo) ListMenu(ctx context.Context, req *sys.MenuListReq) (*sys.MenuListResp, error) {

	list, _ := m.data.SysClient.MenuList(ctx, &sysV1.MenuListReq{
		Name: req.Name,
		Url:  req.Url,
	})

	menus := make([]*sys.Menu, 0)
	copier.Copy(&menus, &list.List)

	return &sys.MenuListResp{
		Total: list.Total,
		List:  menus,
	}, nil
}

func (u menuRepo) DeleteMenu(ctx context.Context, id int64) error {
	panic("implement me")
}
