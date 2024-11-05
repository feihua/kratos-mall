package data

import (
	"context"
	"fmt"
	"github.com/feihua/kratos-mall/app/ums/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

type Users struct {
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

func (Users) TableName() string {
	return "users"
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) CreateGreeter(ctx context.Context, g *biz.Greeter) error {
	var sysUser []Users

	_ = r.data.db.Find(&sysUser)

	fmt.Printf("%v", sysUser)
	return nil
}

func (r *greeterRepo) UpdateGreeter(ctx context.Context, g *biz.Greeter) error {
	return nil
}
