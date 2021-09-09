package pms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/front/admin/internal/biz/pms"
	"kratos-mall/app/front/admin/internal/data"
)

type commentRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewCommentRepo(data *data.Data, logger log.Logger) pms.CommentRepo {
	return &commentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c commentRepo) CreateComment(ctx context.Context, comment *pms.Comment) error {
	panic("implement me")
}

func (c commentRepo) GetComment(ctx context.Context, id int64) error {
	panic("implement me")
}

func (c commentRepo) UpdateComment(ctx context.Context, comment *pms.Comment) error {
	panic("implement me")
}

func (c commentRepo) ListComment(ctx context.Context, req *pms.CommentListReq) ([]*pms.Comment, error) {
	panic("implement me")
}

func (c commentRepo) DeleteComment(ctx context.Context, id int64) error {
	panic("implement me")
}
