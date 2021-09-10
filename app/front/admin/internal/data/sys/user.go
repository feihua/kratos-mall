package sys

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/front/admin/internal/biz/sys"
	"kratos-mall/app/front/admin/internal/data"
)

type userRepo struct {
	data *data.Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *data.Data, logger log.Logger) sys.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u userRepo) QueryUserByName(ctx context.Context, name string) *sys.User {

	//var user model.SysUser
	//_ = u.data.db.WithContext(ctx).Where("name=?", name).First(&user)
	//
	//return &sys.User{
	//	Id:       user.Id,
	//	Name:     user.Name,
	//	Salt:     user.Salt,
	//	Password: user.Password,
	//	Mobile:   user.Mobile,
	//	NickName: user.NickName,
	//	Avatar:   user.Avatar,
	//	Status:   user.Status,
	//}
	return nil
}

func (u userRepo) CreateUser(ctx context.Context, user *sys.UserDTO) error {
	panic("implement me")
}

func (u userRepo) GetUser(ctx context.Context, id int64) *sys.User {

	//var user model.SysUser
	//_ = u.data.db.WithContext(ctx).Where("id=?", id).First(&user)
	//
	//return &sys.User{
	//	Id:       user.Id,
	//	Name:     user.Name,
	//	Salt:     user.Salt,
	//	Password: user.Password,
	//	Mobile:   user.Mobile,
	//	NickName: user.NickName,
	//	Avatar:   user.Avatar,
	//	Status:   user.Status,
	//}
	return nil
}

func (u userRepo) UpdateUser(ctx context.Context, user *sys.UserDTO) error {
	panic("implement me")
}

func (u userRepo) ListUser(ctx context.Context, req *sys.UserListReq) ([]*sys.User, error) {
	//var list []model.SysUser
	//result := u.data.db.WithContext(ctx).
	//	Limit(int(req.PageSize)).
	//	Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
	//	Find(&list)
	//
	//if result.Error != nil {
	//	return nil, result.Error
	//}
	//
	//rv := make([]*sys.User, 0)
	//for _, item := range list {
	//	rv = append(rv, &sys.User{
	//		Id:             item.Id,
	//		Name:           item.Name,
	//		NickName:       item.NickName,
	//		Avatar:         item.Avatar,
	//		Password:       item.Password,
	//		Salt:           item.Salt,
	//		Email:          item.Email,
	//		Mobile:         item.Mobile,
	//		Status:         item.Status,
	//		DeptId:         item.DeptId,
	//		CreateBy:       item.CreateBy,
	//		CreateTime:     item.CreateTime,
	//		LastUpdateBy:   item.LastUpdateBy,
	//		LastUpdateTime: item.LastUpdateTime,
	//		DelFlag:        item.DelFlag,
	//		JobId:          item.JobId,
	//	})
	//}
	//return rv, nil
	return nil, nil
}

func (u userRepo) DeleteUser(ctx context.Context, id int64) error {
	panic("implement me")
}