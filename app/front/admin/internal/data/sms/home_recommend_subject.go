package sms

import (
	"context"
	smsV1 "github.com/feihua/kratos-mall/api/sms/v1"
	"github.com/feihua/kratos-mall/app/front/admin/internal/biz/sms"
	"github.com/feihua/kratos-mall/app/front/admin/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
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

func (h homeRecommendSubjectRepo) ListHomeRecommendSubject(ctx context.Context, req *sms.HomeRecommendSubjectListReq) (*sms.HomeRecommendSubjectListResp, error) {
	list, _ := h.data.SmsClient.HomeRecommendSubjectList(ctx, &smsV1.HomeRecommendSubjectListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*sms.HomeRecommendSubject, 0)
	copier.Copy(&orders, &list.List)

	return &sms.HomeRecommendSubjectListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (h homeRecommendSubjectRepo) DeleteHomeRecommendSubject(ctx context.Context, id int64) error {
	panic("implement me")
}
