package pms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type SkuStockListReq struct {
	Current  int64
	PageSize int64
}

type SkuStock struct {
	Id             int64
	ProductId      int64
	SkuCode        string // sku编码
	Price          string
	Stock          int    // 库存
	LowStock       int    // 预警库存
	Pic            string // 展示图片
	Sale           int    // 销量
	PromotionPrice string // 单品促销价格
	LockStock      int    // 锁定库存
	SpData         string // 商品销售属性，json格式
}
type SkuStockListResp struct {
	Total int64
	List  []*SkuStock
}

type SkuStockRepo interface {
	CreateSkuStock(context.Context, *SkuStock) error
	GetSkuStock(ctx context.Context, id int64) error
	UpdateSkuStock(context.Context, *SkuStock) error
	ListSkuStock(ctx context.Context, req *SkuStockListReq) (*SkuStockListResp, error)
	DeleteSkuStock(ctx context.Context, id int64) error
}

type SkuStockUseCase struct {
	repo SkuStockRepo
	log  *log.Helper
}

func NewSkuStockUseCase(repo SkuStockRepo, logger log.Logger) *SkuStockUseCase {
	return &SkuStockUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (s SkuStockUseCase) CreateSkuStock(ctx context.Context, stock *SkuStock) error {
	panic("implement me")
}

func (s SkuStockUseCase) GetSkuStock(ctx context.Context, id int64) error {
	panic("implement me")
}

func (s SkuStockUseCase) UpdateSkuStock(ctx context.Context, stock *SkuStock) error {
	panic("implement me")
}

func (s SkuStockUseCase) ListSkuStock(ctx context.Context, req *SkuStockListReq) (*SkuStockListResp, error) {
	return s.repo.ListSkuStock(ctx, req)
}

func (s SkuStockUseCase) DeleteSkuStock(ctx context.Context, id int64) error {
	panic("implement me")
}
