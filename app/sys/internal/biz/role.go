package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
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

type RoleRepo interface {
	CreateRole(context.Context, *Role) error
	GetRole(ctx context.Context, id int64) error
	UpdateRole(context.Context, *Role) error
	ListRole(ctx context.Context, req *RoleListReq) ([]*Role, error)
	DeleteRole(ctx context.Context, id int64) error
}

type RoleUseCase struct {
	repo RoleRepo
	log  *log.Helper
}

func NewRoleUseCase(repo RoleRepo, logger log.Logger) *RoleUseCase {
	return &RoleUseCase{repo: repo, log: log.NewHelper(logger)}
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

func (r *RoleUseCase) ListRole(ctx context.Context, pageNum, pageSize int64) ([]*Role, error) {
	panic("implement me")
}

func (r *RoleUseCase) DeleteRole(ctx context.Context, id int64) error {
	panic("implement me")
}
