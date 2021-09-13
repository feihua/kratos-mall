package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/ums/internal/biz"
	"kratos-mall/app/ums/internal/data/model"
	"kratos-mall/pkg/util/pagination"
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

func (t taskRepo) ListTask(ctx context.Context, req *biz.TaskListReq) (*biz.TaskListResp, error) {
	var all []model.UmsMemberTask
	result := t.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	t.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.Task, 0)

	for _, item := range all {
		list = append(list, &biz.Task{
			Id:           item.Id,
			Name:         item.Name,
			Growth:       item.Growth,
			Intergration: item.Intergration,
			Type:         item.Type,
		})
	}

	return &biz.TaskListResp{
		Total: count,
		List:  list,
	}, nil
}

func (t taskRepo) DeleteTask(ctx context.Context, id int64) error {
	panic("implement me")
}
