package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Log struct {
	Hello string
}

type LogRepo interface {
	CreateLog(context.Context, *Log) error
	GetLog(ctx context.Context, id int64) error
	UpdateLog(context.Context, *Log) error
	ListLog(ctx context.Context, pageNum, pageSize int64) ([]*Beer, error)
	DeleteLog(ctx context.Context, id int64) error
}

type LogUseCase struct {
	repo LogRepo
	log  *log.Helper
}

func NewLogUseCase(repo LogRepo, logger log.Logger) *LogUseCase {
	return &LogUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (l *LogUseCase) CreateLog(ctx context.Context, user *Log) error {
	panic("implement me")
}

func (l *LogUseCase) GetLog(ctx context.Context, id int64) error {
	panic("implement me")
}

func (l *LogUseCase) UpdateLog(ctx context.Context, user *Log) error {
	panic("implement me")
}

func (l *LogUseCase) ListLog(ctx context.Context, pageNum, pageSize int64) ([]*Beer, error) {
	panic("implement me")
}

func (l *LogUseCase) DeleteLog(ctx context.Context, id int64) error {
	panic("implement me")
}
