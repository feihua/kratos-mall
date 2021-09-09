package sys

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/front/admin/internal/biz/sys"
	"kratos-mall/app/front/admin/internal/data"
)

type jobRepo struct {
	data *data.Data
	log  *log.Helper
}

// NewJobRepo .
func NewJobRepo(data *data.Data, logger log.Logger) sys.JobRepo {
	return &jobRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (j jobRepo) CreateJob(ctx context.Context, job *sys.Job) error {
	panic("implement me")
}

func (j jobRepo) GetJob(ctx context.Context, id int64) error {
	panic("implement me")
}

func (j jobRepo) UpdateJob(ctx context.Context, job *sys.Job) error {
	panic("implement me")
}

func (j jobRepo) ListJob(ctx context.Context, req *sys.JobListReq) ([]*sys.Job, error) {
	panic("implement me")
}

func (j jobRepo) DeleteJob(ctx context.Context, id int64) error {
	panic("implement me")
}
