package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/sys/internal/biz"
	"time"
)

type logRepo struct {
	data *Data
	log  *log.Helper
}

type Log struct {
	Id          int64
	Logname     string
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

func (Log) TableName() string {
	return "log"
}

// NewLogRepo .
func NewLogRepo(data *Data, logger log.Logger) biz.LogRepo {
	return &logRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u logRepo) CreateLog(ctx context.Context, log *biz.Log) error {
	panic("implement me")
}

func (u logRepo) GetLog(ctx context.Context, id int64) error {
	panic("implement me")
}

func (u logRepo) UpdateLog(ctx context.Context, log *biz.Log) error {
	panic("implement me")
}

func (u logRepo) ListLog(ctx context.Context, pageNum, pageSize int64) ([]*biz.Beer, error) {
	panic("implement me")
}

func (u logRepo) DeleteLog(ctx context.Context, id int64) error {
	panic("implement me")
}
