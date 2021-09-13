package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/pms/internal/biz"
	"kratos-mall/app/pms/internal/data/model"
	"kratos-mall/pkg/util/pagination"
)

type operateHistoryRepo struct {
	data *Data
	log  *log.Helper
}

func NewOperateHistoryRepo(data *Data, logger log.Logger) biz.OperateHistoryRepo {
	return &operateHistoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (o operateHistoryRepo) CreateOperateHistory(ctx context.Context, history *biz.OperateHistory) error {
	panic("implement me")
}

func (o operateHistoryRepo) GetOperateHistory(ctx context.Context, id int64) error {
	panic("implement me")
}

func (o operateHistoryRepo) UpdateOperateHistory(ctx context.Context, history *biz.OperateHistory) error {
	panic("implement me")
}

func (o operateHistoryRepo) ListOperateHistory(ctx context.Context, req *biz.OperateHistoryListReq) (*biz.OperateHistoryListResp, error) {
	var all []model.PmsProductOperateLog
	result := o.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	o.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.OperateHistory, 0)

	for _, item := range all {
		list = append(list, &biz.OperateHistory{
			Id:               item.Id,
			ProductId:        item.ProductId,
			PriceOld:         item.PriceOld,
			PriceNew:         item.PriceNew,
			SalePriceOld:     item.SalePriceOld,
			SalePriceNew:     item.SalePriceNew,
			GiftPointOld:     item.GiftPointOld,
			GiftPointNew:     item.GiftPointNew,
			UsePointLimitOld: item.UsePointLimitOld,
			UsePointLimitNew: item.UsePointLimitNew,
			OperateMan:       item.OperateMan,
			CreateTime:       item.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &biz.OperateHistoryListResp{
		Total: count,
		List:  list,
	}, nil
}

func (o operateHistoryRepo) DeleteOperateHistory(ctx context.Context, id int64) error {
	panic("implement me")
}
