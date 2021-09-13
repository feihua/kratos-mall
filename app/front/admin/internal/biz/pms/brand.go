package pms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type BrandListReq struct {
	Current  int64
	PageSize int64
}

type Brand struct {
	Id                  int64
	Name                string
	FirstLetter         string // 首字母
	Sort                int
	FactoryStatus       int // 是否为品牌制造商：0->不是；1->是
	ShowStatus          int
	ProductCount        int    // 产品数量
	ProductCommentCount int    // 产品评论数量
	Logo                string // 品牌logo
	BigPic              string // 专区大图
	BrandStory          string // 品牌故事
}

type BrandListResp struct {
	Total int64
	List  []*Brand
}

type BrandRepo interface {
	CreateBrand(context.Context, *Brand) error
	GetBrand(ctx context.Context, id int64) error
	UpdateBrand(context.Context, *Brand) error
	ListBrand(ctx context.Context, req *BrandListReq) (*BrandListResp, error)
	DeleteBrand(ctx context.Context, id int64) error
}

type BrandUseCase struct {
	repo BrandRepo
	log  *log.Helper
}

func NewBrandUseCase(repo BrandRepo, logger log.Logger) *BrandUseCase {
	return &BrandUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (b BrandUseCase) CreateBrand(ctx context.Context, brand *Brand) error {
	panic("implement me")
}

func (b BrandUseCase) GetBrand(ctx context.Context, id int64) error {
	panic("implement me")
}

func (b BrandUseCase) UpdateBrand(ctx context.Context, brand *Brand) error {
	panic("implement me")
}

func (b BrandUseCase) ListBrand(ctx context.Context, req *BrandListReq) (*BrandListResp, error) {
	return b.repo.ListBrand(ctx, req)
}

func (b BrandUseCase) DeleteBrand(ctx context.Context, id int64) error {
	panic("implement me")
}
