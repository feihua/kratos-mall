package sms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type CouponHistoryListReq struct {
	Current  int64
	PageSize int64
}

type CouponHistory struct {
	Id             int64
	CouponId       int64
	MemberId       int64
	CouponCode     string
	MemberNickname string // 领取人昵称
	GetType        int    // 获取类型：0->后台赠送；1->主动获取
	CreateTime     string
	UseStatus      int    // 使用状态：0->未使用；1->已使用；2->已过期
	UseTime        string // 使用时间
	OrderId        int64  // 订单编号
	OrderSn        string // 订单号码
}
type CouponHistoryListResp struct {
	Total int64
	List  []*CouponHistory
}

type CouponHistoryRepo interface {
	CreateCouponHistory(context.Context, *CouponHistory) error
	GetCouponHistory(ctx context.Context, id int64) error
	UpdateCouponHistory(context.Context, *CouponHistory) error
	ListCouponHistory(ctx context.Context, req *CouponHistoryListReq) (*CouponHistoryListResp, error)
	DeleteCouponHistory(ctx context.Context, id int64) error
}

type CouponHistoryUseCase struct {
	repo CouponHistoryRepo
	log  *log.Helper
}

func NewCouponHistoryUseCase(repo CouponHistoryRepo, logger log.Logger) *CouponHistoryUseCase {
	return &CouponHistoryUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (c CouponHistoryUseCase) CreateCouponHistory(ctx context.Context, history *CouponHistory) error {
	panic("implement me")
}

func (c CouponHistoryUseCase) GetCouponHistory(ctx context.Context, id int64) error {
	panic("implement me")
}

func (c CouponHistoryUseCase) UpdateCouponHistory(ctx context.Context, history *CouponHistory) error {
	panic("implement me")
}

func (c CouponHistoryUseCase) ListCouponHistory(ctx context.Context, req *CouponHistoryListReq) (*CouponHistoryListResp, error) {
	return c.repo.ListCouponHistory(ctx, req)
}

func (c CouponHistoryUseCase) DeleteCouponHistory(ctx context.Context, id int64) error {
	panic("implement me")
}
