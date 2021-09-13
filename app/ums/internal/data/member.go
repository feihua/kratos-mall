package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/ums/internal/biz"
	"kratos-mall/app/ums/internal/data/model"
	"kratos-mall/pkg/util/pagination"
)

type memberRepo struct {
	data *Data
	log  *log.Helper
}

func NewMemberRepo(data *Data, logger log.Logger) biz.MemberRepo {
	return &memberRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (m memberRepo) CreateMember(ctx context.Context, member *biz.Member) error {
	panic("implement me")
}

func (m memberRepo) GetMember(ctx context.Context, id int64) error {
	panic("implement me")
}

func (m memberRepo) UpdateMember(ctx context.Context, member *biz.Member) error {
	panic("implement me")
}

func (m memberRepo) ListMember(ctx context.Context, req *biz.MemberListReq) (*biz.MemberListResp, error) {
	var all []model.UmsMember
	result := m.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	m.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.Member, 0)

	for _, member := range all {
		list = append(list, &biz.Member{
			Id:                    member.Id,
			MemberLevelId:         member.MemberLevelId,
			Username:              member.Username,
			Password:              member.Password,
			Nickname:              member.Nickname,
			Phone:                 member.Phone,
			Status:                member.Status,
			CreateTime:            member.CreateTime.Format("2006-01-02 15:04:05"),
			Icon:                  member.Icon,
			Gender:                member.Gender,
			Birthday:              member.Birthday.Format("2006-01-02 15:04:05"),
			City:                  member.City,
			Job:                   member.Job,
			PersonalizedSignature: member.PersonalizedSignature,
			SourceType:            member.SourceType,
			Integration:           member.Integration,
			Growth:                member.Growth,
			LuckeyCount:           member.LuckeyCount,
			HistoryIntegration:    member.HistoryIntegration,
		})
	}

	return &biz.MemberListResp{
		Total: count,
		List:  list,
	}, nil
}

func (m memberRepo) DeleteMember(ctx context.Context, id int64) error {
	panic("implement me")
}
