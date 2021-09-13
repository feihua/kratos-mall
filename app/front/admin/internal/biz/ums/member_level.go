package ums

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type MemberLevlListReq struct {
	Current  int64
	PageSize int64
}

type MemberLevel struct {
	Id                    int64
	Name                  string
	GrowthPoint           int
	DefaultStatus         int    // 是否为默认等级：0->不是；1->是
	FreeFreightPoint      string // 免运费标准
	CommentGrowthPoint    int    // 每次评价获取的成长值
	PriviledgeFreeFreight int    // 是否有免邮特权
	PriviledgeSignIn      int    // 是否有签到特权
	PriviledgeComment     int    // 是否有评论获奖励特权
	PriviledgePromotion   int    // 是否有专享活动特权
	PriviledgeMemberPrice int    // 是否有会员价格特权
	PriviledgeBirthday    int    // 是否有生日特权
	Note                  string
}
type MemberLevelListResp struct {
	Total int64
	List  []*MemberLevel
}

type MemberLevelRepo interface {
	CreateMemberLevl(context.Context, *MemberLevel) error
	GetMemberLevl(ctx context.Context, id int64) error
	UpdateMemberLevl(context.Context, *MemberLevel) error
	ListMemberLevl(ctx context.Context, req *MemberLevlListReq) (*MemberLevelListResp, error)
	DeleteMemberLevl(ctx context.Context, id int64) error
}

type MemberLevelUseCase struct {
	repo MemberLevelRepo
	log  *log.Helper
}

func NewMemberLevelUseCase(repo MemberLevelRepo, logger log.Logger) *MemberLevelUseCase {
	return &MemberLevelUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (m MemberLevelUseCase) CreateMemberLevl(ctx context.Context, level *MemberLevel) error {
	panic("implement me")
}

func (m MemberLevelUseCase) GetMemberLevl(ctx context.Context, id int64) error {
	panic("implement me")
}

func (m MemberLevelUseCase) UpdateMemberLevl(ctx context.Context, level *MemberLevel) error {
	panic("implement me")
}

func (m MemberLevelUseCase) ListMemberLevl(ctx context.Context, req *MemberLevlListReq) (*MemberLevelListResp, error) {
	return m.repo.ListMemberLevl(ctx, req)
}

func (m MemberLevelUseCase) DeleteMemberLevl(ctx context.Context, id int64) error {
	panic("implement me")
}
