package sys

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	sysV1 "kratos-mall/api/sys/v1"
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

func (u userRepo) ListUser(ctx context.Context, req *sys.UserListReq) (*sys.UserListResp, error) {
	list, _ := u.data.SysClient.UserList(ctx, &sysV1.UserListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*sys.User, 0)
	copier.Copy(&orders, &list.List)

	return &sys.UserListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (u userRepo) DeleteUser(ctx context.Context, id int64) error {
	panic("implement me")
}
