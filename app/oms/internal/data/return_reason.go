package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/oms/internal/biz"
	"kratos-mall/app/oms/internal/data/model"
	"kratos-mall/pkg/util/pagination"
)

type returnReasonRepo struct {
	data *Data
	log  *log.Helper
}

func NewReturnReasonRepo(data *Data, logger log.Logger) biz.ReturnReasonRepo {
	return &returnReasonRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r returnReasonRepo) CreateReturnReason(ctx context.Context, reason *biz.ReturnReason) error {
	panic("implement me")
}

func (r returnReasonRepo) GetReturnReason(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r returnReasonRepo) UpdateReturnReason(ctx context.Context, reason *biz.ReturnReason) error {
	panic("implement me")
}

func (r returnReasonRepo) ListReturnReason(ctx context.Context, req *biz.ReturnReasonListReq) (*biz.ReturnReasonListResp, error) {
	var all []model.OmsOrderReturnReason
	result := r.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	r.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.ReturnReason, 0)

	for _, item := range all {
		list = append(list, &biz.ReturnReason{
			Id:         item.Id,
			Name:       item.Name,
			Sort:       item.Sort,
			Status:     item.Status,
			CreateTime: item.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &biz.ReturnReasonListResp{
		Total: count,
		List:  list,
	}, nil
}

func (r returnReasonRepo) DeleteReturnReason(ctx context.Context, id int64) error {
	panic("implement me")
}
