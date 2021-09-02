package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/sys/internal/biz"
	"time"
)

type jobRepo struct {
	data *Data
	log  *log.Helper
}

type Job struct {
	Id          int64
	Jobname     string
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

func (Job) TableName() string {
	return "job"
}

// NewJobRepo .
func NewJobRepo(data *Data, logger log.Logger) biz.JobRepo {
	return &jobRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u jobRepo) CreateJob(ctx context.Context, job *biz.Job) error {
	panic("implement me")
}

func (u jobRepo) GetJob(ctx context.Context, id int64) error {
	panic("implement me")
}

func (u jobRepo) UpdateJob(ctx context.Context, job *biz.Job) error {
	panic("implement me")
}

func (u jobRepo) ListJob(ctx context.Context, pageNum, pageSize int64) ([]*biz.Beer, error) {
	panic("implement me")
}

func (u jobRepo) DeleteJob(ctx context.Context, id int64) error {
	panic("implement me")
}
