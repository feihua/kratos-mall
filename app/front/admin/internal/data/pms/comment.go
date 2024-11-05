package pms

import (
	"context"
	pmsV1 "github.com/feihua/kratos-mall/api/pms/v1"
	"github.com/feihua/kratos-mall/app/front/admin/internal/biz/pms"
	"github.com/feihua/kratos-mall/app/front/admin/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
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

func (c commentRepo) ListComment(ctx context.Context, req *pms.CommentListReq) (*pms.CommentListResp, error) {
	list, _ := c.data.PmsClient.CommentList(ctx, &pmsV1.CommentListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*pms.Comment, 0)
	copier.Copy(&orders, &list.List)
	return &pms.CommentListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (c commentRepo) DeleteComment(ctx context.Context, id int64) error {
	panic("implement me")
}
