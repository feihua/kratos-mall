package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/sms/internal/biz"
	"kratos-mall/app/sms/internal/data/model"
	"kratos-mall/pkg/util/pagination"
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

func (h homeRecommendSubjectRepo) ListHomeRecommendSubject(ctx context.Context, req *biz.HomeRecommendSubjectListReq) (*biz.HomeRecommendSubjectListResp, error) {
	var all []model.SmsHomeRecommendSubject
	result := h.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	h.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.HomeRecommendSubject, 0)

	for _, item := range all {
		list = append(list, &biz.HomeRecommendSubject{
			Id:              item.Id,
			SubjectId:       item.SubjectId,
			SubjectName:     item.SubjectName,
			RecommendStatus: item.RecommendStatus,
			Sort:            item.Sort,
		})
	}

	return &biz.HomeRecommendSubjectListResp{
		Total: count,
		List:  list,
	}, nil
}

func (h homeRecommendSubjectRepo) DeleteHomeRecommendSubject(ctx context.Context, id int64) error {
	panic("implement me")
}
