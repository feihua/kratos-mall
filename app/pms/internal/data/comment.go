package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/pms/internal/biz"
)

type commentRepo struct {
	data *Data
	log  *log.Helper
}

func NewCommentRepo(data *Data, logger log.Logger) biz.CommentRepo {
	return &commentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c commentRepo) CreateComment(ctx context.Context, comment *biz.Comment) error {
	panic("implement me")
}

func (c commentRepo) GetComment(ctx context.Context, id int64) error {
	panic("implement me")
}

func (c commentRepo) UpdateComment(ctx context.Context, comment *biz.Comment) error {
	panic("implement me")
}

func (c commentRepo) ListComment(ctx context.Context, req *biz.CommentListReq) ([]*biz.Comment, error) {
	panic("implement me")
}

func (c commentRepo) DeleteComment(ctx context.Context, id int64) error {
	panic("implement me")
}
