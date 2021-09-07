package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type CouponListReq struct {
	Current  int64
	PageSize int64
}

type Coupon struct {
	Id           int64
	Type         int // 优惠券类型；0->全场赠券；1->会员赠券；2->购物赠券；3->注册赠券
	Name         string
	Platform     int    // 使用平台：0->全部；1->移动；2->PC
	Count        int    // 数量
	Amount       string // 金额
	PerLimit     int    // 每人限领张数
	MinPoint     string // 使用门槛；0表示无门槛
	StartTime    time.Time
	EndTime      time.Time
	UseType      int       // 使用类型：0->全场通用；1->指定分类；2->指定商品
	Note         string    // 备注
	PublishCount int       // 发行数量
	UseCount     int       // 已使用数量
	ReceiveCount int       // 领取数量
	EnableTime   time.Time // 可以领取的日期
	Code         string    // 优惠码
	MemberLevel  int       // 可领取的会员类型：0->无限时
}

type CouponRepo interface {
	CreateCoupon(context.Context, *Coupon) error
	GetCoupon(ctx context.Context, id int64) error
	UpdateCoupon(context.Context, *Coupon) error
	ListCoupon(ctx context.Context, req *CouponListReq) ([]*Coupon, error)
	DeleteCoupon(ctx context.Context, id int64) error
}

type CouponUseCase struct {
	repo CouponRepo
	log  *log.Helper
}

func NewCouponUseCase(repo CouponRepo, logger log.Logger) *CouponUseCase {
	return &CouponUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (r *CouponUseCase) CreateCoupon(ctx context.Context, user *Coupon) error {
	panic("implement me")
}

func (r *CouponUseCase) GetCoupon(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r *CouponUseCase) UpdateCoupon(ctx context.Context, user *Coupon) error {
	panic("implement me")
}

func (r *CouponUseCase) ListCoupon(ctx context.Context, pageNum, pageSize int64) ([]*Coupon, error) {
	panic("implement me")
}

func (r *CouponUseCase) DeleteCoupon(ctx context.Context, id int64) error {
	panic("implement me")
}
