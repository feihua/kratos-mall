package model

import (
	"time"
)

// OmsCartItem 购物车表
type OmsCartItem struct {
	Id                int64     `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	ProductId         int64     `gorm:"column:product_id"`
	ProductSkuId      int64     `gorm:"column:product_sku_id"`
	MemberId          int64     `gorm:"column:member_id"`
	Quantity          int       `gorm:"column:quantity"`                // 购买数量
	Price             string    `gorm:"column:price"`                   // 添加到购物车的价格
	ProductPic        string    `gorm:"column:product_pic"`             // 商品主图
	ProductName       string    `gorm:"column:product_name"`            // 商品名称
	ProductSubTitle   string    `gorm:"column:product_sub_title"`       // 商品副标题（卖点）
	ProductSkuCode    string    `gorm:"column:product_sku_code"`        // 商品sku条码
	MemberNickname    string    `gorm:"column:member_nickname"`         // 会员昵称
	CreateDate        time.Time `gorm:"column:create_date"`             // 创建时间
	ModifyDate        time.Time `gorm:"column:modify_date"`             // 修改时间
	DeleteStatus      int       `gorm:"column:delete_status;default:0"` // 是否删除
	ProductCategoryId int64     `gorm:"column:product_category_id"`     // 商品分类
	ProductBrand      string    `gorm:"column:product_brand"`
	ProductSn         string    `gorm:"column:product_sn"`
	ProductAttr       string    `gorm:"column:product_attr"` // 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]
}

func (m *OmsCartItem) TableName() string {
	return "oms_cart_item"
}

// OmsCompanyAddress 公司收发货地址表
type OmsCompanyAddress struct {
	Id            int64  `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	AddressName   string `gorm:"column:address_name"`   // 地址名称
	SendStatus    int    `gorm:"column:send_status"`    // 默认发货地址：0->否；1->是
	ReceiveStatus int    `gorm:"column:receive_status"` // 是否默认收货地址：0->否；1->是
	Name          string `gorm:"column:name"`           // 收发货人姓名
	Phone         string `gorm:"column:phone"`          // 收货人电话
	Province      string `gorm:"column:province"`       // 省/直辖市
	City          string `gorm:"column:city"`           // 市
	Region        string `gorm:"column:region"`         // 区
	DetailAddress string `gorm:"column:detail_address"` // 详细地址
}

func (m *OmsCompanyAddress) TableName() string {
	return "oms_company_address"
}

