package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/oms/internal/biz"
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

func (c companyAddressRepo) ListCompanyAddress(ctx context.Context, req *biz.CompanyAddressListReq) ([]*biz.CompanyAddress, error) {
	panic("implement me")
}

func (c companyAddressRepo) DeleteCompanyAddress(ctx context.Context, id int64) error {
	panic("implement me")
}
