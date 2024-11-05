package data

import (
	"context"
	"github.com/feihua/kratos-mall/app/ums/internal/biz"
	"github.com/feihua/kratos-mall/app/ums/internal/data/model"
	"github.com/feihua/kratos-mall/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
)

type loginLogRepo struct {
	data *Data
	log  *log.Helper
}

func NewLoginLogRepo(data *Data, logger log.Logger) biz.LoginLogRepo {
	return &loginLogRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (l loginLogRepo) CreateLoginLog(ctx context.Context, loginLog *biz.LoginLog) error {
	panic("implement me")
}

func (l loginLogRepo) GetLoginLog(ctx context.Context, id int64) error {
	panic("implement me")
}

func (l loginLogRepo) UpdateLoginLog(ctx context.Context, loginLog *biz.LoginLog) error {
	panic("implement me")
}

func (l loginLogRepo) ListLoginLog(ctx context.Context, req *biz.LoginLogListReq) (*biz.LoginLogListResp, error) {
	var all []model.UmsMemberLoginLog
	result := l.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	l.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.LoginLog, 0)

	for _, item := range all {
		list = append(list, &biz.LoginLog{
			Id:         item.Id,
			MemberId:   item.MemberId,
			CreateTime: item.CreateTime.Format("2006-01-02 15:04:05"),
			Ip:         item.Ip,
			City:       item.City,
			LoginType:  item.LoginType,
			Province:   item.Province,
		})
	}

	return &biz.LoginLogListResp{
		Total: count,
		List:  list,
	}, nil
}

func (l loginLogRepo) DeleteLoginLog(ctx context.Context, id int64) error {
	panic("implement me")
}
