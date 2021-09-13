package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type TaskListReq struct {
	Current  int64
	PageSize int64
}

type Task struct {
	Id           int64
	Name         string
	Growth       int // 赠送成长值
	Intergration int // 赠送积分
	Type         int // 任务类型：0->新手任务；1->日常任务
}
type TaskListResp struct {
	Total int64
	List  []*Task
}

type TaskRepo interface {
	CreateTask(context.Context, *Task) error
	GetTask(ctx context.Context, id int64) error
	UpdateTask(context.Context, *Task) error
	ListTask(ctx context.Context, req *TaskListReq) (*TaskListResp, error)
	DeleteTask(ctx context.Context, id int64) error
}

type TaskUseCase struct {
	repo TaskRepo
	log  *log.Helper
}

func NewTaskUseCase(repo TaskRepo, logger log.Logger) *TaskUseCase {
	return &TaskUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (t TaskUseCase) CreateTask(ctx context.Context, task *Task) error {
	panic("implement me")
}

func (t TaskUseCase) GetTask(ctx context.Context, id int64) error {
	panic("implement me")
}

func (t TaskUseCase) UpdateTask(ctx context.Context, task *Task) error {
	panic("implement me")
}

func (t TaskUseCase) ListTask(ctx context.Context, req *TaskListReq) (*TaskListResp, error) {
	return t.repo.ListTask(ctx, req)
}

func (t TaskUseCase) DeleteTask(ctx context.Context, id int64) error {
	panic("implement me")
}
