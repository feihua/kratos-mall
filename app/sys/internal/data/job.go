package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/sys/internal/biz"
	"kratos-mall/app/sys/internal/data/model"
	"kratos-mall/pkg/util/pagination"
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

func (j jobRepo) ListJob(ctx context.Context, req *biz.JobListReq) (*biz.JobListResp, error) {
	var all []model.SysJob
	result := j.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	j.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.Job, 0)

	for _, job := range all {
		list = append(list, &biz.Job{
			Id:             job.Id,
			JobName:        job.JobName,
			OrderNum:       job.OrderNum,
			CreateBy:       job.CreateBy,
			CreateTime:     job.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   job.LastUpdateBy,
			LastUpdateTime: job.LastUpdateTime.Format("2006-01-02 15:04:05"),
			DelFlag:        job.DelFlag,
			Remarks:        job.Remarks,
		})
	}

	return &biz.JobListResp{
		Total: count,
		List:  list,
	}, nil
}

func (j jobRepo) DeleteJob(ctx context.Context, id int64) error {
	panic("implement me")
}
