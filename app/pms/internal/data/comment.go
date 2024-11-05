package data

import (
	"context"
	"github.com/feihua/kratos-mall/app/pms/internal/biz"
	"github.com/feihua/kratos-mall/app/pms/internal/data/model"
	"github.com/feihua/kratos-mall/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
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

func (c commentRepo) ListComment(ctx context.Context, req *biz.CommentListReq) (*biz.CommentListResp, error) {
	var all []model.PmsComment
	result := c.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	c.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.Comment, 0)

	for _, item := range all {
		list = append(list, &biz.Comment{
			Id:               item.Id,
			ProductId:        item.ProductId,
			MemberNickName:   item.MemberNickName,
			ProductName:      item.ProductName,
			Star:             item.Star,
			MemberIp:         item.MemberIp,
			CreateTime:       item.CreateTime.Format("2006-01-02 15:04:05"),
			ShowStatus:       item.ShowStatus,
			ProductAttribute: item.ProductAttribute,
			CollectCouont:    item.CollectCouont,
			ReadCount:        item.ReadCount,
			Content:          item.Content,
			Pics:             item.Pics,
			MemberIcon:       item.MemberIcon,
			ReplayCount:      item.ReplayCount,
		})
	}

	return &biz.CommentListResp{
		Total: count,
		List:  list,
	}, nil
}

func (c commentRepo) DeleteComment(ctx context.Context, id int64) error {
	panic("implement me")
}
