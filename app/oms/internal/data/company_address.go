package data

import (
	"context"
	"github.com/feihua/kratos-mall/app/oms/internal/biz"
	"github.com/feihua/kratos-mall/app/oms/internal/data/model"
	"github.com/feihua/kratos-mall/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
)

type companyAddressRepo struct {
	data *Data
	log  *log.Helper
}

func NewCompanyAddressRepo(data *Data, logger log.Logger) biz.CompanyAddressRepo {
	return &companyAddressRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c companyAddressRepo) CreateCompanyAddress(ctx context.Context, address *biz.CompanyAddress) error {
	panic("implement me")
}

func (c companyAddressRepo) GetCompanyAddress(ctx context.Context, id int64) error {
	panic("implement me")
}

func (c companyAddressRepo) UpdateCompanyAddress(ctx context.Context, address *biz.CompanyAddress) error {
	panic("implement me")
}

func (c companyAddressRepo) ListCompanyAddress(ctx context.Context, req *biz.CompanyAddressListReq) (*biz.CompanyAddressListResp, error) {
	var all []model.OmsCompanyAddress
	result := c.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	c.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.CompanyAddress, 0)

	for _, item := range all {
		list = append(list, &biz.CompanyAddress{
			Id:            item.Id,
			AddressName:   item.AddressName,
			SendStatus:    item.SendStatus,
			ReceiveStatus: item.ReceiveStatus,
			Name:          item.Name,
			Phone:         item.Phone,
			Province:      item.Province,
			City:          item.City,
			Region:        item.Region,
			DetailAddress: item.DetailAddress,
		})
	}

	return &biz.CompanyAddressListResp{
		Total: count,
		List:  list,
	}, nil
}

func (c companyAddressRepo) DeleteCompanyAddress(ctx context.Context, id int64) error {
	panic("implement me")
}
