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
type HomeBrandListResp struct {
	Total int64
	List  []*HomeBrand
}

type HomeBrandRepo interface {
	CreateHomeBrand(context.Context, *HomeBrand) error
	GetHomeBrand(ctx context.Context, id int64) error
	UpdateHomeBrand(context.Context, *HomeBrand) error
	ListHomeBrand(ctx context.Context, req *HomeBrandListReq) (*HomeBrandListResp, error)
	DeleteHomeBrand(ctx context.Context, id int64) error
}

type HomeBrandUseCase struct {
	repo HomeBrandRepo
	log  *log.Helper
}

func NewHomeBrandUseCase(repo HomeBrandRepo, logger log.Logger) *HomeBrandUseCase {
	return &HomeBrandUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (h HomeBrandUseCase) CreateHomeBrand(ctx context.Context, brand *HomeBrand) error {
	panic("implement me")
}

func (h HomeBrandUseCase) GetHomeBrand(ctx context.Context, id int64) error {
	panic("implement me")
}

func (h HomeBrandUseCase) UpdateHomeBrand(ctx context.Context, brand *HomeBrand) error {
	panic("implement me")
}

func (h HomeBrandUseCase) ListHomeBrand(ctx context.Context, req *HomeBrandListReq) (*HomeBrandListResp, error) {
	return h.repo.ListHomeBrand(ctx, req)
}

func (h HomeBrandUseCase) DeleteHomeBrand(ctx context.Context, id int64) error {
	panic("implement me")
}
