package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type LogDTO struct {
	UserName string
	Status   string
	Ip       string
	CreateBy string
}

type LogRepo interface {
	CreateLog(context.Context, *LogDTO) error
	GetLog(ctx context.Context, id int64) error
	UpdateLog(context.Context, *LogDTO) error
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

func (l *LogUseCase) LoginLogAdd(ctx context.Context, user *LogDTO) error {

	return l.repo.CreateLog(ctx, user)

}

func (l *LogUseCase) GetLog(ctx context.Context, id int64) error {
	panic("implement me")
}

func (l *LogUseCase) UpdateLog(ctx context.Context, user *LogDTO) error {
	panic("implement me")
}

func (l *LogUseCase) ListLog(ctx context.Context, pageNum, pageSize int64) ([]*Beer, error) {
	panic("implement me")
}

func (l *LogUseCase) DeleteLog(ctx context.Context, id int64) error {
	panic("implement me")
}
