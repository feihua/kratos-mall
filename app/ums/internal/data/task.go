package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/ums/internal/biz"
)

type taskRepo struct {
	data *Data
	log  *log.Helper
}

func NewTaskRepo(data *Data, logger log.Logger) biz.TaskRepo {
	return &taskRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (t taskRepo) CreateTask(ctx context.Context, task *biz.Task) error {
	panic("implement me")
}

func (t taskRepo) GetTask(ctx context.Context, id int64) error {
	panic("implement me")
}

func (t taskRepo) UpdateTask(ctx context.Context, task *biz.Task) error {
	panic("implement me")
}

func (t taskRepo) ListTask(ctx context.Context, req *biz.TaskListReq) ([]*biz.Task, error) {
	panic("implement me")
}

func (t taskRepo) DeleteTask(ctx context.Context, id int64) error {
	panic("implement me")
}
