package sys

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	sysV1 "kratos-mall/api/sys/v1"
	"kratos-mall/app/front/admin/internal/biz/sys"
	"kratos-mall/app/front/admin/internal/data"
	"kratos-mall/app/front/admin/internal/pkg/middleware"
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

func (r *userRepo) Login(ctx context.Context, g *sys.LoginDTO) (*sys.LoginRespDTO, error) {
	login, _ := r.data.SysClient.Login(ctx, &sysV1.LoginReq{
		UserName: g.UserName,
		Password: g.Password,
	})

	token, _ := jwt.GenerateToken(login.Id, login.UserName)

	return &sys.LoginRespDTO{
		Status:           login.Status,
		CurrentAuthority: login.CurrentAuthority,
		Id:               login.Id,
		UserName:         login.UserName,
		AccessToken:      token,
		AccessExpire:     login.AccessExpire,
		RefreshAfter:     login.RefreshAfter,
	}, nil
}

func (r *userRepo) LoginLogAdd(ctx context.Context, req *sys.LoginDTO) error {

	_, err := r.data.SysClient.LoginLogAdd(ctx, &sysV1.LoginLogAddReq{
		UserName: req.UserName,
		Status:   "login",
		Ip:       "127.0.0.1",
		CreateBy: req.UserName,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *userRepo) UserInfo(ctx context.Context, id int64) (*sys.UserInfoDTO, error) {

	userInfo, _ := r.data.SysClient.UserInfo(ctx, &sysV1.InfoReq{
		UserId: id,
	})

	rv := make([]*sys.MenuDTO, 0)
	for _, m := range userInfo.MenuListTree {
		rv = append(rv, &sys.MenuDTO{
			Id:       m.Id,
			Url:      m.Path,
			Name:     m.Name,
			ParentId: m.ParentId,
			Icon:     m.Icon,
		})
	}

	return &sys.UserInfoDTO{
		Avatar:         userInfo.Avatar,
		Name:           userInfo.Name,
		MenuListTree:   rv,
		BackgroundUrls: nil,
	}, nil

}

func (r *userRepo) UserList(ctx context.Context, req *sys.UserListReq) (*sys.UserListResp, error) {

	userListResp, _ := r.data.SysClient.UserList(ctx, &sysV1.UserListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		Name:     req.Name,
		NickName: req.NickName,
		Mobile:   req.Mobile,
		Email:    req.Email,
		Status:   req.Status,
		DeptId:   req.DeptId,
	})

	list := make([]*sys.User, 0)
	for _, item := range userListResp.List {
		list = append(list, &sys.User{
			Id:             item.Id,
			Name:           item.Name,
			NickName:       item.NickName,
			Avatar:         item.Avatar,
			Password:       item.Password,
			Salt:           item.Salt,
			Email:          item.Email,
			Mobile:         item.Mobile,
			Status:         int(item.Status),
			DeptId:         item.DeptId,
			CreateBy:       item.CreateBy,
			CreateTime:     item.CreateTime,
			LastUpdateBy:   item.LastUpdateBy,
			LastUpdateTime: item.LastUpdateTime,
			DelFlag:        int(item.DelFlag),
			JobId:          int(item.JobId),
		})
	}

	return &sys.UserListResp{
		Total: userListResp.Total,
		List:  list,
	}, nil
}

func (r *userRepo) UserUpdate(ctx context.Context, req *sys.User) error {

	var tempUser *sysV1.UserUpdateReq

	_ = copier.Copy(&tempUser, &req)
	_, err := r.data.SysClient.UserUpdate(ctx, tempUser)

	if err != nil {
		return err
	}

	return nil
}
