package data

import (
	"context"
	"github.com/feihua/kratos-mall/app/ums/internal/biz"
	"github.com/feihua/kratos-mall/app/ums/internal/data/model"
	"github.com/feihua/kratos-mall/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
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

func (r ruleSettingRepo) ListRuleSetting(ctx context.Context, req *biz.RuleSettingListReq) (*biz.RuleSettingListResp, error) {
	var all []model.UmsMemberRuleSetting
	result := r.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	r.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.RuleSetting, 0)

	for _, item := range all {
		list = append(list, &biz.RuleSetting{
			Id:                item.Id,
			ContinueSignDay:   item.ContinueSignDay,
			ContinueSignPoint: item.ContinueSignPoint,
			ConsumePerPoint:   item.ConsumePerPoint,
			LowOrderAmount:    item.LowOrderAmount,
			MaxPointPerOrder:  item.MaxPointPerOrder,
			Type:              item.Type,
		})
	}

	return &biz.RuleSettingListResp{
		Total: count,
		List:  list,
	}, nil
}

func (r ruleSettingRepo) DeleteRuleSetting(ctx context.Context, id int64) error {
	panic("implement me")
}
