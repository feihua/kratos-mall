package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type UserDTO struct {
	UserName string
	Password string
}

type UserInfoDTO struct {
	Avatar         string
	Name           string
	MenuListTree   []*Menu
	BackgroundUrls []string
}

type TokenDTO struct {
	Status           string
	CurrentAuthority string
	Id               int64
	UserName         string
	AccessToken      string
	AccessExpire     int64
	RefreshAfter     int64
}

type Beer struct {
	Id          int64
	Name        string
	Description string
	Count       int64
}

type UserListReq struct {
	Current  int64
	PageSize int64
	Name     string
	NickName string
	Mobile   string
	Email    string
	Status   int64
	DeptId   int64
}

type User struct {
	Id             int64  // 编号
	Name           string // 用户名
	NickName       string // 昵称
	Avatar         string // 头像
	Password       string // 密码
	Salt           string // 加密盐
	Email          string // 邮箱
	Mobile         string // 手机号
	Status         int    // 状态  0：禁用   1：正常
	DeptId         int64  // 机构ID
	CreateBy       string // 创建人
	CreateTime     string // 创建时间
	LastUpdateBy   string // 更新人
	LastUpdateTime string // 更新时间
	DelFlag        int    // 是否删除  -1：已删除  0：正常
	JobId          int    // 岗位Id
}

type UserRepo interface {
	CreateUser(context.Context, *UserDTO) error
	GetUser(ctx context.Context, id int64) *User
	UpdateUser(context.Context, *UserDTO) error
	ListUser(ctx context.Context, req *UserListReq) ([]*User, error)
	DeleteUser(ctx context.Context, id int64) error
	QueryUserByName(ctx context.Context, name string) *User
}

type UserUseCase struct {
	userRepo UserRepo
	menuRepo MenuRepo
	log      *log.Helper
}

func NewUserUseCase(userRepo UserRepo, menuRepo MenuRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{userRepo: userRepo, menuRepo: menuRepo, log: log.NewHelper(logger)}
}

func (u *UserUseCase) UserLogin(ctx context.Context, userDTO *UserDTO) (*TokenDTO, error) {

	user := u.userRepo.QueryUserByName(ctx, userDTO.UserName)

	return &TokenDTO{
		Status:           "1",
		CurrentAuthority: "admin",
		Id:               user.Id,
		UserName:         user.Name,
		AccessToken:      "test",
		AccessExpire:     0,
		RefreshAfter:     0,
	}, nil

}

func (u *UserUseCase) UserInfo(ctx context.Context, id int64) *UserInfoDTO {

	user := u.userRepo.GetUser(ctx, id)

	listMenu, _ := u.menuRepo.ListMenu(ctx, &MenuListReq{})

	return &UserInfoDTO{
		Avatar:         "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
		Name:           user.Name,
		MenuListTree:   listMenu,
		BackgroundUrls: nil,
	}
}

func (u *UserUseCase) CreateUser(ctx context.Context, user *UserDTO) error {
	panic("implement me")
}

func (u *UserUseCase) GetUser(ctx context.Context, id int64) error {
	panic("implement me")
}

func (u *UserUseCase) UpdateUser(ctx context.Context, user *UserDTO) error {
	panic("implement me")
}

func (u *UserUseCase) ListUser(ctx context.Context, req *UserListReq) ([]*User, error) {
	return u.userRepo.ListUser(ctx, req)
}

func (u *UserUseCase) DeleteUser(ctx context.Context, id int64) error {
	panic("implement me")
}
