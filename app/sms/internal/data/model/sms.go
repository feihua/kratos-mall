// Code generated by sql2gorm. DO NOT EDIT.
package model

import (
	"time"
)

// 优惠券表
type SmsCoupon struct {
	Id           int64     `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	Type         int       `gorm:"column:type"` // 优惠券类型；0->全场赠券；1->会员赠券；2->购物赠券；3->注册赠券
	Name         string    `gorm:"column:name"`
	Platform     int       `gorm:"column:platform"`  // 使用平台：0->全部；1->移动；2->PC
	Count        int       `gorm:"column:count"`     // 数量
	Amount       string    `gorm:"column:amount"`    // 金额
	PerLimit     int       `gorm:"column:per_limit"` // 每人限领张数
	MinPoint     string    `gorm:"column:min_point"` // 使用门槛；0表示无门槛
	StartTime    time.Time `gorm:"column:start_time"`
	EndTime      time.Time `gorm:"column:end_time"`
	UseType      int       `gorm:"column:use_type"`      // 使用类型：0->全场通用；1->指定分类；2->指定商品
	Note         string    `gorm:"column:note"`          // 备注
	PublishCount int       `gorm:"column:publish_count"` // 发行数量
	UseCount     int       `gorm:"column:use_count"`     // 已使用数量
	ReceiveCount int       `gorm:"column:receive_count"` // 领取数量
	EnableTime   time.Time `gorm:"column:enable_time"`   // 可以领取的日期
	Code         string    `gorm:"column:code"`          // 优惠码
	MemberLevel  int       `gorm:"column:member_level"`  // 可领取的会员类型：0->无限时
}

func (m *SmsCoupon) TableName() string {
	return "sms_coupon"
}

// 优惠券使用、领取历史表
type SmsCouponHistory struct {
	Id             int64     `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	CouponId       int64     `gorm:"column:coupon_id"`
	MemberId       int64     `gorm:"column:member_id"`
	CouponCode     string    `gorm:"column:coupon_code"`
	MemberNickname string    `gorm:"column:member_nickname"` // 领取人昵称
	GetType        int       `gorm:"column:get_type"`        // 获取类型：0->后台赠送；1->主动获取
	CreateTime     time.Time `gorm:"column:create_time"`
	UseStatus      int       `gorm:"column:use_status"` // 使用状态：0->未使用；1->已使用；2->已过期
	UseTime        time.Time `gorm:"column:use_time"`   // 使用时间
	OrderId        int64     `gorm:"column:order_id"`   // 订单编号
	OrderSn        string    `gorm:"column:order_sn"`   // 订单号码
}

func (m *SmsCouponHistory) TableName() string {
	return "sms_coupon_history"
}

// 优惠券和产品分类关系表
type SmsCouponProductCategoryRelation struct {
	Id                  int64  `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	CouponId            int64  `gorm:"column:coupon_id"`
	ProductCategoryId   int64  `gorm:"column:product_category_id"`
	ProductCategoryName string `gorm:"column:product_category_name"` // 产品分类名称
	ParentCategoryName  string `gorm:"column:parent_category_name"`  // 父分类名称
}

func (m *SmsCouponProductCategoryRelation) TableName() string {
	return "sms_coupon_product_category_relation"
}

// 优惠券和产品的关系表
type SmsCouponProductRelation struct {
	Id          int64  `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	CouponId    int64  `gorm:"column:coupon_id"`
	ProductId   int64  `gorm:"column:product_id"`
	ProductName string `gorm:"column:product_name"` // 商品名称
	ProductSn   string `gorm:"column:product_sn"`   // 商品编码
}

func (m *SmsCouponProductRelation) TableName() string {
	return "sms_coupon_product_relation"
}

// 限时购表
type SmsFlashPromotion struct {
	Id         int64     `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	Title      string    `gorm:"column:title"`
	StartDate  time.Time `gorm:"column:start_date"`  // 开始日期
	EndDate    time.Time `gorm:"column:end_date"`    // 结束日期
	Status     int       `gorm:"column:status"`      // 上下线状态
	CreateTime time.Time `gorm:"column:create_time"` // 秒杀时间段名称
}

func (m *SmsFlashPromotion) TableName() string {
	return "sms_flash_promotion"
}

// 限时购通知记录
type SmsFlashPromotionLog struct {
	Id            int       `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	MemberId      int       `gorm:"column:member_id"`
	ProductId     int64     `gorm:"column:product_id"`
	MemberPhone   string    `gorm:"column:member_phone"`
	ProductName   string    `gorm:"column:product_name"`
	SubscribeTime time.Time `gorm:"column:subscribe_time"` // 会员订阅时间
	SendTime      time.Time `gorm:"column:send_time"`
}

