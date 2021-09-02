package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/sys/internal/biz"
	"time"
)

type deptRepo struct {
	data *Data
	log  *log.Helper
}

type Dept struct {
	Id          int64
	Deptname    string
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

func (Dept) TableName() string {
	return "dept"
}

// NewDeptRepo .
func NewDeptRepo(data *Data, logger log.Logger) biz.DeptRepo {
	return &deptRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u deptRepo) CreateDept(ctx context.Context, dept *biz.Dept) error {
	panic("implement me")
}

func (u deptRepo) GetDept(ctx context.Context, id int64) error {
	panic("implement me")
}

func (u deptRepo) UpdateDept(ctx context.Context, dept *biz.Dept) error {
	panic("implement me")
}

func (u deptRepo) ListDept(ctx context.Context, pageNum, pageSize int64) ([]*biz.Beer, error) {
	panic("implement me")
}

func (u deptRepo) DeleteDept(ctx context.Context, id int64) error {
	panic("implement me")
}
