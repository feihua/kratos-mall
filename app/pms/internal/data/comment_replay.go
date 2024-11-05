package data

import (
	"context"
	"github.com/feihua/kratos-mall/app/pms/internal/biz"
	"github.com/feihua/kratos-mall/app/pms/internal/data/model"
	"github.com/feihua/kratos-mall/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
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

func (c commentReplayRepo) ListCommentReplay(ctx context.Context, req *biz.CommentReplayListReq) (*biz.CommentReplayListResp, error) {
	var all []model.PmsCommentReplay
	result := c.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	c.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.CommentReplay, 0)

	for _, item := range all {
		list = append(list, &biz.CommentReplay{
			Id:             item.Id,
			CommentId:      item.CommentId,
			MemberNickName: item.MemberNickName,
			MemberIcon:     item.MemberIcon,
			Content:        item.Content,
			CreateTime:     item.CreateTime.Format("2006-01-02 15:04:05"),
			Type:           item.Type,
		})
	}

	return &biz.CommentReplayListResp{
		Total: count,
		List:  list,
	}, nil
}

func (c commentReplayRepo) DeleteCommentReplay(ctx context.Context, id int64) error {
	panic("implement me")
}
