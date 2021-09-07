// Code generated by sql2gorm. DO NOT EDIT.
package model

import (
	"time"
)

// 成长值变化历史记录表
type UmsGrowthChangeHistory struct {
	Id          int64     `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	MemberId    int64     `gorm:"column:member_id"`
	CreateTime  time.Time `gorm:"column:create_time"`
	ChangeType  int       `gorm:"column:change_type"`  // 改变类型：0->增加；1->减少
	ChangeCount int       `gorm:"column:change_count"` // 积分改变数量
	OperateMan  string    `gorm:"column:operate_man"`  // 操作人员
	OperateNote string    `gorm:"column:operate_note"` // 操作备注
	SourceType  int       `gorm:"column:source_type"`  // 积分来源：0->购物；1->管理员修改
}

func (m *UmsGrowthChangeHistory) TableName() string {
	return "ums_growth_change_history"
}

// 积分变化历史记录表
type UmsIntegrationChangeHistory struct {
	Id          int64     `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	MemberId    int64     `gorm:"column:member_id"`
	CreateTime  time.Time `gorm:"column:create_time"`
	ChangeType  int       `gorm:"column:change_type"`  // 改变类型：0->增加；1->减少
	ChangeCount int       `gorm:"column:change_count"` // 积分改变数量
	OperateMan  string    `gorm:"column:operate_man"`  // 操作人员
	OperateNote string    `gorm:"column:operate_note"` // 操作备注
	SourceType  int       `gorm:"column:source_type"`  // 积分来源：0->购物；1->管理员修改
}

func (m *UmsIntegrationChangeHistory) TableName() string {
	return "ums_integration_change_history"
}

// 积分消费设置
type UmsIntegrationConsumeSetting struct {
	Id                 int64 `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	DeductionPerAmount int   `gorm:"column:deduction_per_amount"`  // 每一元需要抵扣的积分数量
	MaxPercentPerOrder int   `gorm:"column:max_percent_per_order"` // 每笔订单最高抵用百分比
	UseUnit            int   `gorm:"column:use_unit"`              // 每次使用积分最小单位100
	CouponStatus       int   `gorm:"column:coupon_status"`         // 是否可以和优惠券同用；0->不可以；1->可以
}

func (m *UmsIntegrationConsumeSetting) TableName() string {
	return "ums_integration_consume_setting"
}

// 会员表
type UmsMember struct {
	Id                    int64     `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	MemberLevelId         int64     `gorm:"column:member_level_id"`
	Username              string    `gorm:"column:username"`               // 用户名
	Password              string    `gorm:"column:password"`               // 密码
	Nickname              string    `gorm:"column:nickname"`               // 昵称
	Phone                 string    `gorm:"column:phone"`                  // 手机号码
	Status                int       `gorm:"column:status"`                 // 帐号启用状态:0->禁用；1->启用
	CreateTime            time.Time `gorm:"column:create_time"`            // 注册时间
	Icon                  string    `gorm:"column:icon"`                   // 头像
	Gender                int       `gorm:"column:gender"`                 // 性别：0->未知；1->男；2->女
	Birthday              time.Time `gorm:"column:birthday"`               // 生日
	City                  string    `gorm:"column:city"`                   // 所做城市
	Job                   string    `gorm:"column:job"`                    // 职业
	PersonalizedSignature string    `gorm:"column:personalized_signature"` // 个性签名
	SourceType            int       `gorm:"column:source_type"`            // 用户来源
	Integration           int       `gorm:"column:integration"`            // 积分
	Growth                int       `gorm:"column:growth"`                 // 成长值
	LuckeyCount           int       `gorm:"column:luckey_count"`           // 剩余抽奖次数
	HistoryIntegration    int       `gorm:"column:history_integration"`    // 历史积分数量
}

func (m *UmsMember) TableName() string {
	return "ums_member"
}

// 会员等级表
type UmsMemberLevel struct {
	Id                    int64  `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	Name                  string `gorm:"column:name"`
	GrowthPoint           int    `gorm:"column:growth_point"`
	DefaultStatus         int    `gorm:"column:default_status"`          // 是否为默认等级：0->不是；1->是
	FreeFreightPoint      string `gorm:"column:free_freight_point"`      // 免运费标准
	CommentGrowthPoint    int    `gorm:"column:comment_growth_point"`    // 每次评价获取的成长值
	PriviledgeFreeFreight int    `gorm:"column:priviledge_free_freight"` // 是否有免邮特权
	PriviledgeSignIn      int    `gorm:"column:priviledge_sign_in"`      // 是否有签到特权
	PriviledgeComment     int    `gorm:"column:priviledge_comment"`      // 是否有评论获奖励特权
	PriviledgePromotion   int    `gorm:"column:priviledge_promotion"`    // 是否有专享活动特权
	PriviledgeMemberPrice int    `gorm:"column:priviledge_member_price"` // 是否有会员价格特权
	PriviledgeBirthday    int    `gorm:"column:priviledge_birthday"`     // 是否有生日特权
	Note                  string `gorm:"column:note"`
}

func (m *UmsMemberLevel) TableName() string {
	return "ums_member_level"
}

// 会员登录记录
type UmsMemberLoginLog struct {
	Id         int64     `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	MemberId   int64     `gorm:"column:member_id"`
	CreateTime time.Time `gorm:"column:create_time"`
	Ip         string    `gorm:"column:ip"`
	City       string    `gorm:"column:city"`
	LoginType  int       `gorm:"column:login_type"` // 登录类型：0->PC；1->android;2->ios;3->小程序
	Province   string    `gorm:"column:province"`
}

