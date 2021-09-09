package sys

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/front/admin/internal/biz/sys"
	"kratos-mall/app/front/admin/internal/data"
)

type logRepo struct {
	data *data.Data
	log  *log.Helper
}

// NewLogRepo .
func NewLogRepo(data *data.Data, logger log.Logger) sys.LogRepo {
	return &logRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u logRepo) CreateLog(ctx context.Context, log *sys.LogDTO) error {

	//sysLog := model.SysLoginLog{
	//	UserName: log.UserName,
	//	Status:   log.Status,
	//	Ip:       log.Ip,
	//	CreateBy: log.CreateBy,
	//}
	//
	//u.data.db.Create(&sysLog)

	return nil
}

func (u logRepo) GetLog(ctx context.Context, id int64) error {
	panic("implement me")
}

func (u logRepo) UpdateLog(ctx context.Context, log *sys.LogDTO) error {
	panic("implement me")
}

func (u logRepo) ListLog(ctx context.Context, req *sys.LogListReq) ([]*sys.Log, error) {
	panic("implement me")
}

func (u logRepo) DeleteLog(ctx context.Context, id int64) error {
	panic("implement me")
}
