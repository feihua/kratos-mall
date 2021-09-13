package oms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	omsV1 "kratos-mall/api/oms/v1"
	"kratos-mall/app/front/admin/internal/biz/oms"
	"kratos-mall/app/front/admin/internal/data"
)

type companyAddressRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewCompanyAddressRepo(data *data.Data, logger log.Logger) oms.CompanyAddressRepo {
	return &companyAddressRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c companyAddressRepo) CreateCompanyAddress(ctx context.Context, address *oms.CompanyAddress) error {
	panic("implement me")
}

func (c companyAddressRepo) GetCompanyAddress(ctx context.Context, id int64) error {
	panic("implement me")
}

func (c companyAddressRepo) UpdateCompanyAddress(ctx context.Context, address *oms.CompanyAddress) error {
	panic("implement me")
}

func (c companyAddressRepo) ListCompanyAddress(ctx context.Context, req *oms.CompanyAddressListReq) (*oms.CompanyAddressListResp, error) {
	list, _ := c.data.OmsClient.CompanyAddressList(ctx, &omsV1.CompanyAddressListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*oms.CompanyAddress, 0)
	copier.Copy(&orders, &list.List)
	return &oms.CompanyAddressListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (c companyAddressRepo) DeleteCompanyAddress(ctx context.Context, id int64) error {
	panic("implement me")
}
