package model

import (
	"time"
)

// SysMenu 菜单管理
type SysMenu struct {
	Id             int64     `gorm:"column:id;AUTO_INCREMENT;primary_key"` // 编号
	Name           string    `gorm:"column:name"`                          // 菜单名称
	ParentId       int64     `gorm:"column:parent_id"`                     // 父菜单ID，一级菜单为0
	Url            string    `gorm:"column:url"`
	Perms          string    `gorm:"column:perms"`                                 // 授权(多个用逗号分隔，如：sys:user:add,sys:user:edit)
	Type           int       `gorm:"column:type"`                                  // 类型   0：目录   1：菜单   2：按钮
	Icon           string    `gorm:"column:icon"`                                  // 菜单图标
	OrderNum       int       `gorm:"column:order_num"`                             // 排序
	CreateBy       string    `gorm:"column:create_by"`                             // 创建人
	CreateTime     time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP"` // 创建时间
	LastUpdateBy   string    `gorm:"column:last_update_by"`                        // 更新人
	LastUpdateTime time.Time `gorm:"column:last_update_time"`                      // 更新时间
	DelFlag        int       `gorm:"column:del_flag;default:0"`                    // 是否删除  -1：已删除  0：正常
}

func (m *SysMenu) TableName() string {
	return "sys_menu"
}
