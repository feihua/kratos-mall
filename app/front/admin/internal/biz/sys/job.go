package sys

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type JobListReq struct {
	Current  int64
	PageSize int64
	JobName  string
}

type Job struct {
	Id             int64  // 编号
	JobName        string // 职位名称
	OrderNum       int    // 排序
	CreateBy       string // 创建人
	CreateTime     string // 创建时间
	LastUpdateBy   string // 更新人
	LastUpdateTime string // 更新时间
	DelFlag        int    // 是否删除  -1：已删除  0：正常
	Remarks        string // 备注
}
type JobListResp struct {
	Total int64
	List  []*Job
}

type JobRepo interface {
	CreateJob(context.Context, *Job) error
	GetJob(ctx context.Context, id int64) error
	UpdateJob(context.Context, *Job) error
	ListJob(ctx context.Context, req *JobListReq) (*JobListResp, error)
	DeleteJob(ctx context.Context, id int64) error
}

type JobUseCase struct {
	repo JobRepo
	log  *log.Helper
}

func NewJobUseCase(repo JobRepo, logger log.Logger) *JobUseCase {
	return &JobUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (j JobUseCase) CreateJob(ctx context.Context, job *Job) error {
	panic("implement me")
}

func (j JobUseCase) GetJob(ctx context.Context, id int64) error {
	panic("implement me")
}

func (j JobUseCase) UpdateJob(ctx context.Context, job *Job) error {
	panic("implement me")
}

func (j JobUseCase) ListJob(ctx context.Context, req *JobListReq) (*JobListResp, error) {
	return j.repo.ListJob(ctx, req)
}

func (j JobUseCase) DeleteJob(ctx context.Context, id int64) error {
	panic("implement me")
}
