package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type CompanyAddressListReq struct {
	Current  int64
	PageSize int64
}

type CompanyAddress struct {
	Id            int64
	AddressName   string // 地址名称
	SendStatus    int    // 默认发货地址：0->否；1->是
	ReceiveStatus int    // 是否默认收货地址：0->否；1->是
	Name          string // 收发货人姓名
	Phone         string // 收货人电话
	Province      string // 省/直辖市
	City          string // 市
	Region        string // 区
	DetailAddress string // 详细地址
}
type CompanyAddressListResp struct {
	Total int64
	List  []*CompanyAddress
}

type CompanyAddressRepo interface {
	CreateCompanyAddress(context.Context, *CompanyAddress) error
	GetCompanyAddress(ctx context.Context, id int64) error
	UpdateCompanyAddress(context.Context, *CompanyAddress) error
	ListCompanyAddress(ctx context.Context, req *CompanyAddressListReq) (*CompanyAddressListResp, error)
	DeleteCompanyAddress(ctx context.Context, id int64) error
}

type CompanyAddressUseCase struct {
	repo CompanyAddressRepo
	log  *log.Helper
}

func NewCompanyAddressUseCase(repo CompanyAddressRepo, logger log.Logger) *CompanyAddressUseCase {
	return &CompanyAddressUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (c CompanyAddressUseCase) CreateCompanyAddress(ctx context.Context, address *CompanyAddress) error {
	panic("implement me")
}

func (c CompanyAddressUseCase) GetCompanyAddress(ctx context.Context, id int64) error {
	panic("implement me")
}

func (c CompanyAddressUseCase) UpdateCompanyAddress(ctx context.Context, address *CompanyAddress) error {
	panic("implement me")
}

func (c CompanyAddressUseCase) ListCompanyAddress(ctx context.Context, req *CompanyAddressListReq) (*CompanyAddressListResp, error) {
	return c.repo.ListCompanyAddress(ctx, req)
}

func (c CompanyAddressUseCase) DeleteCompanyAddress(ctx context.Context, id int64) error {
	panic("implement me")
}
