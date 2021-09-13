package ums

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	umsV1 "kratos-mall/api/ums/v1"
	"kratos-mall/app/front/admin/internal/biz/ums"
	"kratos-mall/app/front/admin/internal/data"
)

type tagRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewTagRepo(data *data.Data, logger log.Logger) ums.TagRepo {
	return &tagRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (t tagRepo) CreateTag(ctx context.Context, tag *ums.Tag) error {
	panic("implement me")
}

func (t tagRepo) GetTag(ctx context.Context, id int64) error {
	panic("implement me")
}

func (t tagRepo) UpdateTag(ctx context.Context, tag *ums.Tag) error {
	panic("implement me")
}

func (t tagRepo) ListTag(ctx context.Context, req *ums.TagListReq) (*ums.TagListResp, error) {
	list, _ := t.data.UmsClient.MemberTagList(ctx, &umsV1.MemberTagListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*ums.Tag, 0)
	copier.Copy(&orders, &list.List)

	return &ums.TagListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (t tagRepo) DeleteTag(ctx context.Context, id int64) error {
	panic("implement me")
}
