package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type DeptListReq struct {
	Name     string
	CreateBy string
}
type Dept struct {
	Id             int64  // 编号
	Name           string // 机构名称
	ParentId       int64  // 上级机构ID，一级机构为0
	OrderNum       int    // 排序
	CreateBy       string // 创建人
	CreateTime     string // 创建时间
	LastUpdateBy   string // 更新人
	LastUpdateTime string // 更新时间
	DelFlag        int    // 是否删除  -1：已删除  0：正常
}

type DeptRepo interface {
	CreateDept(context.Context, *Dept) error
	GetDept(ctx context.Context, id int64) error
	UpdateDept(context.Context, *Dept) error
	ListDept(ctx context.Context, req *DeptListReq) ([]*Dept, error)
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

func (dc *DeptUseCase) ListDept(ctx context.Context, pageNum, pageSize int64) ([]*Dept, error) {
	panic("implement me")
}

func (dc *DeptUseCase) DeleteDept(ctx context.Context, id int64) error {
	panic("implement me")
}
