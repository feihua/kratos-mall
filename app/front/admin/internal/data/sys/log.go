package sys

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	sysV1 "kratos-mall/api/sys/v1"
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

func (u logRepo) ListLog(ctx context.Context, req *sys.LogListReq) (*sys.LogListResp, error) {
	list, _ := u.data.SysClient.LoginLogList(ctx, &sysV1.LoginLogListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*sys.Log, 0)
	copier.Copy(&orders, &list.List)

	return &sys.LogListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (u logRepo) DeleteLog(ctx context.Context, id int64) error {
	panic("implement me")
}
