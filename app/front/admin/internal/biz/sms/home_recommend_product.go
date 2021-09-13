package sms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type HomeRecommendProductListReq struct {
	Current  int64
	PageSize int64
}

type HomeRecommendProduct struct {
	Id              int64
	ProductId       int64
	ProductName     string
	RecommendStatus int
	Sort            int
}
type HomeRecommendProductListResp struct {
	Total int64
	List  []*HomeRecommendProduct
}

type HomeRecommendProductRepo interface {
	CreateHomeRecommendProduct(context.Context, *HomeRecommendProduct) error
	GetHomeRecommendProduct(ctx context.Context, id int64) error
	UpdateHomeRecommendProduct(context.Context, *HomeRecommendProduct) error
	ListHomeRecommendProduct(ctx context.Context, req *HomeRecommendProductListReq) (*HomeRecommendProductListResp, error)
	DeleteHomeRecommendProduct(ctx context.Context, id int64) error
}

type HomeRecommendProductUseCase struct {
	repo HomeRecommendProductRepo
	log  *log.Helper
}

func NewHomeRecommendProductUseCase(repo HomeRecommendProductRepo, logger log.Logger) *HomeRecommendProductUseCase {
	return &HomeRecommendProductUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (h HomeRecommendProductUseCase) CreateHomeRecommendProduct(ctx context.Context, product *HomeRecommendProduct) error {
	panic("implement me")
}

func (h HomeRecommendProductUseCase) GetHomeRecommendProduct(ctx context.Context, id int64) error {
	panic("implement me")
}

func (h HomeRecommendProductUseCase) UpdateHomeRecommendProduct(ctx context.Context, product *HomeRecommendProduct) error {
	panic("implement me")
}

func (h HomeRecommendProductUseCase) ListHomeRecommendProduct(ctx context.Context, req *HomeRecommendProductListReq) (*HomeRecommendProductListResp, error) {
	return h.repo.ListHomeRecommendProduct(ctx, req)
}

func (h HomeRecommendProductUseCase) DeleteHomeRecommendProduct(ctx context.Context, id int64) error {
	panic("implement me")
}
