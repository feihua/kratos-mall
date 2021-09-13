package pay

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type MerchantListReq struct {
	Current  int64
	PageSize int64
}

type Merchant struct {
	Id         int
	MerId      string // 本地系统商户Id,分配给调用方
	AppId      string // 应用ID 微信开放平台审核通过的应用APPID
	MchId      string // 微信支付分配的商户号
	MchKey     string // 微信支付分配的商户秘钥
	PayType    string // APP:APP支付 JSAPI:小程序,公众号 MWEB:H5支付
	NotifyUrl  string // 微信支付回调地址
	Remarks    string // 备注
	CreateTime string
	UpdateTime string
}

type MerchantListResp struct {
	Total int64
	List  []*Merchant
}

type MerchantRepo interface {
	CreateMerchant(context.Context, *Merchant) error
	GetMerchant(ctx context.Context, id int64) error
	UpdateMerchant(context.Context, *Merchant) error
	ListMerchant(ctx context.Context, req *MerchantListReq) (*MerchantListResp, error)
	DeleteMerchant(ctx context.Context, id int64) error
}

type MerchantUseCase struct {
	repo MerchantRepo
	log  *log.Helper
}

func NewMerchantUseCase(repo MerchantRepo, logger log.Logger) *MerchantUseCase {
	return &MerchantUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (m MerchantUseCase) CreateMerchant(ctx context.Context, merchant *Merchant) error {
	panic("implement me")
}

func (m MerchantUseCase) GetMerchant(ctx context.Context, id int64) error {
	panic("implement me")
}

func (m MerchantUseCase) UpdateMerchant(ctx context.Context, merchant *Merchant) error {
	panic("implement me")
}

func (m MerchantUseCase) ListMerchant(ctx context.Context, req *MerchantListReq) (*MerchantListResp, error) {
	panic("implement me")
}

func (m MerchantUseCase) DeleteMerchant(ctx context.Context, id int64) error {
	panic("implement me")
}
