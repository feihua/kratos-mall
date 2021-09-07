package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type FlashPromotionSessionListReq struct {
	Current  int64
	PageSize int64
}

type FlashPromotionSession struct {
	Id         int64     // 编号
	Name       string    // 场次名称
	StartTime  time.Time // 每日开始时间
	EndTime    time.Time // 每日结束时间
	Status     int       // 启用状态：0->不启用；1->启用
	CreateTime time.Time // 创建时间
}

type FlashPromotionSessionRepo interface {
	CreateFlashPromotionSession(context.Context, *FlashPromotionSession) error
	GetFlashPromotionSession(ctx context.Context, id int64) error
	UpdateFlashPromotionSession(context.Context, *FlashPromotionSession) error
	ListFlashPromotionSession(ctx context.Context, req *FlashPromotionSessionListReq) ([]*FlashPromotionSession, error)
	DeleteFlashPromotionSession(ctx context.Context, id int64) error
}

type FlashPromotionSessionUseCase struct {
	repo FlashPromotionSessionRepo
	log  *log.Helper
}

func NewFlashPromotionSessionUseCase(repo FlashPromotionSessionRepo, logger log.Logger) *FlashPromotionSessionUseCase {
	return &FlashPromotionSessionUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (r *FlashPromotionSessionUseCase) CreateFlashPromotionSession(ctx context.Context, user *FlashPromotionSession) error {
	panic("implement me")
}

func (r *FlashPromotionSessionUseCase) GetFlashPromotionSession(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r *FlashPromotionSessionUseCase) UpdateFlashPromotionSession(ctx context.Context, user *FlashPromotionSession) error {
	panic("implement me")
}

func (r *FlashPromotionSessionUseCase) ListFlashPromotionSession(ctx context.Context, pageNum, pageSize int64) ([]*FlashPromotionSession, error) {
	panic("implement me")
}

func (r *FlashPromotionSessionUseCase) DeleteFlashPromotionSession(ctx context.Context, id int64) error {
	panic("implement me")
}