// OmsOrder 订单表
type OmsOrder struct {
	Id                    int64     `gorm:"column:id;AUTO_INCREMENT;primary_key"` // 订单id
	MemberId              int64     `gorm:"column:member_id;NOT NULL"`
	CouponId              int64     `gorm:"column:coupon_id"`
	OrderSn               string    `gorm:"column:order_sn"`                         // 订单编号
	CreateTime            time.Time `gorm:"column:create_time"`                      // 提交时间
	MemberUsername        string    `gorm:"column:member_username"`                  // 用户帐号
	TotalAmount           string    `gorm:"column:total_amount"`                     // 订单总金额
	PayAmount             string    `gorm:"column:pay_amount"`                       // 应付金额（实际支付金额）
	FreightAmount         string    `gorm:"column:freight_amount"`                   // 运费金额
	PromotionAmount       string    `gorm:"column:promotion_amount"`                 // 促销优化金额（促销价、满减、阶梯价）
	IntegrationAmount     string    `gorm:"column:integration_amount"`               // 积分抵扣金额
	CouponAmount          string    `gorm:"column:coupon_amount"`                    // 优惠券抵扣金额
	DiscountAmount        string    `gorm:"column:discount_amount"`                  // 管理员后台调整订单使用的折扣金额
	PayType               int       `gorm:"column:pay_type"`                         // 支付方式：0->未支付；1->支付宝；2->微信
	SourceType            int       `gorm:"column:source_type"`                      // 订单来源：0->PC订单；1->app订单
	Status                int       `gorm:"column:status"`                           // 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
	OrderType             int       `gorm:"column:order_type"`                       // 订单类型：0->正常订单；1->秒杀订单
	DeliveryCompany       string    `gorm:"column:delivery_company"`                 // 物流公司(配送方式)
	DeliverySn            string    `gorm:"column:delivery_sn"`                      // 物流单号
	AutoConfirmDay        int       `gorm:"column:auto_confirm_day"`                 // 自动确认时间（天）
	Integration           int       `gorm:"column:integration"`                      // 可以获得的积分
	Growth                int       `gorm:"column:growth"`                           // 可以活动的成长值
	PromotionInfo         string    `gorm:"column:promotion_info"`                   // 活动信息
	BillType              int       `gorm:"column:bill_type"`                        // 发票类型：0->不开发票；1->电子发票；2->纸质发票
	BillHeader            string    `gorm:"column:bill_header"`                      // 发票抬头
	BillContent           string    `gorm:"column:bill_content"`                     // 发票内容
	BillReceiverPhone     string    `gorm:"column:bill_receiver_phone"`              // 收票人电话
	BillReceiverEmail     string    `gorm:"column:bill_receiver_email"`              // 收票人邮箱
	ReceiverName          string    `gorm:"column:receiver_name;NOT NULL"`           // 收货人姓名
	ReceiverPhone         string    `gorm:"column:receiver_phone;NOT NULL"`          // 收货人电话
	ReceiverPostCode      string    `gorm:"column:receiver_post_code"`               // 收货人邮编
	ReceiverProvince      string    `gorm:"column:receiver_province"`                // 省份/直辖市
	ReceiverCity          string    `gorm:"column:receiver_city"`                    // 城市
	ReceiverRegion        string    `gorm:"column:receiver_region"`                  // 区
	ReceiverDetailAddress string    `gorm:"column:receiver_detail_address"`          // 详细地址
	Note                  string    `gorm:"column:note"`                             // 订单备注
	ConfirmStatus         int       `gorm:"column:confirm_status"`                   // 确认收货状态：0->未确认；1->已确认
	DeleteStatus          int       `gorm:"column:delete_status;default:0;NOT NULL"` // 删除状态：0->未删除；1->已删除
	UseIntegration        int       `gorm:"column:use_integration"`                  // 下单时使用的积分
	PaymentTime           time.Time `gorm:"column:payment_time"`                     // 支付时间
	DeliveryTime          time.Time `gorm:"column:delivery_time"`                    // 发货时间
	ReceiveTime           time.Time `gorm:"column:receive_time"`                     // 确认收货时间
	CommentTime           time.Time `gorm:"column:comment_time"`                     // 评价时间
	ModifyTime            time.Time `gorm:"column:modify_time"`                      // 修改时间
}

func (m *OmsOrder) TableName() string {
	return "oms_order"
}

// OmsOrderItem 订单中所包含的商品
type OmsOrderItem struct {
	Id                int64  `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	OrderId           int64  `gorm:"column:order_id"` // 订单id
	OrderSn           string `gorm:"column:order_sn"` // 订单编号
	ProductId         int64  `gorm:"column:product_id"`
	ProductPic        string `gorm:"column:product_pic"`
	ProductName       string `gorm:"column:product_name"`
	ProductBrand      string `gorm:"column:product_brand"`
	ProductSn         string `gorm:"column:product_sn"`
	ProductPrice      string `gorm:"column:product_price"`       // 销售价格
	ProductQuantity   int    `gorm:"column:product_quantity"`    // 购买数量
	ProductSkuId      int64  `gorm:"column:product_sku_id"`      // 商品sku编号
	ProductSkuCode    string `gorm:"column:product_sku_code"`    // 商品sku条码
	ProductCategoryId int64  `gorm:"column:product_category_id"` // 商品分类id
	PromotionName     string `gorm:"column:promotion_name"`      // 商品促销名称
	PromotionAmount   string `gorm:"column:promotion_amount"`    // 商品促销分解金额
	CouponAmount      string `gorm:"column:coupon_amount"`       // 优惠券优惠分解金额
	IntegrationAmount string `gorm:"column:integration_amount"`  // 积分优惠分解金额
	RealAmount        string `gorm:"column:real_amount"`         // 该商品经过优惠后的分解金额
	GiftIntegration   int    `gorm:"column:gift_integration;default:0"`
	GiftGrowth        int    `gorm:"column:gift_growth;default:0"`
	ProductAttr       string `gorm:"column:product_attr"` // 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]
}

func (m *OmsOrderItem) TableName() string {
	return "oms_order_item"
}

// OmsOrderOperateHistory 订单操作历史记录
type OmsOrderOperateHistory struct {
	Id          int64     `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	OrderId     int64     `gorm:"column:order_id"`     // 订单id
	OperateMan  string    `gorm:"column:operate_man"`  // 操作人：用户；系统；后台管理员
	CreateTime  time.Time `gorm:"column:create_time"`  // 操作时间
	OrderStatus int       `gorm:"column:order_status"` // 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
	Note        string    `gorm:"column:note"`         // 备注
}

