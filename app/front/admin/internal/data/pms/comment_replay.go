package pms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/front/admin/internal/biz/pms"
	"kratos-mall/app/front/admin/internal/data"
)

type commentReplayRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewCommentReplayRepo(data *data.Data, logger log.Logger) pms.CommentReplayRepo {
	return &commentReplayRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c commentReplayRepo) CreateCommentReplay(ctx context.Context, replay *pms.CommentReplay) error {
	panic("implement me")
}

func (c commentReplayRepo) GetCommentReplay(ctx context.Context, id int64) error {
	panic("implement me")
}

func (c commentReplayRepo) UpdateCommentReplay(ctx context.Context, replay *pms.CommentReplay) error {
	panic("implement me")
}

func (c commentReplayRepo) ListCommentReplay(ctx context.Context, req *pms.CommentReplayListReq) ([]*pms.CommentReplay, error) {
	panic("implement me")
}

func (c commentReplayRepo) DeleteCommentReplay(ctx context.Context, id int64) error {
	panic("implement me")
}
