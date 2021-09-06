package model

import (
	"time"
)

// SysUser 用户管理
type SysUser struct {
	Id             int64     `gorm:"column:id;AUTO_INCREMENT;primary_key"`         // 编号
	Name           string    `gorm:"column:name;NOT NULL"`                         // 用户名
	NickName       string    `gorm:"column:nick_name"`                             // 昵称
	Avatar         string    `gorm:"column:avatar"`                                // 头像
	Password       string    `gorm:"column:password"`                              // 密码
	Salt           string    `gorm:"column:salt"`                                  // 加密盐
	Email          string    `gorm:"column:email"`                                 // 邮箱
	Mobile         string    `gorm:"column:mobile"`                                // 手机号
	Status         int       `gorm:"column:status"`                                // 状态  0：禁用   1：正常
	DeptId         int64     `gorm:"column:dept_id"`                               // 机构ID
	CreateBy       string    `gorm:"column:create_by"`                             // 创建人
	CreateTime     time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP"` // 创建时间
	LastUpdateBy   string    `gorm:"column:last_update_by"`                        // 更新人
	LastUpdateTime time.Time `gorm:"column:last_update_time"`                      // 更新时间
	DelFlag        int       `gorm:"column:del_flag;default:0"`                    // 是否删除  -1：已删除  0：正常
	JobId          int       `gorm:"column:job_id"`                                // 岗位Id
}

func (m *SysUser) TableName() string {
	return "sys_user"
}
