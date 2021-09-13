package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/ums/internal/biz"
	"kratos-mall/app/ums/internal/data/model"
	"kratos-mall/pkg/util/pagination"
)

type memberLevelRepo struct {
	data *Data
	log  *log.Helper
}

func NewMemberLevelRepo(data *Data, logger log.Logger) biz.MemberLevelRepo {
	return &memberLevelRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (m memberLevelRepo) CreateMemberLevl(ctx context.Context, levl *biz.MemberLevel) error {
	panic("implement me")
}

func (m memberLevelRepo) GetMemberLevl(ctx context.Context, id int64) error {
	panic("implement me")
}

func (m memberLevelRepo) UpdateMemberLevl(ctx context.Context, levl *biz.MemberLevel) error {
	panic("implement me")
}

func (m memberLevelRepo) ListMemberLevl(ctx context.Context, req *biz.MemberLevlListReq) (*biz.MemberLevelListResp, error) {
	var all []model.UmsMemberLevel
	result := m.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	if result.Error != nil {
		return nil, result.Error
	}

	var count int64
	m.data.db.WithContext(ctx).Model(&all).Count(&count)

	list := make([]*biz.MemberLevel, 0)

	for _, item := range all {
		list = append(list, &biz.MemberLevel{
			Id:                    item.Id,
			Name:                  item.Name,
			GrowthPoint:           item.GrowthPoint,
			DefaultStatus:         item.DefaultStatus,
			FreeFreightPoint:      item.FreeFreightPoint,
			CommentGrowthPoint:    item.CommentGrowthPoint,
			PriviledgeFreeFreight: item.PriviledgeFreeFreight,
			PriviledgeSignIn:      item.PriviledgeSignIn,
			PriviledgeComment:     item.PriviledgeComment,
			PriviledgePromotion:   item.PriviledgePromotion,
			PriviledgeMemberPrice: item.PriviledgeMemberPrice,
			PriviledgeBirthday:    item.PriviledgeBirthday,
			Note:                  item.Note,
		})
	}

	return &biz.MemberLevelListResp{
		Total: count,
		List:  list,
	}, nil
}

func (m memberLevelRepo) DeleteMemberLevl(ctx context.Context, id int64) error {
	panic("implement me")
}
