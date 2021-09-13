package sms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type FlashPromotionSessionListReq struct {
	Current  int64
	PageSize int64
}

type FlashPromotionSession struct {
	Id         int64  // 编号
	Name       string // 场次名称
	StartTime  string // 每日开始时间
	EndTime    string // 每日结束时间
	Status     int    // 启用状态：0->不启用；1->启用
	CreateTime string // 创建时间
}
type FlashPromotionSessionListResp struct {
	Total int64
	List  []*FlashPromotionSession
}

type FlashPromotionSessionRepo interface {
	CreateFlashPromotionSession(context.Context, *FlashPromotionSession) error
	GetFlashPromotionSession(ctx context.Context, id int64) error
	UpdateFlashPromotionSession(context.Context, *FlashPromotionSession) error
	ListFlashPromotionSession(ctx context.Context, req *FlashPromotionSessionListReq) (*FlashPromotionSessionListResp, error)
	DeleteFlashPromotionSession(ctx context.Context, id int64) error
}

type FlashPromotionSessionUseCase struct {
	repo FlashPromotionSessionRepo
	log  *log.Helper
}

func NewFlashPromotionSessionUseCase(repo FlashPromotionSessionRepo, logger log.Logger) *FlashPromotionSessionUseCase {
	return &FlashPromotionSessionUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (f FlashPromotionSessionUseCase) CreateFlashPromotionSession(ctx context.Context, session *FlashPromotionSession) error {
	panic("implement me")
}

func (f FlashPromotionSessionUseCase) GetFlashPromotionSession(ctx context.Context, id int64) error {
	panic("implement me")
}

func (f FlashPromotionSessionUseCase) UpdateFlashPromotionSession(ctx context.Context, session *FlashPromotionSession) error {
	panic("implement me")
}

func (f FlashPromotionSessionUseCase) ListFlashPromotionSession(ctx context.Context, req *FlashPromotionSessionListReq) (*FlashPromotionSessionListResp, error) {
	return f.repo.ListFlashPromotionSession(ctx, req)
}

func (f FlashPromotionSessionUseCase) DeleteFlashPromotionSession(ctx context.Context, id int64) error {
	panic("implement me")
}
