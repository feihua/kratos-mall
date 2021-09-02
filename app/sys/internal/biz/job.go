package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Job struct {
	Hello string
}

type JobRepo interface {
	CreateJob(context.Context, *Job) error
	GetJob(ctx context.Context, id int64) error
	UpdateJob(context.Context, *Job) error
	ListJob(ctx context.Context, pageNum, pageSize int64) ([]*Beer, error)
	DeleteJob(ctx context.Context, id int64) error
}

type JobUseCase struct {
	repo JobRepo
	log  *log.Helper
}

func NewJobUseCase(repo JobRepo, logger log.Logger) *JobUseCase {
	return &JobUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (j *JobUseCase) CreateJob(ctx context.Context, user *Job) error {
	panic("implement me")
}

func (j *JobUseCase) GetJob(ctx context.Context, id int64) error {
	panic("implement me")
}

func (j *JobUseCase) UpdateJob(ctx context.Context, user *Job) error {
	panic("implement me")
}

func (j *JobUseCase) ListJob(ctx context.Context, pageNum, pageSize int64) ([]*Beer, error) {
	panic("implement me")
}

func (j *JobUseCase) DeleteJob(ctx context.Context, id int64) error {
	panic("implement me")
}
