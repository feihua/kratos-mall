package sys

import (
	"context"
	sysV1 "github.com/feihua/kratos-mall/api/sys/v1"
	"github.com/feihua/kratos-mall/app/front/admin/internal/biz/sys"
	"github.com/feihua/kratos-mall/app/front/admin/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
)

type dictRepo struct {
	data *data.Data
	log  *log.Helper
}

// NewDictRepo .
func NewDictRepo(data *data.Data, logger log.Logger) sys.DictRepo {
	return &dictRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (d dictRepo) CreateDict(ctx context.Context, dict *sys.Dict) error {
	panic("implement me")
}

func (d dictRepo) GetDict(ctx context.Context, id int64) error {
	panic("implement me")
}

func (d dictRepo) UpdateDict(ctx context.Context, dict *sys.Dict) error {
	panic("implement me")
}

func (d dictRepo) ListDict(ctx context.Context, req *sys.DictListReq) (*sys.DictListResp, error) {
	list, _ := d.data.SysClient.DictList(ctx, &sysV1.DictListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*sys.Dict, 0)
	copier.Copy(&orders, &list.List)

	return &sys.DictListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (d dictRepo) DeleteDict(ctx context.Context, id int64) error {
	panic("implement me")
}
