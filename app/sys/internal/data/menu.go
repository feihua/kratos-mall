package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/sys/internal/biz"
	"time"
)

type menuRepo struct {
	data *Data
	log  *log.Helper
}

type Menu struct {
	Id          int64
	Menuname    string
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

func (Menu) TableName() string {
	return "menu"
}

// NewMenuRepo .
func NewMenuRepo(data *Data, logger log.Logger) biz.MenuRepo {
	return &menuRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u menuRepo) CreateMenu(ctx context.Context, menu *biz.Menu) error {
	panic("implement me")
}

func (u menuRepo) GetMenu(ctx context.Context, id int64) error {
	panic("implement me")
}

func (u menuRepo) UpdateMenu(ctx context.Context, menu *biz.Menu) error {
	panic("implement me")
}

func (u menuRepo) ListMenu(ctx context.Context, pageNum, pageSize int64) ([]*biz.Beer, error) {
	panic("implement me")
}

func (u menuRepo) DeleteMenu(ctx context.Context, id int64) error {
	panic("implement me")
}
