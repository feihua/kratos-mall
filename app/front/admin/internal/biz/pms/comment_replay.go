package pms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type CommentReplayListReq struct {
	Current  int64
	PageSize int64
}

type CommentReplay struct {
	Id             int64
	CommentId      int64
	MemberNickName string
	MemberIcon     string
	Content        string
	CreateTime     string
	Type           int // 评论人员类型；0->会员；1->管理员
}
type CommentReplayListResp struct {
	Total int64
	List  []*CommentReplay
}

type CommentReplayRepo interface {
	CreateCommentReplay(context.Context, *CommentReplay) error
	GetCommentReplay(ctx context.Context, id int64) error
	UpdateCommentReplay(context.Context, *CommentReplay) error
	ListCommentReplay(ctx context.Context, req *CommentReplayListReq) (*CommentReplayListResp, error)
	DeleteCommentReplay(ctx context.Context, id int64) error
}

type CommentReplayUseCase struct {
	repo CommentReplayRepo
	log  *log.Helper
}

func NewCommentReplayUseCase(repo CommentReplayRepo, logger log.Logger) *CommentReplayUseCase {
	return &CommentReplayUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (c CommentReplayUseCase) CreateCommentReplay(ctx context.Context, replay *CommentReplay) error {
	panic("implement me")
}

func (c CommentReplayUseCase) GetCommentReplay(ctx context.Context, id int64) error {
	panic("implement me")
}

func (c CommentReplayUseCase) UpdateCommentReplay(ctx context.Context, replay *CommentReplay) error {
	panic("implement me")
}

func (c CommentReplayUseCase) ListCommentReplay(ctx context.Context, req *CommentReplayListReq) (*CommentReplayListResp, error) {
	return c.repo.ListCommentReplay(ctx, req)
}

func (c CommentReplayUseCase) DeleteCommentReplay(ctx context.Context, id int64) error {
	panic("implement me")
}
