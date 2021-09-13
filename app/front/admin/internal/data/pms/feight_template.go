package pms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	pmsV1 "kratos-mall/api/pms/v1"
	"kratos-mall/app/front/admin/internal/biz/pms"
	"kratos-mall/app/front/admin/internal/data"
)

type feightTemplateRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewFeightTemplateRepo(data *data.Data, logger log.Logger) pms.FeightTemplateRepo {
	return &feightTemplateRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (f feightTemplateRepo) CreateFeightTemplate(ctx context.Context, template *pms.FeightTemplate) error {
	panic("implement me")
}

func (f feightTemplateRepo) GetFeightTemplate(ctx context.Context, id int64) error {
	panic("implement me")
}

func (f feightTemplateRepo) UpdateFeightTemplate(ctx context.Context, template *pms.FeightTemplate) error {
	panic("implement me")
}

func (f feightTemplateRepo) ListFeightTemplate(ctx context.Context, req *pms.FeightTemplateListReq) (*pms.FeightTemplateListResp, error) {
	list, _ := f.data.PmsClient.FeightTemplateList(ctx, &pmsV1.FeightTemplateListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*pms.FeightTemplate, 0)
	copier.Copy(&orders, &list.List)

	return &pms.FeightTemplateListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (f feightTemplateRepo) DeleteFeightTemplate(ctx context.Context, id int64) error {
	panic("implement me")
}
