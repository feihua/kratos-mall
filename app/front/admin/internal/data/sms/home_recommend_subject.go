package sms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/front/admin/internal/biz/sms"
	"kratos-mall/app/front/admin/internal/data"
)

type homeRecommendSubjectRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewHomeRecommendSubjectRepo(data *data.Data, logger log.Logger) sms.HomeRecommendSubjectRepo {
	return &homeRecommendSubjectRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (h homeRecommendSubjectRepo) CreateHomeRecommendSubject(ctx context.Context, subject *sms.HomeRecommendSubject) error {
	panic("implement me")
}

func (h homeRecommendSubjectRepo) GetHomeRecommendSubject(ctx context.Context, id int64) error {
	panic("implement me")
}

func (h homeRecommendSubjectRepo) UpdateHomeRecommendSubject(ctx context.Context, subject *sms.HomeRecommendSubject) error {
	panic("implement me")
}

func (h homeRecommendSubjectRepo) ListHomeRecommendSubject(ctx context.Context, req *sms.HomeRecommendSubjectListReq) ([]*sms.HomeRecommendSubject, error) {
	panic("implement me")
}

func (h homeRecommendSubjectRepo) DeleteHomeRecommendSubject(ctx context.Context, id int64) error {
	panic("implement me")
}