func (m *UmsMemberLoginLog) TableName() string {
	return "ums_member_login_log"
}

// 用户和标签关系表
type UmsMemberMemberTagRelation struct {
	Id       int64 `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	MemberId int64 `gorm:"column:member_id"`
	TagId    int64 `gorm:"column:tag_id"`
}

func (m *UmsMemberMemberTagRelation) TableName() string {
	return "ums_member_member_tag_relation"
}

// 会员与产品分类关系表（用户喜欢的分类）
type UmsMemberProductCategoryRelation struct {
	Id                int64 `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	MemberId          int64 `gorm:"column:member_id"`
	ProductCategoryId int64 `gorm:"column:product_category_id"`
}

func (m *UmsMemberProductCategoryRelation) TableName() string {
	return "ums_member_product_category_relation"
}

// 会员收货地址表
type UmsMemberReceiveAddress struct {
	Id            int64  `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	MemberId      int64  `gorm:"column:member_id"`
	Name          string `gorm:"column:name"` // 收货人名称
	PhoneNumber   string `gorm:"column:phone_number"`
	DefaultStatus int    `gorm:"column:default_status"` // 是否为默认
	PostCode      string `gorm:"column:post_code"`      // 邮政编码
	Province      string `gorm:"column:province"`       // 省份/直辖市
	City          string `gorm:"column:city"`           // 城市
	Region        string `gorm:"column:region"`         // 区
	DetailAddress string `gorm:"column:detail_address"` // 详细地址(街道)
}

func (m *UmsMemberReceiveAddress) TableName() string {
	return "ums_member_receive_address"
}

// 会员积分成长规则表
type UmsMemberRuleSetting struct {
	Id                int64  `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	ContinueSignDay   int    `gorm:"column:continue_sign_day"`   // 连续签到天数
	ContinueSignPoint int    `gorm:"column:continue_sign_point"` // 连续签到赠送数量
	ConsumePerPoint   string `gorm:"column:consume_per_point"`   // 每消费多少元获取1个点
	LowOrderAmount    string `gorm:"column:low_order_amount"`    // 最低获取点数的订单金额
	MaxPointPerOrder  int    `gorm:"column:max_point_per_order"` // 每笔订单最高获取点数
	Type              int    `gorm:"column:type"`                // 类型：0->积分规则；1->成长值规则
}

func (m *UmsMemberRuleSetting) TableName() string {
	return "ums_member_rule_setting"
}

// 会员统计信息
type UmsMemberStatisticsInfo struct {
	Id                  int64     `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	MemberId            int64     `gorm:"column:member_id"`
	ConsumeAmount       string    `gorm:"column:consume_amount"`     // 累计消费金额
	OrderCount          int       `gorm:"column:order_count"`        // 订单数量
	CouponCount         int       `gorm:"column:coupon_count"`       // 优惠券数量
	CommentCount        int       `gorm:"column:comment_count"`      // 评价数
	ReturnOrderCount    int       `gorm:"column:return_order_count"` // 退货数量
	LoginCount          int       `gorm:"column:login_count"`        // 登录次数
	AttendCount         int       `gorm:"column:attend_count"`       // 关注数量
	FansCount           int       `gorm:"column:fans_count"`         // 粉丝数量
	CollectProductCount int       `gorm:"column:collect_product_count"`
	CollectSubjectCount int       `gorm:"column:collect_subject_count"`
	CollectTopicCount   int       `gorm:"column:collect_topic_count"`
	CollectCommentCount int       `gorm:"column:collect_comment_count"`
	InviteFriendCount   int       `gorm:"column:invite_friend_count"`
	RecentOrderTime     time.Time `gorm:"column:recent_order_time"` // 最后一次下订单时间
}

func (m *UmsMemberStatisticsInfo) TableName() string {
	return "ums_member_statistics_info"
}

// 用户标签表
type UmsMemberTag struct {
	Id                int64  `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	Name              string `gorm:"column:name"`
	FinishOrderCount  int    `gorm:"column:finish_order_count"`  // 自动打标签完成订单数量
	FinishOrderAmount string `gorm:"column:finish_order_amount"` // 自动打标签完成订单金额
}

func (m *UmsMemberTag) TableName() string {
	return "ums_member_tag"
}

// 会员任务表
type UmsMemberTask struct {
	Id           int64  `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	Name         string `gorm:"column:name"`
	Growth       int    `gorm:"column:growth"`       // 赠送成长值
	Intergration int    `gorm:"column:intergration"` // 赠送积分
	Type         int    `gorm:"column:type"`         // 任务类型：0->新手任务；1->日常任务
}

func (m *UmsMemberTask) TableName() string {
	return "ums_member_task"
}