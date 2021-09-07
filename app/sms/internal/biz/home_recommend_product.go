package biz

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

type HomeRecommendProductRepo interface {
	CreateHomeRecommendProduct(context.Context, *HomeRecommendProduct) error
	GetHomeRecommendProduct(ctx context.Context, id int64) error
	UpdateHomeRecommendProduct(context.Context, *HomeRecommendProduct) error
	ListHomeRecommendProduct(ctx context.Context, req *HomeRecommendProductListReq) ([]*HomeRecommendProduct, error)
	DeleteHomeRecommendProduct(ctx context.Context, id int64) error
}

type HomeRecommendProductUseCase struct {
	repo HomeRecommendProductRepo
	log  *log.Helper
}

func NewHomeRecommendProductUseCase(repo HomeRecommendProductRepo, logger log.Logger) *HomeRecommendProductUseCase {
	return &HomeRecommendProductUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (r *HomeRecommendProductUseCase) CreateHomeRecommendProduct(ctx context.Context, user *HomeRecommendProduct) error {
	panic("implement me")
}

func (r *HomeRecommendProductUseCase) GetHomeRecommendProduct(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r *HomeRecommendProductUseCase) UpdateHomeRecommendProduct(ctx context.Context, user *HomeRecommendProduct) error {
	panic("implement me")
}

func (r *HomeRecommendProductUseCase) ListHomeRecommendProduct(ctx context.Context, pageNum, pageSize int64) ([]*HomeRecommendProduct, error) {
	panic("implement me")
}

func (r *HomeRecommendProductUseCase) DeleteHomeRecommendProduct(ctx context.Context, id int64) error {
	panic("implement me")
}
