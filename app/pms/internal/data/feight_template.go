package data

import (
	"context"
	"github.com/feihua/kratos-mall/app/pms/internal/biz"
	"github.com/feihua/kratos-mall/app/pms/internal/data/model"
	"github.com/feihua/kratos-mall/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
)

type feightTemplateRepo struct {
	data *Data
	log  *log.Helper
}

func NewFeightTemplateRepo(data *Data, logger log.Logger) biz.FeightTemplateRepo {
	return &feightTemplateRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (f feightTemplateRepo) CreateFeightTemplate(ctx context.Context, template *biz.FeightTemplate) error {
	panic("implement me")
}

func (f feightTemplateRepo) GetFeightTemplate(ctx context.Context, id int64) error {
	panic("implement me")
}

func (f feightTemplateRepo) UpdateFeightTemplate(ctx context.Context, template *biz.FeightTemplate) error {
	panic("implement me")
}

func (f feightTemplateRepo) ListFeightTemplate(ctx context.Context, req *biz.FeightTemplateListReq) (*biz.FeightTemplateListResp, error) {
	var all []model.PmsFeightTemplate
	result := f.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	f.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.FeightTemplate, 0)

	for _, item := range all {
		list = append(list, &biz.FeightTemplate{
			Id:             item.Id,
			Name:           item.Name,
			ChargeType:     item.ChargeType,
			FirstWeight:    item.FirstWeight,
			FirstFee:       item.FirstFee,
			ContinueWeight: item.ContinueWeight,
			ContinmeFee:    item.ContinmeFee,
			Dest:           item.Dest,
		})
	}

	return &biz.FeightTemplateListResp{
		Total: count,
		List:  list,
	}, nil
}

func (f feightTemplateRepo) DeleteFeightTemplate(ctx context.Context, id int64) error {
	panic("implement me")
}
