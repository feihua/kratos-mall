package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type FlashPromotionListReq struct {
	Current  int64
	PageSize int64
}

type FlashPromotion struct {
	Id         int64
	Title      string
	StartDate  string // 开始日期
	EndDate    string // 结束日期
	Status     int    // 上下线状态
	CreateTime string // 秒杀时间段名称
}
type FlashPromotionListResp struct {
	Total int64
	List  []*FlashPromotion
}

type FlashPromotionRepo interface {
	CreateFlashPromotion(context.Context, *FlashPromotion) error
	GetFlashPromotion(ctx context.Context, id int64) error
	UpdateFlashPromotion(context.Context, *FlashPromotion) error
	ListFlashPromotion(ctx context.Context, req *FlashPromotionListReq) (*FlashPromotionListResp, error)
	DeleteFlashPromotion(ctx context.Context, id int64) error
}

type FlashPromotionUseCase struct {
	repo FlashPromotionRepo
	log  *log.Helper
}

func NewFlashPromotionUseCase(repo FlashPromotionRepo, logger log.Logger) *FlashPromotionUseCase {
	return &FlashPromotionUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (f FlashPromotionUseCase) CreateFlashPromotion(ctx context.Context, promotion *FlashPromotion) error {
	panic("implement me")
}

func (f FlashPromotionUseCase) GetFlashPromotion(ctx context.Context, id int64) error {
	panic("implement me")
}

func (f FlashPromotionUseCase) UpdateFlashPromotion(ctx context.Context, promotion *FlashPromotion) error {
	panic("implement me")
}

func (f FlashPromotionUseCase) ListFlashPromotion(ctx context.Context, req *FlashPromotionListReq) (*FlashPromotionListResp, error) {
	return f.repo.ListFlashPromotion(ctx, req)
}

func (f FlashPromotionUseCase) DeleteFlashPromotion(ctx context.Context, id int64) error {
	panic("implement me")
}
