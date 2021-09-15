package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"strconv"
)

type RoleListReq struct {
	Current  int64
	PageSize int64
	Name     string
	Status   int64
}

type Role struct {
	Id             int64  // 编号
	Name           string // 角色名称
	Remark         string // 备注
	CreateBy       string // 创建人
	CreateTime     string // 创建时间
	LastUpdateBy   string // 更新人
	LastUpdateTime string // 更新时间
	DelFlag        int    // 是否删除  -1：已删除  0：正常
	Status         int    // 状态  1:启用,0:禁用
}
type RoleListResp struct {
	Total int64
	List  []*Role
}

type ListMenuDataResp struct {
	Key      string
	Title    string
	ParentId int32
	Label    string
	Id       int32
}

type QueryMenuByRoleIdResp struct {
	Ids  []int32
	List []*ListMenuDataResp
}

type RoleRepo interface {
	CreateRole(context.Context, *Role) error
	GetRole(ctx context.Context, id int64) error
	UpdateRole(context.Context, *Role) error
	ListRole(ctx context.Context, req *RoleListReq) (*RoleListResp, error)
	DeleteRole(ctx context.Context, id int64) error
	UpdateMenuRole(ctx context.Context, id int64, menuId int64) error
	DeleteMenuRole(ctx context.Context, id int64) error
	QueryMenuByRoleId(ctx context.Context, id int64) ([]int32, error)
}

type RoleUseCase struct {
	repo     RoleRepo
	menuRepo MenuRepo
	log      *log.Helper
}

func NewRoleUseCase(repo RoleRepo, menuRepo MenuRepo, logger log.Logger) *RoleUseCase {
	return &RoleUseCase{repo: repo, menuRepo: menuRepo, log: log.NewHelper(logger)}
}

func (r *RoleUseCase) CreateRole(ctx context.Context, user *Role) error {
	panic("implement me")
}

func (r *RoleUseCase) GetRole(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r *RoleUseCase) UpdateRole(ctx context.Context, user *Role) error {
	panic("implement me")
}

func (r *RoleUseCase) ListRole(ctx context.Context, req *RoleListReq) (*RoleListResp, error) {
	return r.repo.ListRole(ctx, req)
}

func (r *RoleUseCase) DeleteRole(ctx context.Context, id int64) error {
	panic("implement me")
}
func (r RoleUseCase) UpdateMenuRole(ctx context.Context, id int64, menuIds []int64) error {

	_ = r.repo.DeleteMenuRole(ctx, id)

	for _, menuId := range menuIds {
		_ = r.repo.UpdateMenuRole(ctx, id, menuId)
	}

	return nil

}

func (r RoleUseCase) QueryMenuByRoleId(ctx context.Context, id int64) (*QueryMenuByRoleIdResp, error) {

	roleId, _ := r.repo.QueryMenuByRoleId(ctx, id)

	listMenu, _ := r.menuRepo.ListMenu(ctx, &MenuListReq{})
	resps := make([]*ListMenuDataResp, 0)

	for _, menu := range listMenu {
		resps = append(resps, &ListMenuDataResp{
			Key:      strconv.FormatInt(menu.Id, 10),
			Title:    menu.Name,
			ParentId: int32(menu.ParentId),
			Label:    menu.Name,
			Id:       int32(menu.Id),
		})
	}
	return &QueryMenuByRoleIdResp{
		Ids:  roleId,
		List: resps,
	}, nil
}
