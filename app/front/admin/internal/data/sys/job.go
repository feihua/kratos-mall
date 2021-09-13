package sys

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	sysV1 "kratos-mall/api/sys/v1"
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

func (j jobRepo) ListJob(ctx context.Context, req *sys.JobListReq) (*sys.JobListResp, error) {
	list, _ := j.data.SysClient.JobList(ctx, &sysV1.JobListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*sys.Job, 0)
	copier.Copy(&orders, &list.List)

	return &sys.JobListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (j jobRepo) DeleteJob(ctx context.Context, id int64) error {
	panic("implement me")
}
