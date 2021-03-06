package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type MemberListReq struct {
	Current  int64
	PageSize int64
}

type Member struct {
	Id                    int64
	MemberLevelId         int64
	Username              string // 用户名
	Password              string // 密码
	Nickname              string // 昵称
	Phone                 string // 手机号码
	Status                int    // 帐号启用状态:0->禁用；1->启用
	CreateTime            string // 注册时间
	Icon                  string // 头像
	Gender                int    // 性别：0->未知；1->男；2->女
	Birthday              string // 生日
	City                  string // 所做城市
	Job                   string // 职业
	PersonalizedSignature string // 个性签名
	SourceType            int    // 用户来源
	Integration           int    // 积分
	Growth                int    // 成长值
	LuckeyCount           int    // 剩余抽奖次数
	HistoryIntegration    int    // 历史积分数量
}
type MemberListResp struct {
	Total int64
	List  []*Member
}

type MemberRepo interface {
	CreateMember(context.Context, *Member) error
	GetMember(ctx context.Context, id int64) error
	UpdateMember(context.Context, *Member) error
	ListMember(ctx context.Context, req *MemberListReq) (*MemberListResp, error)
	DeleteMember(ctx context.Context, id int64) error
}

type MemberUseCase struct {
	repo MemberRepo
	log  *log.Helper
}

func NewMemberUseCase(repo MemberRepo, logger log.Logger) *MemberUseCase {
	return &MemberUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (m MemberUseCase) CreateMember(ctx context.Context, member *Member) error {
	panic("implement me")
}

func (m MemberUseCase) GetMember(ctx context.Context, id int64) error {
	panic("implement me")
}

func (m MemberUseCase) UpdateMember(ctx context.Context, member *Member) error {
	panic("implement me")
}

func (m MemberUseCase) ListMember(ctx context.Context, req *MemberListReq) (*MemberListResp, error) {
	return m.repo.ListMember(ctx, req)
}

func (m MemberUseCase) DeleteMember(ctx context.Context, id int64) error {
	panic("implement me")
}
