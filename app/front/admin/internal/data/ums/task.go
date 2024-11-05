package ums

import (
	"context"
	umsV1 "github.com/feihua/kratos-mall/api/ums/v1"
	"github.com/feihua/kratos-mall/app/front/admin/internal/biz/ums"
	"github.com/feihua/kratos-mall/app/front/admin/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
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

func (t taskRepo) ListTask(ctx context.Context, req *ums.TaskListReq) (*ums.TaskListResp, error) {
	list, _ := t.data.UmsClient.MemberTaskList(ctx, &umsV1.MemberTaskListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*ums.Task, 0)
	copier.Copy(&orders, &list.List)

	return &ums.TaskListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (t taskRepo) DeleteTask(ctx context.Context, id int64) error {
	panic("implement me")
}
