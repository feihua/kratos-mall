package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type CommentListReq struct {
	Current  int64
	PageSize int64
}

type Comment struct {
	Id               int64
	ProductId        int64
	MemberNickName   string
	ProductName      string
	Star             int    // 评价星数：0->5
	MemberIp         string // 评价的ip
	CreateTime       string
	ShowStatus       int
	ProductAttribute string // 购买时的商品属性
	CollectCouont    int
	ReadCount        int
	Content          string
	Pics             string // 上传图片地址，以逗号隔开
	MemberIcon       string // 评论用户头像
	ReplayCount      int
}

type CommentRepo interface {
	CreateComment(context.Context, *Comment) error
	GetComment(ctx context.Context, id int64) error
	UpdateComment(context.Context, *Comment) error
	ListComment(ctx context.Context, req *CommentListReq) ([]*Comment, error)
	DeleteComment(ctx context.Context, id int64) error
}

type CommentUseCase struct {
	repo CommentRepo
	log  *log.Helper
}

func NewCommentUseCase(repo CommentRepo, logger log.Logger) *CommentUseCase {
	return &CommentUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (r *CommentUseCase) CreateComment(ctx context.Context, user *Comment) error {
	panic("implement me")
}

func (r *CommentUseCase) GetComment(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r *CommentUseCase) UpdateComment(ctx context.Context, user *Comment) error {
	panic("implement me")
}

func (r *CommentUseCase) ListComment(ctx context.Context, pageNum, pageSize int64) ([]*Comment, error) {
	panic("implement me")
}

func (r *CommentUseCase) DeleteComment(ctx context.Context, id int64) error {
	panic("implement me")
}
