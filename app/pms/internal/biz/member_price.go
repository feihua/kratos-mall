package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type MemberPriceListReq struct {
	Current  int64
	PageSize int64
}

type MemberPrice struct {
	Id              int64
	ProductId       int64
	MemberLevelId   int64
	MemberPrice     string
	MemberLevelName string
}
type MemberPriceListResp struct {
	Total int64
	List  []*MemberPrice
}

type MemberPriceRepo interface {
	CreateMemberPrice(context.Context, *MemberPrice) error
	GetMemberPrice(ctx context.Context, id int64) error
	UpdateMemberPrice(context.Context, *MemberPrice) error
	ListMemberPrice(ctx context.Context, req *MemberPriceListReq) (*MemberPriceListResp, error)
	DeleteMemberPrice(ctx context.Context, id int64) error
}

type MemberPriceUseCase struct {
	repo MemberPriceRepo
	log  *log.Helper
}

func NewMemberPriceUseCase(repo MemberPriceRepo, logger log.Logger) *MemberPriceUseCase {
	return &MemberPriceUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (m MemberPriceUseCase) CreateMemberPrice(ctx context.Context, price *MemberPrice) error {
	panic("implement me")
}

func (m MemberPriceUseCase) GetMemberPrice(ctx context.Context, id int64) error {
	panic("implement me")
}

func (m MemberPriceUseCase) UpdateMemberPrice(ctx context.Context, price *MemberPrice) error {
	panic("implement me")
}

func (m MemberPriceUseCase) ListMemberPrice(ctx context.Context, req *MemberPriceListReq) (*MemberPriceListResp, error) {
	return m.repo.ListMemberPrice(ctx, req)
}

func (m MemberPriceUseCase) DeleteMemberPrice(ctx context.Context, id int64) error {
	panic("implement me")
}
