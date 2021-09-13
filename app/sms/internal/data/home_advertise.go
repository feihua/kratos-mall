package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/sms/internal/biz"
	"kratos-mall/app/sms/internal/data/model"
	"kratos-mall/pkg/util/pagination"
)

type homeAdvertiseRepo struct {
	data *Data
	log  *log.Helper
}

func NewHomeAdvertiseRepo(data *Data, logger log.Logger) biz.HomeAdvertiseRepo {
	return &homeAdvertiseRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (h homeAdvertiseRepo) CreateHomeAdvertise(ctx context.Context, advertise *biz.HomeAdvertise) error {
	panic("implement me")
}

func (h homeAdvertiseRepo) GetHomeAdvertise(ctx context.Context, id int64) error {
	panic("implement me")
}

func (h homeAdvertiseRepo) UpdateHomeAdvertise(ctx context.Context, advertise *biz.HomeAdvertise) error {
	panic("implement me")
}

func (h homeAdvertiseRepo) ListHomeAdvertise(ctx context.Context, req *biz.HomeAdvertiseListReq) (*biz.HomeAdvertiseListResp, error) {
	var all []model.SmsHomeAdvertise
	result := h.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	h.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.HomeAdvertise, 0)

	for _, item := range all {
		list = append(list, &biz.HomeAdvertise{
			Id:         item.Id,
			Name:       item.Name,
			Type:       item.Type,
			Pic:        item.Pic,
			StartTime:  item.StartTime.Format("2006-01-02 15:04:05"),
			EndTime:    item.EndTime.Format("2006-01-02 15:04:05"),
			Status:     item.Status,
			ClickCount: item.ClickCount,
			OrderCount: item.OrderCount,
			Url:        item.Url,
			Note:       item.Note,
			Sort:       item.Sort,
		})
	}

	return &biz.HomeAdvertiseListResp{
		Total: count,
		List:  list,
	}, nil
}

func (h homeAdvertiseRepo) DeleteHomeAdvertise(ctx context.Context, id int64) error {
	panic("implement me")
}
