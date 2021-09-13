package sms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
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
	SubscribeTime string // 会员订阅时间
	SendTime      string
}
type FlashPromotionHistoryListResp struct {
	Total int64
	List  []*FlashPromotionHistory
}

type FlashPromotionHistoryRepo interface {
	CreateFlashPromotionHistory(context.Context, *FlashPromotionHistory) error
	GetFlashPromotionHistory(ctx context.Context, id int64) error
	UpdateFlashPromotionHistory(context.Context, *FlashPromotionHistory) error
	ListFlashPromotionHistory(ctx context.Context, req *FlashPromotionHistoryListReq) (*FlashPromotionHistoryListResp, error)
	DeleteFlashPromotionHistory(ctx context.Context, id int64) error
}

type FlashPromotionHistoryUseCase struct {
	repo FlashPromotionHistoryRepo
	log  *log.Helper
}

func NewFlashPromotionHistoryUseCase(repo FlashPromotionHistoryRepo, logger log.Logger) *FlashPromotionHistoryUseCase {
	return &FlashPromotionHistoryUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (f FlashPromotionHistoryUseCase) CreateFlashPromotionHistory(ctx context.Context, history *FlashPromotionHistory) error {
	panic("implement me")
}

func (f FlashPromotionHistoryUseCase) GetFlashPromotionHistory(ctx context.Context, id int64) error {
	panic("implement me")
}

func (f FlashPromotionHistoryUseCase) UpdateFlashPromotionHistory(ctx context.Context, history *FlashPromotionHistory) error {
	panic("implement me")
}

func (f FlashPromotionHistoryUseCase) ListFlashPromotionHistory(ctx context.Context, req *FlashPromotionHistoryListReq) (*FlashPromotionHistoryListResp, error) {
	return f.repo.ListFlashPromotionHistory(ctx, req)
}

func (f FlashPromotionHistoryUseCase) DeleteFlashPromotionHistory(ctx context.Context, id int64) error {
	panic("implement me")
}
