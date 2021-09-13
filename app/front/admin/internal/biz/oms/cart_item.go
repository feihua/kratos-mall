package oms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type CartItemListReq struct {
	Current  int64
	PageSize int64
}

type CartItem struct {
	Id                int64
	ProductId         int64
	ProductSkuId      int64
	MemberId          int64
	Quantity          int    // 购买数量
	Price             string // 添加到购物车的价格
	ProductPic        string // 商品主图
	ProductName       string // 商品名称
	ProductSubTitle   string // 商品副标题（卖点）
	ProductSkuCode    string // 商品sku条码
	MemberNickname    string // 会员昵称
	CreateDate        string // 创建时间
	ModifyDate        string // 修改时间
	DeleteStatus      int    // 是否删除
	ProductCategoryId int64  // 商品分类
	ProductBrand      string
	ProductSn         string
	ProductAttr       string // 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]

}

type CartItemListResp struct {
	Total int64
	List  []*CartItem
}

type CartItemRepo interface {
	CreateCartItem(context.Context, *CartItem) error
	GetCartItem(ctx context.Context, id int64) error
	UpdateCartItem(context.Context, *CartItem) error
	ListCartItem(ctx context.Context, req *CartItemListReq) (*CartItemListResp, error)
	DeleteCartItem(ctx context.Context, id int64) error
}

type CartItemUseCase struct {
	repo CartItemRepo
	log  *log.Helper
}

func NewCartItemUseCase(repo CartItemRepo, logger log.Logger) *CartItemUseCase {
	return &CartItemUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (c CartItemUseCase) CreateCartItem(ctx context.Context, item *CartItem) error {
	panic("implement me")
}

func (c CartItemUseCase) GetCartItem(ctx context.Context, id int64) error {
	panic("implement me")
}

func (c CartItemUseCase) UpdateCartItem(ctx context.Context, item *CartItem) error {
	panic("implement me")
}

func (c CartItemUseCase) ListCartItem(ctx context.Context, req *CartItemListReq) (*CartItemListResp, error) {
	return c.repo.ListCartItem(ctx, req)
}

func (c CartItemUseCase) DeleteCartItem(ctx context.Context, id int64) error {
	panic("implement me")
}
