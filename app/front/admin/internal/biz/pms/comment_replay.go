package pms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
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
	CreateTime     time.Time
	Type           int // 评论人员类型；0->会员；1->管理员
}

type CommentReplayRepo interface {
	CreateCommentReplay(context.Context, *CommentReplay) error
	GetCommentReplay(ctx context.Context, id int64) error
	UpdateCommentReplay(context.Context, *CommentReplay) error
	ListCommentReplay(ctx context.Context, req *CommentReplayListReq) ([]*CommentReplay, error)
	DeleteCommentReplay(ctx context.Context, id int64) error
}

type CommentReplayUseCase struct {
	repo CommentReplayRepo
	log  *log.Helper
}

func NewCommentReplayUseCase(repo CommentReplayRepo, logger log.Logger) *CommentReplayUseCase {
	return &CommentReplayUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (r *CommentReplayUseCase) CreateCommentReplay(ctx context.Context, user *CommentReplay) error {
	panic("implement me")
}

func (r *CommentReplayUseCase) GetCommentReplay(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r *CommentReplayUseCase) UpdateCommentReplay(ctx context.Context, user *CommentReplay) error {
	panic("implement me")
}

func (r *CommentReplayUseCase) ListCommentReplay(ctx context.Context, pageNum, pageSize int64) ([]*CommentReplay, error) {
	panic("implement me")
}

func (r *CommentReplayUseCase) DeleteCommentReplay(ctx context.Context, id int64) error {
	panic("implement me")
}
