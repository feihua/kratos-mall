package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/pms/internal/biz"
	"kratos-mall/app/pms/internal/data/model"
	"kratos-mall/pkg/util/pagination"
)

type memberPriceRepo struct {
	data *Data
	log  *log.Helper
}

func NewMemberPriceRepo(data *Data, logger log.Logger) biz.MemberPriceRepo {
	return &memberPriceRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (m memberPriceRepo) CreateMemberPrice(ctx context.Context, price *biz.MemberPrice) error {
	panic("implement me")
}

func (m memberPriceRepo) GetMemberPrice(ctx context.Context, id int64) error {
	panic("implement me")
}

func (m memberPriceRepo) UpdateMemberPrice(ctx context.Context, price *biz.MemberPrice) error {
	panic("implement me")
}

func (m memberPriceRepo) ListMemberPrice(ctx context.Context, req *biz.MemberPriceListReq) (*biz.MemberPriceListResp, error) {
	var all []model.PmsMemberPrice
	result := m.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	m.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.MemberPrice, 0)

	for _, item := range all {
		list = append(list, &biz.MemberPrice{
			Id:              item.Id,
			ProductId:       item.ProductId,
			MemberLevelId:   item.MemberLevelId,
			MemberPrice:     item.MemberPrice,
			MemberLevelName: item.MemberLevelName,
		})
	}

	return &biz.MemberPriceListResp{
		Total: count,
		List:  list,
	}, nil
}

func (m memberPriceRepo) DeleteMemberPrice(ctx context.Context, id int64) error {
	panic("implement me")
}
