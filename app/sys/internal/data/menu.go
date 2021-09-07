package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/sys/internal/biz"
	"kratos-mall/app/sys/internal/data/model"
)

type menuRepo struct {
	data *Data
	log  *log.Helper
}

// NewMenuRepo .
func NewMenuRepo(data *Data, logger log.Logger) biz.MenuRepo {
	return &menuRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u menuRepo) CreateMenu(ctx context.Context, menu *biz.Menu) error {
	panic("implement me")
}

func (u menuRepo) GetMenu(ctx context.Context, id int64) error {
	panic("implement me")
}

func (u menuRepo) UpdateMenu(ctx context.Context, menu *biz.Menu) error {
	panic("implement me")
}

func (m menuRepo) ListMenu(ctx context.Context, req *biz.MenuListReq) ([]*biz.Menu, error) {

	var memus []model.SysMenu
	m.data.db.WithContext(ctx).Find(&memus)

	rv := make([]*biz.Menu, 0)
	for _, m := range memus {
		rv = append(rv, &biz.Menu{
			Id:             m.Id,
			Name:           m.Name,
			ParentId:       m.ParentId,
			Url:            m.Url,
			Perms:          m.Perms,
			Type:           m.Type,
			Icon:           m.Icon,
			OrderNum:       m.OrderNum,
			CreateBy:       m.CreateBy,
			CreateTime:     m.CreateTime,
			LastUpdateBy:   m.LastUpdateBy,
			LastUpdateTime: m.LastUpdateTime,
			DelFlag:        m.DelFlag,
		})
	}
	return rv, nil
}

func (u menuRepo) DeleteMenu(ctx context.Context, id int64) error {
	panic("implement me")
}
