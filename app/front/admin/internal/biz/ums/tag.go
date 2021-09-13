package ums

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type TagListReq struct {
	Current  int64
	PageSize int64
}

type Tag struct {
	Id                int64
	Name              string
	FinishOrderCount  int    // 自动打标签完成订单数量
	FinishOrderAmount string // 自动打标签完成订单金额
}
type TagListResp struct {
	Total int64
	List  []*Tag
}

type TagRepo interface {
	CreateTag(context.Context, *Tag) error
	GetTag(ctx context.Context, id int64) error
	UpdateTag(context.Context, *Tag) error
	ListTag(ctx context.Context, req *TagListReq) (*TagListResp, error)
	DeleteTag(ctx context.Context, id int64) error
}

type TagUseCase struct {
	repo TagRepo
	log  *log.Helper
}

func NewTagUseCase(repo TagRepo, logger log.Logger) *TagUseCase {
	return &TagUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (t TagUseCase) CreateTag(ctx context.Context, tag *Tag) error {
	panic("implement me")
}

func (t TagUseCase) GetTag(ctx context.Context, id int64) error {
	panic("implement me")
}

func (t TagUseCase) UpdateTag(ctx context.Context, tag *Tag) error {
	panic("implement me")
}

func (t TagUseCase) ListTag(ctx context.Context, req *TagListReq) (*TagListResp, error) {
	return t.repo.ListTag(ctx, req)
}

func (t TagUseCase) DeleteTag(ctx context.Context, id int64) error {
	panic("implement me")
}
