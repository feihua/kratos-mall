package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type ProductListReq struct {
	Current  int64
	PageSize int64
}

type Product struct {
	Id                         int64
	BrandId                    int64
	ProductCategoryId          int64
	FeightTemplateId           int64
	ProductAttributeCategoryId int64
	Name                       string
	Pic                        string
	ProductSn                  string // 货号
	DeleteStatus               int    // 删除状态：0->未删除；1->已删除
	PublishStatus              int    // 上架状态：0->下架；1->上架
	NewStatus                  int    // 新品状态:0->不是新品；1->新品
	RecommandStatus            int    // 推荐状态；0->不推荐；1->推荐
	VerifyStatus               int    // 审核状态：0->未审核；1->审核通过
	Sort                       int    // 排序
	Sale                       int    // 销量
	Price                      string
	PromotionPrice             string // 促销价格
	GiftGrowth                 int    // 赠送的成长值
	GiftPoint                  int    // 赠送的积分
	UsePointLimit              int    // 限制使用的积分数
	SubTitle                   string // 副标题
	Description                string // 商品描述
	OriginalPrice              string // 市场价
	Stock                      int    // 库存
	LowStock                   int    // 库存预警值
	Unit                       string // 单位
	Weight                     string // 商品重量，默认为克
	PreviewStatus              int    // 是否为预告商品：0->不是；1->是
	ServiceIds                 string // 以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮
	Keywords                   string
	Note                       string
	AlbumPics                  string // 画册图片，连产品图片限制为5张，以逗号分割
	DetailTitle                string
	DetailDesc                 string
	DetailHtml                 string // 产品详情网页内容
	DetailMobileHtml           string // 移动端网页详情
	PromotionStartTime         string // 促销开始时间
	PromotionEndTime           string // 促销结束时间
	PromotionPerLimit          int    // 活动限购数量
	PromotionType              int    // 促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购
	BrandName                  string // 品牌名称
	ProductCategoryName        string // 商品分类名称
}
type ProductListResp struct {
	Total int64
	List  []*Product
}

type ProductRepo interface {
	CreateProduct(context.Context, *Product) error
	GetProduct(ctx context.Context, id int64) error
	UpdateProduct(context.Context, *Product) error
	ListProduct(ctx context.Context, req *ProductListReq) (*ProductListResp, error)
	DeleteProduct(ctx context.Context, id int64) error
}

type ProductUseCase struct {
	repo ProductRepo
	log  *log.Helper
}

func NewProductUseCase(repo ProductRepo, logger log.Logger) *ProductUseCase {
	return &ProductUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (p ProductUseCase) CreateProduct(ctx context.Context, product *Product) error {
	panic("implement me")
}

func (p ProductUseCase) GetProduct(ctx context.Context, id int64) error {
	panic("implement me")
}

func (p ProductUseCase) UpdateProduct(ctx context.Context, product *Product) error {
	panic("implement me")
}

func (p ProductUseCase) ListProduct(ctx context.Context, req *ProductListReq) (*ProductListResp, error) {
	return p.repo.ListProduct(ctx, req)
}

func (p ProductUseCase) DeleteProduct(ctx context.Context, id int64) error {
	panic("implement me")
}
