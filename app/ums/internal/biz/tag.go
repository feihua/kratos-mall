package biz

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

type TagRepo interface {
	CreateTag(context.Context, *Tag) error
	GetTag(ctx context.Context, id int64) error
	UpdateTag(context.Context, *Tag) error
	ListTag(ctx context.Context, req *TagListReq) ([]*Tag, error)
	DeleteTag(ctx context.Context, id int64) error
}

type TagUseCase struct {
	repo TagRepo
	log  *log.Helper
}

func NewTagUseCase(repo TagRepo, logger log.Logger) *TagUseCase {
	return &TagUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (r *TagUseCase) CreateTag(ctx context.Context, user *Tag) error {
	panic("implement me")
}

func (r *TagUseCase) GetTag(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r *TagUseCase) UpdateTag(ctx context.Context, user *Tag) error {
	panic("implement me")
}

func (r *TagUseCase) ListTag(ctx context.Context, pageNum, pageSize int64) ([]*Tag, error) {
	panic("implement me")
}

func (r *TagUseCase) DeleteTag(ctx context.Context, id int64) error {
	panic("implement me")
}
