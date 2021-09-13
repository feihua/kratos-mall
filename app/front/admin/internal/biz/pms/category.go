package pms

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
type CategoryListResp struct {
	Total int64
	List  []*Category
}

type CategoryRepo interface {
	CreateCategory(context.Context, *Category) error
	GetCategory(ctx context.Context, id int64) error
	UpdateCategory(context.Context, *Category) error
	ListCategory(ctx context.Context, req *CategoryListReq) (*CategoryListResp, error)
	DeleteCategory(ctx context.Context, id int64) error
}

type CategoryUseCase struct {
	repo CategoryRepo
	log  *log.Helper
}

func NewCategoryUseCase(repo CategoryRepo, logger log.Logger) *CategoryUseCase {
	return &CategoryUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (c CategoryUseCase) CreateCategory(ctx context.Context, category *Category) error {
	panic("implement me")
}

func (c CategoryUseCase) GetCategory(ctx context.Context, id int64) error {
	panic("implement me")
}

func (c CategoryUseCase) UpdateCategory(ctx context.Context, category *Category) error {
	panic("implement me")
}

func (c CategoryUseCase) ListCategory(ctx context.Context, req *CategoryListReq) (*CategoryListResp, error) {
	return c.repo.ListCategory(ctx, req)
}

func (c CategoryUseCase) DeleteCategory(ctx context.Context, id int64) error {
	panic("implement me")
}
