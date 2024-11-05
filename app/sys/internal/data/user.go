package data

import (
	"context"
	"github.com/feihua/kratos-mall/app/sys/internal/biz"
	"github.com/feihua/kratos-mall/app/sys/internal/data/model"
	"github.com/feihua/kratos-mall/pkg/util/pagination"
	"github.com/feihua/kratos-mall/pkg/util/time2str"
	"github.com/go-kratos/kratos/v2/log"
	"time"
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
		Name:     user.Name,
		Salt:     user.Salt,
		Password: user.Password,
		Mobile:   user.Mobile,
		NickName: user.NickName,
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
		Name:     user.Name,
		Salt:     user.Salt,
		Password: user.Password,
		Mobile:   user.Mobile,
		NickName: user.NickName,
		Avatar:   user.Avatar,
		Status:   user.Status,
	}
}

func (u userRepo) UpdateUser(ctx context.Context, user *biz.User) error {
	panic("implement me")
}

func (u userRepo) ListUser(ctx context.Context, req *biz.UserListReq) ([]*biz.User, error) {
	var list []model.SysUser
	result := u.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&list)

	var count int64
	u.data.db.WithContext(ctx).Model(&list).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	rv := make([]*biz.User, 0)
	for _, item := range list {
		rv = append(rv, &biz.User{
			Id:             item.Id,
			Name:           item.Name,
			NickName:       item.NickName,
			Avatar:         item.Avatar,
			Password:       item.Password,
			Salt:           item.Salt,
			Email:          item.Email,
			Mobile:         item.Mobile,
			Status:         item.Status,
			DeptId:         item.DeptId,
			CreateBy:       item.CreateBy,
			CreateTime:     time2str.String(item.CreateTime),
			LastUpdateBy:   item.LastUpdateBy,
			LastUpdateTime: time2str.String(item.LastUpdateTime),
			DelFlag:        item.DelFlag,
			JobId:          item.JobId,
		})
	}
	return rv, nil
}

func (u userRepo) DeleteUser(ctx context.Context, id int64) error {
	panic("implement me")
}

func (u userRepo) DeleteUserRole(ctx context.Context, userId int64) error {
	return u.data.db.Where(&model.SysUserRole{UserId: userId}).Delete(&model.SysUserRole{}).Error
}

func (u userRepo) UpdateUserRole(ctx context.Context, userId int64, roleId int64) error {

	u.data.db.Create(&model.SysUserRole{
		UserId:         userId,
		RoleId:         roleId,
		CreateBy:       "admin",
		CreateTime:     time.Now(),
		LastUpdateBy:   "admin",
		LastUpdateTime: time.Now(),
	})

	return nil
}
