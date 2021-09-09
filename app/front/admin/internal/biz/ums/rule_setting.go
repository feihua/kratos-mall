package ums

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type RuleSettingListReq struct {
	Current  int64
	PageSize int64
}

type RuleSetting struct {
	Id                int64
	ContinueSignDay   int    // 连续签到天数
	ContinueSignPoint int    // 连续签到赠送数量
	ConsumePerPoint   string // 每消费多少元获取1个点
	LowOrderAmount    string // 最低获取点数的订单金额
	MaxPointPerOrder  int    // 每笔订单最高获取点数
	Type              int    // 类型：0->积分规则；1->成长值规则
}

type RuleSettingRepo interface {
	CreateRuleSetting(context.Context, *RuleSetting) error
	GetRuleSetting(ctx context.Context, id int64) error
	UpdateRuleSetting(context.Context, *RuleSetting) error
	ListRuleSetting(ctx context.Context, req *RuleSettingListReq) ([]*RuleSetting, error)
	DeleteRuleSetting(ctx context.Context, id int64) error
}

type RuleSettingUseCase struct {
	repo RuleSettingRepo
	log  *log.Helper
}

func NewRuleSettingUseCase(repo RuleSettingRepo, logger log.Logger) *RuleSettingUseCase {
	return &RuleSettingUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (r *RuleSettingUseCase) CreateRuleSetting(ctx context.Context, user *RuleSetting) error {
	panic("implement me")
}

func (r *RuleSettingUseCase) GetRuleSetting(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r *RuleSettingUseCase) UpdateRuleSetting(ctx context.Context, user *RuleSetting) error {
	panic("implement me")
}

func (r *RuleSettingUseCase) ListRuleSetting(ctx context.Context, pageNum, pageSize int64) ([]*RuleSetting, error) {
	panic("implement me")
}

func (r *RuleSettingUseCase) DeleteRuleSetting(ctx context.Context, id int64) error {
	panic("implement me")
}
