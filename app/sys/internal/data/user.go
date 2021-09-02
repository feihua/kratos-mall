package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/sys/internal/biz"
	"time"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

type User struct {
	Id          int64
	Username    string
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

func (User) TableName() string {
	return "user"
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u userRepo) CreateUser(ctx context.Context, user *biz.User) error {
	panic("implement me")
}

func (u userRepo) GetUser(ctx context.Context, id int64) error {
	panic("implement me")
}

func (u userRepo) UpdateUser(ctx context.Context, user *biz.User) error {
	panic("implement me")
}

func (u userRepo) ListUser(ctx context.Context, pageNum, pageSize int64) ([]*biz.Beer, error) {
	panic("implement me")
}

func (u userRepo) DeleteUser(ctx context.Context, id int64) error {
	panic("implement me")
}
