package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Dept struct {
	Hello string
}

type DeptRepo interface {
	CreateDept(context.Context, *Dept) error
	GetDept(ctx context.Context, id int64) error
	UpdateDept(context.Context, *Dept) error
	ListDept(ctx context.Context, pageNum, pageSize int64) ([]*Beer, error)
	DeleteDept(ctx context.Context, id int64) error
}

type DeptUseCase struct {
	repo DeptRepo
	log  *log.Helper
}

func NewDeptUseCase(repo DeptRepo, logger log.Logger) *DeptUseCase {
	return &DeptUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (dc *DeptUseCase) CreateDept(ctx context.Context, user *Dept) error {
	panic("implement me")
}

func (dc *DeptUseCase) GetDept(ctx context.Context, id int64) error {
	panic("implement me")
}

func (dc *DeptUseCase) UpdateDept(ctx context.Context, user *Dept) error {
	panic("implement me")
}

func (dc *DeptUseCase) ListDept(ctx context.Context, pageNum, pageSize int64) ([]*Beer, error) {
	panic("implement me")
}

func (dc *DeptUseCase) DeleteDept(ctx context.Context, id int64) error {
	panic("implement me")
}
