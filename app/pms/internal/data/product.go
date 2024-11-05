package data

import (
	"context"
	"github.com/feihua/kratos-mall/app/pms/internal/biz"
	"github.com/feihua/kratos-mall/app/pms/internal/data/model"
	"github.com/feihua/kratos-mall/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
)

type productRepo struct {
	data *Data
	log  *log.Helper
}

func NewProductRepo(data *Data, logger log.Logger) biz.ProductRepo {
	return &productRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (p productRepo) CreateProduct(ctx context.Context, product *biz.Product) error {
	panic("implement me")
}

func (p productRepo) GetProduct(ctx context.Context, id int64) error {
	panic("implement me")
}

func (p productRepo) UpdateProduct(ctx context.Context, product *biz.Product) error {
	panic("implement me")
}

func (p productRepo) ListProduct(ctx context.Context, req *biz.ProductListReq) (*biz.ProductListResp, error) {
	var all []model.PmsProduct
	result := p.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	p.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.Product, 0)

	for _, product := range all {
		list = append(list, &biz.Product{
			Id:                         product.Id,
			BrandId:                    product.BrandId,
			ProductCategoryId:          product.ProductCategoryId,
			FeightTemplateId:           product.FeightTemplateId,
			ProductAttributeCategoryId: product.ProductAttributeCategoryId,
			Name:                       product.Name,
			Pic:                        product.Pic,
			ProductSn:                  product.ProductSn,
			DeleteStatus:               product.DeleteStatus,
			PublishStatus:              product.PublishStatus,
			NewStatus:                  product.NewStatus,
			RecommandStatus:            product.RecommandStatus,
			VerifyStatus:               product.VerifyStatus,
			Sort:                       product.Sort,
			Sale:                       product.Sale,
			Price:                      product.Price,
			PromotionPrice:             product.PromotionPrice,
			GiftGrowth:                 product.GiftGrowth,
			GiftPoint:                  product.GiftPoint,
			UsePointLimit:              product.UsePointLimit,
			SubTitle:                   product.SubTitle,
			Description:                product.Description,
			OriginalPrice:              product.OriginalPrice,
			Stock:                      product.Stock,
			LowStock:                   product.LowStock,
			Unit:                       product.Unit,
			Weight:                     product.Weight,
			PreviewStatus:              product.PreviewStatus,
			ServiceIds:                 product.ServiceIds,
			Keywords:                   product.Keywords,
			Note:                       product.Note,
			AlbumPics:                  product.AlbumPics,
			DetailTitle:                product.DetailTitle,
			DetailDesc:                 product.DetailDesc,
			DetailHtml:                 product.DetailHtml,
			DetailMobileHtml:           product.DetailMobileHtml,
			PromotionStartTime:         product.PromotionStartTime.Format("2006-01-02 15:04:05"),
			PromotionEndTime:           product.PromotionEndTime.Format("2006-01-02 15:04:05"),
			PromotionPerLimit:          product.PromotionPerLimit,
			PromotionType:              product.PromotionType,
			BrandName:                  product.BrandName,
			ProductCategoryName:        product.ProductCategoryName,
		})
	}

	return &biz.ProductListResp{
		Total: count,
		List:  list,
	}, nil
}

func (p productRepo) DeleteProduct(ctx context.Context, id int64) error {
	panic("implement me")
}
