package ums

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	umsV1 "kratos-mall/api/ums/v1"
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

func (r ruleSettingRepo) ListRuleSetting(ctx context.Context, req *ums.RuleSettingListReq) (*ums.RuleSettingListResp, error) {
	list, _ := r.data.UmsClient.MemberRuleSettingList(ctx, &umsV1.MemberRuleSettingListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*ums.RuleSetting, 0)
	copier.Copy(&orders, &list.List)

	return &ums.RuleSettingListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (r ruleSettingRepo) DeleteRuleSetting(ctx context.Context, id int64) error {
	panic("implement me")
}
