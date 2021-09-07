package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type CategoryListReq struct {
	Current  int64
	PageSize int64
}

type Category struct {
	Id           int64
	ParentId     int64 // 上机分类的编号：0表示一级分类
	Name         string
	Level        int // 分类级别：0->1级；1->2级
	ProductCount int
	ProductUnit  string
	NavStatus    int // 是否显示在导航栏：0->不显示；1->显示
	ShowStatus   int // 显示状态：0->不显示；1->显示
	Sort         int
	Icon         string // 图标
	Keywords     string
	Description  string // 描述
}

type CategoryRepo interface {
	CreateCategory(context.Context, *Category) error
	GetCategory(ctx context.Context, id int64) error
	UpdateCategory(context.Context, *Category) error
	ListCategory(ctx context.Context, req *CategoryListReq) ([]*Category, error)
	DeleteCategory(ctx context.Context, id int64) error
}

type CategoryUseCase struct {
	repo CategoryRepo
	log  *log.Helper
}

func NewCategoryUseCase(repo CategoryRepo, logger log.Logger) *CategoryUseCase {
	return &CategoryUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (r *CategoryUseCase) CreateCategory(ctx context.Context, user *Category) error {
	panic("implement me")
}

func (r *CategoryUseCase) GetCategory(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r *CategoryUseCase) UpdateCategory(ctx context.Context, user *Category) error {
	panic("implement me")
}

func (r *CategoryUseCase) ListCategory(ctx context.Context, pageNum, pageSize int64) ([]*Category, error) {
	panic("implement me")
}

func (r *CategoryUseCase) DeleteCategory(ctx context.Context, id int64) error {
	panic("implement me")
}
