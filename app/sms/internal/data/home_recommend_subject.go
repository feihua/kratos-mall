package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/sms/internal/biz"
)

type homeRecommendSubjectRepo struct {
	data *Data
	log  *log.Helper
}

func NewHomeRecommendSubjectRepo(data *Data, logger log.Logger) biz.HomeRecommendSubjectRepo {
	return &homeRecommendSubjectRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (h homeRecommendSubjectRepo) CreateHomeRecommendSubject(ctx context.Context, subject *biz.HomeRecommendSubject) error {
	panic("implement me")
}

func (h homeRecommendSubjectRepo) GetHomeRecommendSubject(ctx context.Context, id int64) error {
	panic("implement me")
}

func (h homeRecommendSubjectRepo) UpdateHomeRecommendSubject(ctx context.Context, subject *biz.HomeRecommendSubject) error {
	panic("implement me")
}

func (h homeRecommendSubjectRepo) ListHomeRecommendSubject(ctx context.Context, req *biz.HomeRecommendSubjectListReq) ([]*biz.HomeRecommendSubject, error) {
	panic("implement me")
}

func (h homeRecommendSubjectRepo) DeleteHomeRecommendSubject(ctx context.Context, id int64) error {
	panic("implement me")
}
