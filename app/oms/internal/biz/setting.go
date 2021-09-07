package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type SettingListReq struct {
	Current  int64
	PageSize int64
}

type Setting struct {
	Id                  int64
	FlashOrderOvertime  int // 秒杀订单超时关闭时间(分)
	NormalOrderOvertime int // 正常订单超时时间(分)
	ConfirmOvertime     int // 发货后自动确认收货时间（天）
	FinishOvertime      int // 自动完成交易时间，不能申请售后（天）
	CommentOvertime     int // 订单完成后自动好评时间（天）
}

type SettingRepo interface {
	CreateSetting(context.Context, *Setting) error
	GetSetting(ctx context.Context, id int64) error
	UpdateSetting(context.Context, *Setting) error
	ListSetting(ctx context.Context, req *SettingListReq) ([]*Setting, error)
	DeleteSetting(ctx context.Context, id int64) error
}

type SettingUseCase struct {
	repo SettingRepo
	log  *log.Helper
}

func NewSettingUseCase(repo SettingRepo, logger log.Logger) *SettingUseCase {
	return &SettingUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (r *SettingUseCase) CreateSetting(ctx context.Context, user *Setting) error {
	panic("implement me")
}

func (r *SettingUseCase) GetSetting(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r *SettingUseCase) UpdateSetting(ctx context.Context, user *Setting) error {
	panic("implement me")
}

func (r *SettingUseCase) ListSetting(ctx context.Context, pageNum, pageSize int64) ([]*Setting, error) {
	panic("implement me")
}

func (r *SettingUseCase) DeleteSetting(ctx context.Context, id int64) error {
	panic("implement me")
}
