package ums

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type MemberAddressListReq struct {
	Current  int64
	PageSize int64
}

type MemberAddress struct {
	Id            int64
	MemberId      int64
	Name          string // 收货人名称
	PhoneNumber   string
	DefaultStatus int    // 是否为默认
	PostCode      string // 邮政编码
	Province      string // 省份/直辖市
	City          string // 城市
	Region        string // 区
	DetailAddress string // 详细地址(街道)
}

type MemberAddressRepo interface {
	CreateMemberAddress(context.Context, *MemberAddress) error
	GetMemberAddress(ctx context.Context, id int64) error
	UpdateMemberAddress(context.Context, *MemberAddress) error
	ListMemberAddress(ctx context.Context, req *MemberAddressListReq) ([]*MemberAddress, error)
	DeleteMemberAddress(ctx context.Context, id int64) error
}

type MemberAddressUseCase struct {
	repo MemberAddressRepo
	log  *log.Helper
}

func NewMemberAddressUseCase(repo MemberAddressRepo, logger log.Logger) *MemberAddressUseCase {
	return &MemberAddressUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (r *MemberAddressUseCase) CreateMemberAddress(ctx context.Context, user *MemberAddress) error {
	panic("implement me")
}

func (r *MemberAddressUseCase) GetMemberAddress(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r *MemberAddressUseCase) UpdateMemberAddress(ctx context.Context, user *MemberAddress) error {
	panic("implement me")
}

func (r *MemberAddressUseCase) ListMemberAddress(ctx context.Context, pageNum, pageSize int64) ([]*MemberAddress, error) {
	panic("implement me")
}

func (r *MemberAddressUseCase) DeleteMemberAddress(ctx context.Context, id int64) error {
	panic("implement me")
}
