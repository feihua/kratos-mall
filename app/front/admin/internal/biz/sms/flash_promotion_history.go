package sms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type FlashPromotionHistoryListReq struct {
	Current  int64
	PageSize int64
}

type FlashPromotionHistory struct {
	Id            int
	MemberId      int
	ProductId     int64
	MemberPhone   string
	ProductName   string
	SubscribeTime time.Time // 会员订阅时间
	SendTime      time.Time
}

type FlashPromotionHistoryRepo interface {
	CreateFlashPromotionHistory(context.Context, *FlashPromotionHistory) error
	GetFlashPromotionHistory(ctx context.Context, id int64) error
	UpdateFlashPromotionHistory(context.Context, *FlashPromotionHistory) error
	ListFlashPromotionHistory(ctx context.Context, req *FlashPromotionHistoryListReq) ([]*FlashPromotionHistory, error)
	DeleteFlashPromotionHistory(ctx context.Context, id int64) error
}

type FlashPromotionHistoryUseCase struct {
	repo FlashPromotionHistoryRepo
	log  *log.Helper
}

func NewFlashPromotionHistoryUseCase(repo FlashPromotionHistoryRepo, logger log.Logger) *FlashPromotionHistoryUseCase {
	return &FlashPromotionHistoryUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (r *FlashPromotionHistoryUseCase) CreateFlashPromotionHistory(ctx context.Context, user *FlashPromotionHistory) error {
	panic("implement me")
}

func (r *FlashPromotionHistoryUseCase) GetFlashPromotionHistory(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r *FlashPromotionHistoryUseCase) UpdateFlashPromotionHistory(ctx context.Context, user *FlashPromotionHistory) error {
	panic("implement me")
}

func (r *FlashPromotionHistoryUseCase) ListFlashPromotionHistory(ctx context.Context, pageNum, pageSize int64) ([]*FlashPromotionHistory, error) {
	panic("implement me")
}

func (r *FlashPromotionHistoryUseCase) DeleteFlashPromotionHistory(ctx context.Context, id int64) error {
	panic("implement me")
}