func (m *SmsFlashPromotionLog) TableName() string {
	return "sms_flash_promotion_log"
}

// 商品限时购与商品关系表
type SmsFlashPromotionProductRelation struct {
	Id                      int64  `gorm:"column:id;AUTO_INCREMENT;primary_key"` // 编号
	FlashPromotionId        int64  `gorm:"column:flash_promotion_id"`
	FlashPromotionSessionId int64  `gorm:"column:flash_promotion_session_id"` // 编号
	ProductId               int64  `gorm:"column:product_id"`
	FlashPromotionPrice     string `gorm:"column:flash_promotion_price"` // 限时购价格
	FlashPromotionCount     int    `gorm:"column:flash_promotion_count"` // 限时购数量
	FlashPromotionLimit     int    `gorm:"column:flash_promotion_limit"` // 每人限购数量
	Sort                    int    `gorm:"column:sort"`                  // 排序
}

func (m *SmsFlashPromotionProductRelation) TableName() string {
	return "sms_flash_promotion_product_relation"
}

// 限时购场次表
type SmsFlashPromotionSession struct {
	Id         int64     `gorm:"column:id;AUTO_INCREMENT;primary_key"` // 编号
	Name       string    `gorm:"column:name"`                          // 场次名称
	StartTime  time.Time `gorm:"column:start_time"`                    // 每日开始时间
	EndTime    time.Time `gorm:"column:end_time"`                      // 每日结束时间
	Status     int       `gorm:"column:status"`                        // 启用状态：0->不启用；1->启用
	CreateTime time.Time `gorm:"column:create_time"`                   // 创建时间
}

func (m *SmsFlashPromotionSession) TableName() string {
	return "sms_flash_promotion_session"
}

// 首页轮播广告表
type SmsHomeAdvertise struct {
	Id         int64     `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	Name       string    `gorm:"column:name"`
	Type       int       `gorm:"column:type"` // 轮播位置：0->PC首页轮播；1->app首页轮播
	Pic        string    `gorm:"column:pic"`
	StartTime  time.Time `gorm:"column:start_time"`
	EndTime    time.Time `gorm:"column:end_time"`
	Status     int       `gorm:"column:status"`         // 上下线状态：0->下线；1->上线
	ClickCount int       `gorm:"column:click_count"`    // 点击数
	OrderCount int       `gorm:"column:order_count"`    // 下单数
	Url        string    `gorm:"column:url"`            // 链接地址
	Note       string    `gorm:"column:note"`           // 备注
	Sort       int       `gorm:"column:sort;default:0"` // 排序
}

func (m *SmsHomeAdvertise) TableName() string {
	return "sms_home_advertise"
}

// 首页推荐品牌表
type SmsHomeBrand struct {
	Id              int64  `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	BrandId         int64  `gorm:"column:brand_id"`
	BrandName       string `gorm:"column:brand_name"`
	RecommendStatus int    `gorm:"column:recommend_status"`
	Sort            int    `gorm:"column:sort"`
}

func (m *SmsHomeBrand) TableName() string {
	return "sms_home_brand"
}

// 新鲜好物表
type SmsHomeNewProduct struct {
	Id              int64  `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	ProductId       int64  `gorm:"column:product_id"`
	ProductName     string `gorm:"column:product_name"`
	RecommendStatus int    `gorm:"column:recommend_status"`
	Sort            int    `gorm:"column:sort"`
}

func (m *SmsHomeNewProduct) TableName() string {
	return "sms_home_new_product"
}

// 人气推荐商品表
type SmsHomeRecommendProduct struct {
	Id              int64  `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	ProductId       int64  `gorm:"column:product_id"`
	ProductName     string `gorm:"column:product_name"`
	RecommendStatus int    `gorm:"column:recommend_status"`
	Sort            int    `gorm:"column:sort"`
}

func (m *SmsHomeRecommendProduct) TableName() string {
	return "sms_home_recommend_product"
}

// 首页推荐专题表
type SmsHomeRecommendSubject struct {
	Id              int64  `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	SubjectId       int64  `gorm:"column:subject_id"`
	SubjectName     string `gorm:"column:subject_name"`
	RecommendStatus int    `gorm:"column:recommend_status"`
	Sort            int    `gorm:"column:sort"`
}

func (m *SmsHomeRecommendSubject) TableName() string {
	return "sms_home_recommend_subject"
}
