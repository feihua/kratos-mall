package biz

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

type MemberLevelRepo interface {
	CreateMemberLevl(context.Context, *MemberLevel) error
	GetMemberLevl(ctx context.Context, id int64) error
	UpdateMemberLevl(context.Context, *MemberLevel) error
	ListMemberLevl(ctx context.Context, req *MemberLevlListReq) ([]*MemberLevel, error)
	DeleteMemberLevl(ctx context.Context, id int64) error
}

type MemberLevelUseCase struct {
	repo MemberLevelRepo
	log  *log.Helper
}

func NewMemberLevelUseCase(repo MemberLevelRepo, logger log.Logger) *MemberLevelUseCase {
	return &MemberLevelUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (r *MemberLevelUseCase) CreateMemberLevl(ctx context.Context, user *MemberLevel) error {
	panic("implement me")
}

func (r *MemberLevelUseCase) GetMemberLevl(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r *MemberLevelUseCase) UpdateMemberLevl(ctx context.Context, user *MemberLevel) error {
	panic("implement me")
}

func (r *MemberLevelUseCase) ListMemberLevl(ctx context.Context, pageNum, pageSize int64) ([]*MemberLevel, error) {
	panic("implement me")
}

func (r *MemberLevelUseCase) DeleteMemberLevl(ctx context.Context, id int64) error {
	panic("implement me")
}
