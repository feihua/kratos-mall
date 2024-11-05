package data

import (
	"context"
	"github.com/feihua/kratos-mall/app/sys/internal/biz"
	"github.com/feihua/kratos-mall/app/sys/internal/data/model"
	"github.com/feihua/kratos-mall/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
)

type logRepo struct {
	data *Data
	log  *log.Helper
}

// NewLogRepo .
func NewLogRepo(data *Data, logger log.Logger) biz.LogRepo {
	return &logRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u logRepo) CreateLog(ctx context.Context, log *biz.LogDTO) error {

	sysLog := model.SysLoginLog{
		UserName: log.UserName,
		Status:   log.Status,
		Ip:       log.Ip,
		CreateBy: log.CreateBy,
	}

	u.data.db.Create(&sysLog)

	return nil
}

func (u logRepo) GetLog(ctx context.Context, id int64) error {
	panic("implement me")
}

func (u logRepo) UpdateLog(ctx context.Context, log *biz.LogDTO) error {
	panic("implement me")
}

func (u logRepo) ListLog(ctx context.Context, req *biz.LogListReq) (*biz.LogListResp, error) {
	var all []model.SysLoginLog
	result := u.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	u.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.Log, 0)

	for _, item := range all {
		list = append(list, &biz.Log{
			Id:             item.Id,
			UserName:       item.UserName,
			Status:         item.Status,
			Ip:             item.Ip,
			CreateBy:       item.CreateBy,
			CreateTime:     item.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   item.LastUpdateBy,
			LastUpdateTime: item.LastUpdateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &biz.LogListResp{
		Total: count,
		List:  list,
	}, nil
}

func (u logRepo) DeleteLog(ctx context.Context, id int64) error {
	panic("implement me")
}
