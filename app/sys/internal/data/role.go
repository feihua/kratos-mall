package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/sys/internal/biz"
	"time"
)

type roleRepo struct {
	data *Data
	log  *log.Helper
}

type Role struct {
	Id          int64
	Rolename    string
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

func (Role) TableName() string {
	return "role"
}

// NewRoleRepo .
func NewRoleRepo(data *Data, logger log.Logger) biz.RoleRepo {
	return &roleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u roleRepo) CreateRole(ctx context.Context, role *biz.Role) error {
	panic("implement me")
}

func (u roleRepo) GetRole(ctx context.Context, id int64) error {
	panic("implement me")
}

func (u roleRepo) UpdateRole(ctx context.Context, role *biz.Role) error {
	panic("implement me")
}

func (u roleRepo) ListRole(ctx context.Context, pageNum, pageSize int64) ([]*biz.Beer, error) {
	panic("implement me")
}

func (u roleRepo) DeleteRole(ctx context.Context, id int64) error {
	panic("implement me")
}
