package ums

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/front/admin/internal/biz/ums"
	"kratos-mall/app/front/admin/internal/data"
)

type taskRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewTaskRepo(data *data.Data, logger log.Logger) ums.TaskRepo {
	return &taskRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (t taskRepo) CreateTask(ctx context.Context, task *ums.Task) error {
	panic("implement me")
}

func (t taskRepo) GetTask(ctx context.Context, id int64) error {
	panic("implement me")
}

func (t taskRepo) UpdateTask(ctx context.Context, task *ums.Task) error {
	panic("implement me")
}

func (t taskRepo) ListTask(ctx context.Context, req *ums.TaskListReq) ([]*ums.Task, error) {
	panic("implement me")
}

func (t taskRepo) DeleteTask(ctx context.Context, id int64) error {
	panic("implement me")
}
