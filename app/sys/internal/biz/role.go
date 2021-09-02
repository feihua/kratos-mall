package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Role struct {
	Hello string
}

type RoleRepo interface {
	CreateRole(context.Context, *Role) error
	GetRole(ctx context.Context, id int64) error
	UpdateRole(context.Context, *Role) error
	ListRole(ctx context.Context, pageNum, pageSize int64) ([]*Beer, error)
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

func (r *RoleUseCase) ListRole(ctx context.Context, pageNum, pageSize int64) ([]*Beer, error) {
	panic("implement me")
}

func (r *RoleUseCase) DeleteRole(ctx context.Context, id int64) error {
	panic("implement me")
}
