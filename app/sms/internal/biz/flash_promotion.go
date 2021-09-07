package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type FlashPromotionListReq struct {
	Current  int64
	PageSize int64
}

type FlashPromotion struct {
	Id         int64
	Title      string
	StartDate  time.Time // 开始日期
	EndDate    time.Time // 结束日期
	Status     int       // 上下线状态
	CreateTime time.Time // 秒杀时间段名称
}

type FlashPromotionRepo interface {
	CreateFlashPromotion(context.Context, *FlashPromotion) error
	GetFlashPromotion(ctx context.Context, id int64) error
	UpdateFlashPromotion(context.Context, *FlashPromotion) error
	ListFlashPromotion(ctx context.Context, req *FlashPromotionListReq) ([]*FlashPromotion, error)
	DeleteFlashPromotion(ctx context.Context, id int64) error
}

type FlashPromotionUseCase struct {
	repo FlashPromotionRepo
	log  *log.Helper
}

func NewFlashPromotionUseCase(repo FlashPromotionRepo, logger log.Logger) *FlashPromotionUseCase {
	return &FlashPromotionUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (r *FlashPromotionUseCase) CreateFlashPromotion(ctx context.Context, user *FlashPromotion) error {
	panic("implement me")
}

func (r *FlashPromotionUseCase) GetFlashPromotion(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r *FlashPromotionUseCase) UpdateFlashPromotion(ctx context.Context, user *FlashPromotion) error {
	panic("implement me")
}

func (r *FlashPromotionUseCase) ListFlashPromotion(ctx context.Context, pageNum, pageSize int64) ([]*FlashPromotion, error) {
	panic("implement me")
}

func (r *FlashPromotionUseCase) DeleteFlashPromotion(ctx context.Context, id int64) error {
	panic("implement me")
}
