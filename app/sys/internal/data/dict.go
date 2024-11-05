package data

import (
	"context"
	"github.com/feihua/kratos-mall/app/sys/internal/biz"
	"github.com/feihua/kratos-mall/app/sys/internal/data/model"
	"github.com/feihua/kratos-mall/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
)

type dictRepo struct {
	data *Data
	log  *log.Helper
}

// NewDictRepo .
func NewDictRepo(data *Data, logger log.Logger) biz.DictRepo {
	return &dictRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (d dictRepo) CreateDict(ctx context.Context, dict *biz.Dict) error {
	panic("implement me")
}

func (d dictRepo) GetDict(ctx context.Context, id int64) error {
	panic("implement me")
}

func (d dictRepo) UpdateDict(ctx context.Context, dict *biz.Dict) error {
	panic("implement me")
}

func (d dictRepo) ListDict(ctx context.Context, req *biz.DictListReq) (*biz.DictListResp, error) {
	var all []model.SysDict
	result := d.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	d.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.Dict, 0)

	for _, dict := range all {
		list = append(list, &biz.Dict{
			Id:             dict.Id,
			Value:          dict.Value,
			Label:          dict.Label,
			Type:           dict.Type,
			Description:    dict.Description,
			Sort:           dict.Sort,
			Remarks:        dict.Remarks,
			CreateBy:       dict.CreateBy,
			CreateTime:     dict.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   dict.LastUpdateBy,
			LastUpdateTime: dict.LastUpdateTime.Format("2006-01-02 15:04:05"),
			DelFlag:        dict.DelFlag,
		})
	}

	return &biz.DictListResp{
		Total: count,
		List:  list,
	}, nil
}

func (d dictRepo) DeleteDict(ctx context.Context, id int64) error {
	panic("implement me")
}
