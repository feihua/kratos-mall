package model

import (
	"time"
)

// SysLoginLog 系统登录日志
type SysLoginLog struct {
	Id             int64     `gorm:"column:id;AUTO_INCREMENT;primary_key"`         // 编号
	UserName       string    `gorm:"column:user_name"`                             // 用户名
	Status         string    `gorm:"column:status"`                                // 登录状态（online:在线，登录初始状态，方便统计在线人数；login:退出登录后将online置为login；logout:退出登录）
	Ip             string    `gorm:"column:ip"`                                    // IP地址
	CreateBy       string    `gorm:"column:create_by"`                             // 创建人
	CreateTime     time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP"` // 创建时间
	LastUpdateBy   string    `gorm:"column:last_update_by"`                        // 更新人
	LastUpdateTime time.Time `gorm:"column:last_update_time"`                      // 更新时间
}

func (m *SysLoginLog) TableName() string {
	return "sys_login_log"
}
