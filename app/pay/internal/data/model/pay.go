package model

import (
	"time"
)

// PayWxMerchants 微信商户信息
type PayWxMerchants struct {
	Id         int       `gorm:"column:id;AUTO_INCREMENT"`
	MerId      string    `gorm:"column:mer_id;NOT NULL"`     // 本地系统商户Id,分配给调用方
	AppId      string    `gorm:"column:app_id;NOT NULL"`     // 应用ID 微信开放平台审核通过的应用APPID
	MchId      string    `gorm:"column:mch_id;NOT NULL"`     // 微信支付分配的商户号
	MchKey     string    `gorm:"column:mch_key;NOT NULL"`    // 微信支付分配的商户秘钥
	PayType    string    `gorm:"column:pay_type;NOT NULL"`   // APP:APP支付 JSAPI:小程序,公众号 MWEB:H5支付
	NotifyUrl  string    `gorm:"column:notify_url;NOT NULL"` // 微信支付回调地址
	Remarks    string    `gorm:"column:remarks;NOT NULL"`    // 备注
	CreateTime time.Time `gorm:"column:create_time;NOT NULL"`
	UpdateTime time.Time `gorm:"column:update_time;NOT NULL"`
}

// PayWxRecord 微信支付记录
type PayWxRecord struct {
	Id         int       `gorm:"column:id;AUTO_INCREMENT"`   // 主键
	BusinessId string    `gorm:"column:businessId;NOT NULL"` // 业务编号
	Amount     string    `gorm:"column:amount;NOT NULL"`     // 金额
	PayType    string    `gorm:"column:pay_type;NOT NULL"`   // 支付类型(APP:APP支付 JSAPI:小程序,公众号 MWEB:H5支付)
	Remarks    string    `gorm:"column:remarks;NOT NULL"`    // 备注
	ReturnCode string    `gorm:"column:return_code;NOT NULL"`
	ReturnMsg  string    `gorm:"column:return_msg;NOT NULL"`
	ResultCode string    `gorm:"column:result_code;NOT NULL"`
	ResultMsg  string    `gorm:"column:result_msg;NOT NULL"`
	PayStatus  int       `gorm:"column:pay_status;default:0;NOT NULL"` // 0：初始化 1：已发送 2：成功 3：失败
	CreateTime time.Time `gorm:"column:create_time;NOT NULL"`          // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;NOT NULL"`          // 更新时间
}

func (m *PayWxRecord) TableName() string {
	return "pay_wx_record"
}
