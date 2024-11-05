package data

import (
	"context"
	"github.com/feihua/kratos-mall/app/oms/internal/biz"
	"github.com/feihua/kratos-mall/app/oms/internal/data/model"
	"github.com/feihua/kratos-mall/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
)

type cartItemRepo struct {
	data *Data
	log  *log.Helper
}

// NewCartItemRepo .
func NewCartItemRepo(data *Data, logger log.Logger) biz.CartItemRepo {
	return &cartItemRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c cartItemRepo) CreateCartItem(ctx context.Context, item *biz.CartItem) error {
	panic("implement me")
}

func (c cartItemRepo) GetCartItem(ctx context.Context, id int64) error {
	panic("implement me")
}

func (c cartItemRepo) UpdateCartItem(ctx context.Context, item *biz.CartItem) error {
	panic("implement me")
}

func (c cartItemRepo) ListCartItem(ctx context.Context, req *biz.CartItemListReq) (*biz.CartItemListResp, error) {
	var all []model.OmsCartItem
	result := c.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	c.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.CartItem, 0)

	for _, item := range all {
		list = append(list, &biz.CartItem{
			Id:                item.Id,
			ProductId:         item.ProductId,
			ProductSkuId:      item.ProductSkuId,
			MemberId:          item.MemberId,
			Quantity:          item.Quantity,
			Price:             item.Price,
			ProductPic:        item.ProductPic,
			ProductName:       item.ProductName,
			ProductSubTitle:   item.ProductSubTitle,
			ProductSkuCode:    item.ProductSkuCode,
			MemberNickname:    item.MemberNickname,
			CreateDate:        item.CreateDate.Format("2006-01-02 15:04:05"),
			ModifyDate:        item.ModifyDate.Format("2006-01-02 15:04:05"),
			DeleteStatus:      item.DeleteStatus,
			ProductCategoryId: item.ProductCategoryId,
			ProductBrand:      item.ProductBrand,
			ProductSn:         item.ProductSn,
			ProductAttr:       item.ProductAttr,
		})
	}

	return &biz.CartItemListResp{
		Total: count,
		List:  list,
	}, nil
}

func (c cartItemRepo) DeleteCartItem(ctx context.Context, id int64) error {
	panic("implement me")
}
