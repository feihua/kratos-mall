package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/sys/internal/biz"
	"kratos-mall/app/sys/internal/data/model"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u userRepo) QueryUserByName(ctx context.Context, name string) *biz.User {

	var user model.SysUser
	_ = u.data.db.WithContext(ctx).Where("name=?", name).First(&user)

	return &biz.User{
		Id:       user.Id,
		Username: user.Name,
		Salt:     user.Salt,
		Password: user.Password,
		Mobile:   user.Mobile,
		Nickname: user.NickName,
		Avatar:   user.Avatar,
		Status:   user.Status,
	}
}

func (u userRepo) CreateUser(ctx context.Context, user *biz.UserDTO) error {
	panic("implement me")
}

func (u userRepo) GetUser(ctx context.Context, id int64) *biz.User {

	var user model.SysUser
	_ = u.data.db.WithContext(ctx).Where("id=?", id).First(&user)

	return &biz.User{
		Id:       user.Id,
		Username: user.Name,
		Salt:     user.Salt,
		Password: user.Password,
		Mobile:   user.Mobile,
		Nickname: user.NickName,
		Avatar:   user.Avatar,
		Status:   user.Status,
	}
}

func (u userRepo) UpdateUser(ctx context.Context, user *biz.UserDTO) error {
	panic("implement me")
}

func (u userRepo) ListUser(ctx context.Context, pageNum, pageSize int64) ([]*biz.Beer, error) {
	panic("implement me")
}

func (u userRepo) DeleteUser(ctx context.Context, id int64) error {
	panic("implement me")
}
