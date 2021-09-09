package sms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type HomeBrandListReq struct {
	Current  int64
	PageSize int64
}

type HomeBrand struct {
	Id              int64
	BrandId         int64
	BrandName       string
	RecommendStatus int
	Sort            int
}

type HomeBrandRepo interface {
	CreateHomeBrand(context.Context, *HomeBrand) error
	GetHomeBrand(ctx context.Context, id int64) error
	UpdateHomeBrand(context.Context, *HomeBrand) error
	ListHomeBrand(ctx context.Context, req *HomeBrandListReq) ([]*HomeBrand, error)
	DeleteHomeBrand(ctx context.Context, id int64) error
}

type HomeBrandUseCase struct {
	repo HomeBrandRepo
	log  *log.Helper
}

func NewHomeBrandUseCase(repo HomeBrandRepo, logger log.Logger) *HomeBrandUseCase {
	return &HomeBrandUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (r *HomeBrandUseCase) CreateHomeBrand(ctx context.Context, user *HomeBrand) error {
	panic("implement me")
}

func (r *HomeBrandUseCase) GetHomeBrand(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r *HomeBrandUseCase) UpdateHomeBrand(ctx context.Context, user *HomeBrand) error {
	panic("implement me")
}

func (r *HomeBrandUseCase) ListHomeBrand(ctx context.Context, pageNum, pageSize int64) ([]*HomeBrand, error) {
	panic("implement me")
}

func (r *HomeBrandUseCase) DeleteHomeBrand(ctx context.Context, id int64) error {
	panic("implement me")
}
