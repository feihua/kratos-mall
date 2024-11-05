package data

import (
	"context"
	"github.com/feihua/kratos-mall/app/ums/internal/biz"
	"github.com/feihua/kratos-mall/app/ums/internal/data/model"
	"github.com/feihua/kratos-mall/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
)

type memberAddressRepo struct {
	data *Data
	log  *log.Helper
}

func NewMemberAddressRepo(data *Data, logger log.Logger) biz.MemberAddressRepo {
	return &memberAddressRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (m memberAddressRepo) CreateMemberAddress(ctx context.Context, address *biz.MemberAddress) error {
	panic("implement me")
}

func (m memberAddressRepo) GetMemberAddress(ctx context.Context, id int64) error {
	panic("implement me")
}

func (m memberAddressRepo) UpdateMemberAddress(ctx context.Context, address *biz.MemberAddress) error {
	panic("implement me")
}

func (m memberAddressRepo) ListMemberAddress(ctx context.Context, req *biz.MemberAddressListReq) (*biz.MemberAddressListResp, error) {
	var all []model.UmsMemberReceiveAddress
	result := m.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	m.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.MemberAddress, 0)

	for _, item := range all {
		list = append(list, &biz.MemberAddress{
			Id:            item.Id,
			MemberId:      item.MemberId,
			Name:          item.Name,
			PhoneNumber:   item.PhoneNumber,
			DefaultStatus: item.DefaultStatus,
			PostCode:      item.PostCode,
			Province:      item.Province,
			City:          item.City,
			Region:        item.Region,
			DetailAddress: item.DetailAddress,
		})
	}

	return &biz.MemberAddressListResp{
		Total: count,
		List:  list,
	}, nil
}

func (m memberAddressRepo) DeleteMemberAddress(ctx context.Context, id int64) error {
	panic("implement me")
}
