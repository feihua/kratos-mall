package sys

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
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

func (m menuRepo) ListMenu(ctx context.Context, req *sys.MenuListReq) ([]*sys.Menu, error) {

	//var memus []model.SysMenu
	//m.data.db.WithContext(ctx).Find(&memus)
	//
	//rv := make([]*sys.Menu, 0)
	//for _, m := range memus {
	//	rv = append(rv, &sys.Menu{
	//		Id:             m.Id,
	//		Name:           m.Name,
	//		ParentId:       m.ParentId,
	//		Url:            m.Url,
	//		Perms:          m.Perms,
	//		Type:           m.Type,
	//		Icon:           m.Icon,
	//		OrderNum:       m.OrderNum,
	//		CreateBy:       m.CreateBy,
	//		CreateTime:     m.CreateTime,
	//		LastUpdateBy:   m.LastUpdateBy,
	//		LastUpdateTime: m.LastUpdateTime,
	//		DelFlag:        m.DelFlag,
	//	})
	//}
	//return rv, nil
	return nil, nil
}

func (u menuRepo) DeleteMenu(ctx context.Context, id int64) error {
	panic("implement me")
}
