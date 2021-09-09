package sys

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/front/admin/internal/biz/sys"
	"kratos-mall/app/front/admin/internal/data"
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

func (d dictRepo) ListDict(ctx context.Context, req *sys.DictListReq) ([]*sys.Dict, error) {
	panic("implement me")
}

func (d dictRepo) DeleteDict(ctx context.Context, id int64) error {
	panic("implement me")
}
