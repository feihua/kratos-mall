package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/sys/internal/biz"
)

type jobRepo struct {
	data *Data
	log  *log.Helper
}

// NewJobRepo .
func NewJobRepo(data *Data, logger log.Logger) biz.JobRepo {
	return &jobRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (j jobRepo) CreateJob(ctx context.Context, job *biz.Job) error {
	panic("implement me")
}

func (j jobRepo) GetJob(ctx context.Context, id int64) error {
	panic("implement me")
}

func (j jobRepo) UpdateJob(ctx context.Context, job *biz.Job) error {
	panic("implement me")
}

func (j jobRepo) ListJob(ctx context.Context, req *biz.JobListReq) ([]*biz.Job, error) {
	panic("implement me")
}

func (j jobRepo) DeleteJob(ctx context.Context, id int64) error {
	panic("implement me")
}
