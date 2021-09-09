package sms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type HomeNewProductListReq struct {
	Current  int64
	PageSize int64
}

type HomeNewProduct struct {
	Id              int64
	ProductId       int64
	ProductName     string
	RecommendStatus int
	Sort            int
}

type HomeNewProductRepo interface {
	CreateHomeNewProduct(context.Context, *HomeNewProduct) error
	GetHomeNewProduct(ctx context.Context, id int64) error
	UpdateHomeNewProduct(context.Context, *HomeNewProduct) error
	ListHomeNewProduct(ctx context.Context, req *HomeNewProductListReq) ([]*HomeNewProduct, error)
	DeleteHomeNewProduct(ctx context.Context, id int64) error
}

type HomeNewProductUseCase struct {
	repo HomeNewProductRepo
	log  *log.Helper
}

func NewHomeNewProductUseCase(repo HomeNewProductRepo, logger log.Logger) *HomeNewProductUseCase {
	return &HomeNewProductUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (r *HomeNewProductUseCase) CreateHomeNewProduct(ctx context.Context, user *HomeNewProduct) error {
	panic("implement me")
}

func (r *HomeNewProductUseCase) GetHomeNewProduct(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r *HomeNewProductUseCase) UpdateHomeNewProduct(ctx context.Context, user *HomeNewProduct) error {
	panic("implement me")
}

func (r *HomeNewProductUseCase) ListHomeNewProduct(ctx context.Context, pageNum, pageSize int64) ([]*HomeNewProduct, error) {
	panic("implement me")
}

func (r *HomeNewProductUseCase) DeleteHomeNewProduct(ctx context.Context, id int64) error {
	panic("implement me")
}
