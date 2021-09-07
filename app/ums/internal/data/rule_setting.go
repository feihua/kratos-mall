package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/ums/internal/biz"
)

type ruleSettingRepo struct {
	data *Data
	log  *log.Helper
}

func NewRuleSettingRepo(data *Data, logger log.Logger) biz.RuleSettingRepo {
	return &ruleSettingRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r ruleSettingRepo) CreateRuleSetting(ctx context.Context, setting *biz.RuleSetting) error {
	panic("implement me")
}

func (r ruleSettingRepo) GetRuleSetting(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r ruleSettingRepo) UpdateRuleSetting(ctx context.Context, setting *biz.RuleSetting) error {
	panic("implement me")
}

func (r ruleSettingRepo) ListRuleSetting(ctx context.Context, req *biz.RuleSettingListReq) ([]*biz.RuleSetting, error) {
	panic("implement me")
}

func (r ruleSettingRepo) DeleteRuleSetting(ctx context.Context, id int64) error {
	panic("implement me")
}
