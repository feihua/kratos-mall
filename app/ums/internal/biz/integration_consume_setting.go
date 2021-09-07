package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type ConsumeSettingListReq struct {
	Current  int64
	PageSize int64
}

type ConsumeSetting struct {
	Id                 int64
	DeductionPerAmount int // 每一元需要抵扣的积分数量
	MaxPercentPerOrder int // 每笔订单最高抵用百分比
	UseUnit            int // 每次使用积分最小单位100
	CouponStatus       int // 是否可以和优惠券同用；0->不可以；1->可以
}

type ConsumeSettingRepo interface {
	CreateConsumeSetting(context.Context, *ConsumeSetting) error
	GetConsumeSetting(ctx context.Context, id int64) error
	UpdateConsumeSetting(context.Context, *ConsumeSetting) error
	ListConsumeSetting(ctx context.Context, req *ConsumeSettingListReq) ([]*ConsumeSetting, error)
	DeleteConsumeSetting(ctx context.Context, id int64) error
}

type ConsumeSettingUseCase struct {
	repo ConsumeSettingRepo
	log  *log.Helper
}

func NewConsumeSettingUseCase(repo ConsumeSettingRepo, logger log.Logger) *ConsumeSettingUseCase {
	return &ConsumeSettingUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (r *ConsumeSettingUseCase) CreateConsumeSetting(ctx context.Context, user *ConsumeSetting) error {
	panic("implement me")
}

func (r *ConsumeSettingUseCase) GetConsumeSetting(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r *ConsumeSettingUseCase) UpdateConsumeSetting(ctx context.Context, user *ConsumeSetting) error {
	panic("implement me")
}

func (r *ConsumeSettingUseCase) ListConsumeSetting(ctx context.Context, pageNum, pageSize int64) ([]*ConsumeSetting, error) {
	panic("implement me")
}

func (r *ConsumeSettingUseCase) DeleteConsumeSetting(ctx context.Context, id int64) error {
	panic("implement me")
}
