package data

import (
	"context"
	"github.com/feihua/kratos-mall/app/sms/internal/biz"
	"github.com/feihua/kratos-mall/app/sms/internal/data/model"
	"github.com/feihua/kratos-mall/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
)

type homeBrandRepo struct {
	data *Data
	log  *log.Helper
}

func NewHomeBrandRepo(data *Data, logger log.Logger) biz.HomeBrandRepo {
	return &homeBrandRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (h homeBrandRepo) CreateHomeBrand(ctx context.Context, brand *biz.HomeBrand) error {
	panic("implement me")
}

func (h homeBrandRepo) GetHomeBrand(ctx context.Context, id int64) error {
	panic("implement me")
}

func (h homeBrandRepo) UpdateHomeBrand(ctx context.Context, brand *biz.HomeBrand) error {
	panic("implement me")
}

func (h homeBrandRepo) ListHomeBrand(ctx context.Context, req *biz.HomeBrandListReq) (*biz.HomeBrandListResp, error) {
	var all []model.SmsHomeBrand
	result := h.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	h.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.HomeBrand, 0)

	for _, item := range all {
		list = append(list, &biz.HomeBrand{
			Id:              item.Id,
			BrandId:         item.BrandId,
			BrandName:       item.BrandName,
			RecommendStatus: item.RecommendStatus,
			Sort:            item.Sort,
		})
	}

	return &biz.HomeBrandListResp{
		Total: count,
		List:  list,
	}, nil
}

func (h homeBrandRepo) DeleteHomeBrand(ctx context.Context, id int64) error {
	panic("implement me")
}
