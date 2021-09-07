package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	sysV1 "kratos-mall/api/sys/v1"
	"kratos-mall/app/front/admin/internal/biz"
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

func (r *userRepo) Login(ctx context.Context, g *biz.LoginDTO) (*biz.LoginRespDTO, error) {
	login, _ := r.data.sysClient.Login(ctx, &sysV1.LoginReq{
		UserName: g.UserName,
		Password: g.Password,
	})

	return &biz.LoginRespDTO{
		Status:           login.Status,
		CurrentAuthority: login.CurrentAuthority,
		Id:               login.Id,
		UserName:         login.UserName,
		AccessToken:      login.AccessToken,
		AccessExpire:     login.AccessExpire,
		RefreshAfter:     login.RefreshAfter,
	}, nil
}

func (r *userRepo) LoginLogAdd(ctx context.Context, req *biz.LoginDTO) error {

	_, err := r.data.sysClient.LoginLogAdd(ctx, &sysV1.LoginLogAddReq{
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

func (r *userRepo) UserInfo(ctx context.Context, id int64) (*biz.UserInfoDTO, error) {

	userInfo, _ := r.data.sysClient.UserInfo(ctx, &sysV1.InfoReq{
		UserId: id,
	})

	rv := make([]*biz.MenuDTO, 0)
	for _, m := range userInfo.MenuListTree {
		rv = append(rv, &biz.MenuDTO{
			Id:       m.Id,
			Url:      m.Path,
			Name:     m.Name,
			ParentId: m.ParentId,
			Icon:     m.Icon,
		})
	}

	return &biz.UserInfoDTO{
		Avatar:         userInfo.Avatar,
		Name:           userInfo.Name,
		MenuListTree:   rv,
		BackgroundUrls: nil,
	}, nil

}

func (r *userRepo) UserList(ctx context.Context, req *biz.UserListReq) (*biz.UserListResp, error) {

	userListResp, _ := r.data.sysClient.UserList(ctx, &sysV1.UserListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		Name:     req.Name,
		NickName: req.NickName,
		Mobile:   req.Mobile,
		Email:    req.Email,
		Status:   req.Status,
		DeptId:   req.DeptId,
	})

	list := make([]*biz.User, 0)
	for _, item := range userListResp.List {
		list = append(list, &biz.User{
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

	return &biz.UserListResp{
		Total: userListResp.Total,
		List:  list,
	}, nil

}
