package ums

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/front/admin/internal/biz/ums"
	"kratos-mall/app/front/admin/internal/data"
)

type ruleSettingRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewRuleSettingRepo(data *data.Data, logger log.Logger) ums.RuleSettingRepo {
	return &ruleSettingRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r ruleSettingRepo) CreateRuleSetting(ctx context.Context, setting *ums.RuleSetting) error {
	panic("implement me")
}

func (r ruleSettingRepo) GetRuleSetting(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r ruleSettingRepo) UpdateRuleSetting(ctx context.Context, setting *ums.RuleSetting) error {
	panic("implement me")
}

func (r ruleSettingRepo) ListRuleSetting(ctx context.Context, req *ums.RuleSettingListReq) ([]*ums.RuleSetting, error) {
	panic("implement me")
}

func (r ruleSettingRepo) DeleteRuleSetting(ctx context.Context, id int64) error {
	panic("implement me")
}
