package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Menu struct {
	Id             int64
	Name           string
	ParentId       int64
	Url            string
	Perms          string
	Type           int
	Icon           string
	OrderNum       int
	CreateBy       string
	CreateTime     time.Time
	LastUpdateBy   string
	LastUpdateTime time.Time
	DelFlag        int
}

type MenuRepo interface {
	CreateMenu(context.Context, *Menu) error
	GetMenu(ctx context.Context, id int64) error
	UpdateMenu(context.Context, *Menu) error
	ListMenu(ctx context.Context, pageNum, pageSize int64) ([]*Menu, error)
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

func (m *MenuUseCase) ListMenu(ctx context.Context, pageNum, pageSize int64) ([]*Menu, error) {

	listMenu, _ := m.repo.ListMenu(ctx, pageNum, pageSize)

	return listMenu, nil
}

func (m *MenuUseCase) DeleteMenu(ctx context.Context, id int64) error {
	panic("implement me")
}
