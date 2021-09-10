package oms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type OrderListReq struct {
	Current  int64
	PageSize int64
}

type Order struct {
	Id                    int64 // 订单id
	MemberId              int64
	CouponId              int64
	OrderSn               string // 订单编号
	CreateTime            string // 提交时间
	MemberUsername        string // 用户帐号
	TotalAmount           string // 订单总金额
	PayAmount             string // 应付金额（实际支付金额）
	FreightAmount         string // 运费金额
	PromotionAmount       string // 促销优化金额（促销价、满减、阶梯价）
	IntegrationAmount     string // 积分抵扣金额
	CouponAmount          string // 优惠券抵扣金额
	DiscountAmount        string // 管理员后台调整订单使用的折扣金额
	PayType               int    // 支付方式：0->未支付；1->支付宝；2->微信
	SourceType            int    // 订单来源：0->PC订单；1->app订单
	Status                int    // 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
	OrderType             int    // 订单类型：0->正常订单；1->秒杀订单
	DeliveryCompany       string // 物流公司(配送方式)
	DeliverySn            string // 物流单号
	AutoConfirmDay        int    // 自动确认时间（天）
	Integration           int    // 可以获得的积分
	Growth                int    // 可以活动的成长值
	PromotionInfo         string // 活动信息
	BillType              int    // 发票类型：0->不开发票；1->电子发票；2->纸质发票
	BillHeader            string // 发票抬头
	BillContent           string // 发票内容
	BillReceiverPhone     string // 收票人电话
	BillReceiverEmail     string // 收票人邮箱
	ReceiverName          string // 收货人姓名
	ReceiverPhone         string // 收货人电话
	ReceiverPostCode      string // 收货人邮编
	ReceiverProvince      string // 省份/直辖市
	ReceiverCity          string // 城市
	ReceiverRegion        string // 区
	ReceiverDetailAddress string // 详细地址
	Note                  string // 订单备注
	ConfirmStatus         int    // 确认收货状态：0->未确认；1->已确认
	DeleteStatus          int    // 删除状态：0->未删除；1->已删除
	UseIntegration        int    // 下单时使用的积分
	PaymentTime           string // 支付时间
	DeliveryTime          string // 发货时间
	ReceiveTime           string // 确认收货时间
	CommentTime           string // 评价时间
	ModifyTime            string // 修改时间
}

type OrderRepo interface {
	CreateOrder(context.Context, *Order) error
	GetOrder(ctx context.Context, id int64) error
	UpdateOrder(context.Context, *Order) error
	ListOrder(ctx context.Context, req *OrderListReq) ([]*Order, error)
	DeleteOrder(ctx context.Context, id int64) error
}

type OrderUseCase struct {
	repo OrderRepo
	log  *log.Helper
}

func NewOrderUseCase(repo OrderRepo, logger log.Logger) *OrderUseCase {
	return &OrderUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (r *OrderUseCase) CreateOrder(ctx context.Context, user *Order) error {
	panic("implement me")
}

func (r *OrderUseCase) GetOrder(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r *OrderUseCase) UpdateOrder(ctx context.Context, user *Order) error {
	panic("implement me")
}

func (r *OrderUseCase) ListOrder(ctx context.Context, req *OrderListReq) ([]*Order, error) {
	return r.repo.ListOrder(ctx, req)
}

func (r *OrderUseCase) DeleteOrder(ctx context.Context, id int64) error {
	panic("implement me")
}
