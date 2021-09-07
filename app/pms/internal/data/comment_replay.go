package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/pms/internal/biz"
)

type commentReplayRepo struct {
	data *Data
	log  *log.Helper
}

func NewCommentReplayRepo(data *Data, logger log.Logger) biz.CommentReplayRepo {
	return &commentReplayRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c commentReplayRepo) CreateCommentReplay(ctx context.Context, replay *biz.CommentReplay) error {
	panic("implement me")
}

func (c commentReplayRepo) GetCommentReplay(ctx context.Context, id int64) error {
	panic("implement me")
}

func (c commentReplayRepo) UpdateCommentReplay(ctx context.Context, replay *biz.CommentReplay) error {
	panic("implement me")
}

func (c commentReplayRepo) ListCommentReplay(ctx context.Context, req *biz.CommentReplayListReq) ([]*biz.CommentReplay, error) {
	panic("implement me")
}

func (c commentReplayRepo) DeleteCommentReplay(ctx context.Context, id int64) error {
	panic("implement me")
}
