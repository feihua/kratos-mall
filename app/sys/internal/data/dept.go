package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/sys/internal/biz"
	"kratos-mall/app/sys/internal/data/model"
)

type deptRepo struct {
	data *Data
	log  *log.Helper
}

// NewDeptRepo .
func NewDeptRepo(data *Data, logger log.Logger) biz.DeptRepo {
	return &deptRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (d deptRepo) CreateDept(ctx context.Context, dept *biz.Dept) error {
	panic("implement me")
}

func (d deptRepo) GetDept(ctx context.Context, id int64) error {
	panic("implement me")
}

func (d deptRepo) UpdateDept(ctx context.Context, dept *biz.Dept) error {
	panic("implement me")
}

func (d deptRepo) ListDept(ctx context.Context, req *biz.DeptListReq) (*biz.DeptListResp, error) {
	var all []model.SysDept
	result := d.data.db.WithContext(ctx).
		Find(&all)

	var count int64
	d.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.Dept, 0)

	for _, dept := range all {
		list = append(list, &biz.Dept{
			Id:             dept.Id,
			Name:           dept.Name,
			ParentId:       dept.ParentId,
			OrderNum:       dept.OrderNum,
			CreateBy:       dept.CreateBy,
			CreateTime:     dept.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   dept.LastUpdateBy,
			LastUpdateTime: dept.LastUpdateTime.Format("2006-01-02 15:04:05"),
			DelFlag:        dept.DelFlag,
		})
	}

	return &biz.DeptListResp{
		Total: count,
		List:  list,
	}, nil
}

func (d deptRepo) DeleteDept(ctx context.Context, id int64) error {
	panic("implement me")
}
