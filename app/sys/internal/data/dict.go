package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/sys/internal/biz"
)

type dictRepo struct {
	data *Data
	log  *log.Helper
}

// NewDictRepo .
func NewDictRepo(data *Data, logger log.Logger) biz.DictRepo {
	return &dictRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (d dictRepo) CreateDict(ctx context.Context, dict *biz.Dict) error {
	panic("implement me")
}

func (d dictRepo) GetDict(ctx context.Context, id int64) error {
	panic("implement me")
}

func (d dictRepo) UpdateDict(ctx context.Context, dict *biz.Dict) error {
	panic("implement me")
}

func (d dictRepo) ListDict(ctx context.Context, req *biz.DictListReq) ([]*biz.Dict, error) {
	panic("implement me")
}

func (d dictRepo) DeleteDict(ctx context.Context, id int64) error {
	panic("implement me")
}
