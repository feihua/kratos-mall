package pms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	pmsV1 "kratos-mall/api/pms/v1"
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

func (c commentReplayRepo) ListCommentReplay(ctx context.Context, req *pms.CommentReplayListReq) (*pms.CommentReplayListResp, error) {
	list, _ := c.data.PmsClient.CommentReplayList(ctx, &pmsV1.CommentReplayListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*pms.CommentReplay, 0)
	copier.Copy(&orders, &list.List)
	return &pms.CommentReplayListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (c commentReplayRepo) DeleteCommentReplay(ctx context.Context, id int64) error {
	panic("implement me")
}