func (m *OmsOrderOperateHistory) TableName() string {
	return "oms_order_operate_history"
}

// OmsOrderReturnApply 订单退货申请
type OmsOrderReturnApply struct {
	Id               int64     `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	OrderId          int64     `gorm:"column:order_id"`           // 订单id
	CompanyAddressId int64     `gorm:"column:company_address_id"` // 收货地址表id
	ProductId        int64     `gorm:"column:product_id"`         // 退货商品id
	OrderSn          string    `gorm:"column:order_sn"`           // 订单编号
	CreateTime       time.Time `gorm:"column:create_time"`        // 申请时间
	MemberUsername   string    `gorm:"column:member_username"`    // 会员用户名
	ReturnAmount     string    `gorm:"column:return_amount"`      // 退款金额
	ReturnName       string    `gorm:"column:return_name"`        // 退货人姓名
	ReturnPhone      string    `gorm:"column:return_phone"`       // 退货人电话
	Status           int       `gorm:"column:status"`             // 申请状态：0->待处理；1->退货中；2->已完成；3->已拒绝
	HandleTime       time.Time `gorm:"column:handle_time"`        // 处理时间
	ProductPic       string    `gorm:"column:product_pic"`        // 商品图片
	ProductName      string    `gorm:"column:product_name"`       // 商品名称
	ProductBrand     string    `gorm:"column:product_brand"`      // 商品品牌
	ProductAttr      string    `gorm:"column:product_attr"`       // 商品销售属性：颜色：红色；尺码：xl;
	ProductCount     int       `gorm:"column:product_count"`      // 退货数量
	ProductPrice     string    `gorm:"column:product_price"`      // 商品单价
	ProductRealPrice string    `gorm:"column:product_real_price"` // 商品实际支付单价
	Reason           string    `gorm:"column:reason"`             // 原因
	Description      string    `gorm:"column:description"`        // 描述
	ProofPics        string    `gorm:"column:proof_pics"`         // 凭证图片，以逗号隔开
	HandleNote       string    `gorm:"column:handle_note"`        // 处理备注
	HandleMan        string    `gorm:"column:handle_man"`         // 处理人员
	ReceiveMan       string    `gorm:"column:receive_man"`        // 收货人
	ReceiveTime      time.Time `gorm:"column:receive_time"`       // 收货时间
	ReceiveNote      string    `gorm:"column:receive_note"`       // 收货备注
}

func (m *OmsOrderReturnApply) TableName() string {
	return "oms_order_return_apply"
}

// OmsOrderReturnReason 退货原因表
type OmsOrderReturnReason struct {
	Id         int64     `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	Name       string    `gorm:"column:name"` // 退货类型
	Sort       int       `gorm:"column:sort"`
	Status     int       `gorm:"column:status"`      // 状态：0->不启用；1->启用
	CreateTime time.Time `gorm:"column:create_time"` // 添加时间
}

func (m *OmsOrderReturnReason) TableName() string {
	return "oms_order_return_reason"
}

// OmsOrderSetting 订单设置表
type OmsOrderSetting struct {
	Id                  int64 `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	FlashOrderOvertime  int   `gorm:"column:flash_order_overtime"`  // 秒杀订单超时关闭时间(分)
	NormalOrderOvertime int   `gorm:"column:normal_order_overtime"` // 正常订单超时时间(分)
	ConfirmOvertime     int   `gorm:"column:confirm_overtime"`      // 发货后自动确认收货时间（天）
	FinishOvertime      int   `gorm:"column:finish_overtime"`       // 自动完成交易时间，不能申请售后（天）
	CommentOvertime     int   `gorm:"column:comment_overtime"`      // 订单完成后自动好评时间（天）
}

func (m *OmsOrderSetting) TableName() string {
	return "oms_order_setting"
}
