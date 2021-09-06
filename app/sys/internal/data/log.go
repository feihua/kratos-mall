package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/sys/internal/biz"
	"kratos-mall/app/sys/internal/data/model"
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

func (u logRepo) ListLog(ctx context.Context, pageNum, pageSize int64) ([]*biz.Beer, error) {
	panic("implement me")
}

func (u logRepo) DeleteLog(ctx context.Context, id int64) error {
	panic("implement me")
}
