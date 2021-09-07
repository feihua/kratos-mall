package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/pms/internal/biz"
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

func (f feightTemplateRepo) ListFeightTemplate(ctx context.Context, req *biz.FeightTemplateListReq) ([]*biz.FeightTemplate, error) {
	panic("implement me")
}

func (f feightTemplateRepo) DeleteFeightTemplate(ctx context.Context, id int64) error {
	panic("implement me")
}
