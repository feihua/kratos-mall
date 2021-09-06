package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"strings"
)

type LoginDTO struct {
	UserName string
	Password string
}

type LoginRespDTO struct {
	Status           string
	CurrentAuthority string
	Id               int64
	UserName         string
	AccessToken      string
	AccessExpire     int64
	RefreshAfter     int64
}

type UserInfoDTO struct {
	Avatar         string
	Name           string
	MenuListTree   []*MenuDTO
	BackgroundUrls []string
}

type MenuDTO struct {
	Id       int64
	Name     string
	ParentId int64
	Url      string
	Perms    string
	Type     int
	Icon     string
}

type UserRepo interface {
	Login(context.Context, *LoginDTO) (*LoginRespDTO, error)
	LoginLogAdd(context.Context, *LoginDTO) error
	UserInfo(context.Context, int64) (*UserInfoDTO, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUseCase) Login(ctx context.Context, req *LoginDTO) (*LoginRespDTO, error) {

	if len(strings.TrimSpace(req.UserName)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		uc.log.WithContext(ctx).Errorf("用户名或密码不能为空,请求信息失败,参数: %v", req)
		return nil, errors.BadRequest("用户名或密码不能为空", "用户名或密码不能为空")
	}

	loginResp, _ := uc.repo.Login(ctx, req)

	_ = uc.repo.LoginLogAdd(ctx, req)

	uc.log.WithContext(ctx).Infof("登录成功: %v", loginResp)

	return loginResp, nil
}

func (uc *UserUseCase) UserInfo(ctx context.Context, id int64) (*UserInfoDTO, error) {

	return uc.repo.UserInfo(ctx, id)

}
