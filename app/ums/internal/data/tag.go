package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/ums/internal/biz"
	"kratos-mall/app/ums/internal/data/model"
	"kratos-mall/pkg/util/pagination"
)

type tagRepo struct {
	data *Data
	log  *log.Helper
}

func NewTagRepo(data *Data, logger log.Logger) biz.TagRepo {
	return &tagRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (t tagRepo) CreateTag(ctx context.Context, tag *biz.Tag) error {
	panic("implement me")
}

func (t tagRepo) GetTag(ctx context.Context, id int64) error {
	panic("implement me")
}

func (t tagRepo) UpdateTag(ctx context.Context, tag *biz.Tag) error {
	panic("implement me")
}

func (t tagRepo) ListTag(ctx context.Context, req *biz.TagListReq) (*biz.TagListResp, error) {
	var all []model.UmsMemberTag
	result := t.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	t.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.Tag, 0)

	for _, item := range all {
		list = append(list, &biz.Tag{
			Id:                item.Id,
			Name:              item.Name,
			FinishOrderCount:  item.FinishOrderCount,
			FinishOrderAmount: item.FinishOrderAmount,
		})
	}

	return &biz.TagListResp{
		Total: count,
		List:  list,
	}, nil
}

func (t tagRepo) DeleteTag(ctx context.Context, id int64) error {
	panic("implement me")
}
