package sys

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type MenuListReq struct {
	Name string
	Url  string
}

type Menu struct {
	Id             int64  // 编号
	Name           string // 菜单名称
	ParentId       int64  // 父菜单ID，一级菜单为0
	Url            string
	Perms          string // 授权(多个用逗号分隔，如：sys:user:add,sys:user:edit)
	Type           int    // 类型   0：目录   1：菜单   2：按钮
	Icon           string // 菜单图标
	OrderNum       int    // 排序
	CreateBy       string // 创建人
	CreateTime     string // 创建时间
	LastUpdateBy   string // 更新人
	LastUpdateTime string // 更新时间
	DelFlag        int    // 是否删除  -1：已删除  0：正常
}
type MenuListResp struct {
	Total int64
	List  []*Menu
}

type MenuRepo interface {
	CreateMenu(context.Context, *Menu) error
	GetMenu(ctx context.Context, id int64) error
	UpdateMenu(context.Context, *Menu) error
	ListMenu(ctx context.Context, req *MenuListReq) ([]*Menu, error)
	DeleteMenu(ctx context.Context, id int64) error
}

type MenuUseCase struct {
	repo MenuRepo
	log  *log.Helper
}

func NewMenuUseCase(repo MenuRepo, logger log.Logger) *MenuUseCase {
	return &MenuUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (m *MenuUseCase) CreateMenu(ctx context.Context, user *Menu) error {
	panic("implement me")
}

func (m *MenuUseCase) GetMenu(ctx context.Context, id int64) error {
	panic("implement me")
}

func (m *MenuUseCase) UpdateMenu(ctx context.Context, user *Menu) error {
	panic("implement me")
}

func (m *MenuUseCase) ListMenu(ctx context.Context, req *MenuListReq) ([]*Menu, error) {

	listMenu, _ := m.repo.ListMenu(ctx, req)

	return listMenu, nil
}

func (m *MenuUseCase) DeleteMenu(ctx context.Context, id int64) error {
	panic("implement me")
}
