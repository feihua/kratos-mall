package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/sys/internal/biz"
	"time"
)

type dictRepo struct {
	data *Data
	log  *log.Helper
}

type Dict struct {
	Id          int64
	Dictname    string
	Salt        string
	Password    string
	Mobile      string
	Nickname    string
	Avatar      string
	Status      int
	LastLoginAt *time.Time
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

func (Dict) TableName() string {
	return "dict"
}

// NewDictRepo .
func NewDictRepo(data *Data, logger log.Logger) biz.DictRepo {
	return &dictRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u dictRepo) CreateDict(ctx context.Context, dict *biz.Dict) error {
	panic("implement me")
}

func (u dictRepo) GetDict(ctx context.Context, id int64) error {
	panic("implement me")
}

func (u dictRepo) UpdateDict(ctx context.Context, dict *biz.Dict) error {
	panic("implement me")
}

func (u dictRepo) ListDict(ctx context.Context, pageNum, pageSize int64) ([]*biz.Beer, error) {
	panic("implement me")
}

func (u dictRepo) DeleteDict(ctx context.Context, id int64) error {
	panic("implement me")
}
