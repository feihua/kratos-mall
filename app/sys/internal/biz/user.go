package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
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
	return "sys_user"
}

type UserRepo interface {
	CreateUser(context.Context, *UserDTO) error
	GetUser(ctx context.Context, id int64) *User
	UpdateUser(context.Context, *UserDTO) error
	ListUser(ctx context.Context, pageNum, pageSize int64) ([]*Beer, error)
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
		UserName:         user.Username,
		AccessToken:      "test",
		AccessExpire:     0,
		RefreshAfter:     0,
	}, nil

}

func (u *UserUseCase) UserInfo(ctx context.Context, id int64) *UserInfoDTO {

	user := u.userRepo.GetUser(ctx, id)

	listMenu, _ := u.menuRepo.ListMenu(ctx, 1, 1000)

	return &UserInfoDTO{
		Avatar:         "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
		Name:           user.Username,
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

func (u *UserUseCase) ListUser(ctx context.Context, pageNum, pageSize int64) ([]*Beer, error) {
	panic("implement me")
}

func (u *UserUseCase) DeleteUser(ctx context.Context, id int64) error {
	panic("implement me")
}
