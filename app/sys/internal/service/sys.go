package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	"kratos-mall/app/sys/internal/biz"

	pb "kratos-mall/api/sys/v1"
)

type SysService struct {
	pb.UnimplementedSysServer
	uc          *biz.UserUseCase
	lc          *biz.LogUseCase
	mc          *biz.MenuUseCase
	roleUseCase *biz.RoleUseCase
	jobUseCase  *biz.JobUseCase
	dictUseCase *biz.DictUseCase
	deptUseCase *biz.DeptUseCase
	log         *log.Helper
}

func NewSysService(uc *biz.UserUseCase,
	lc *biz.LogUseCase,
	mc *biz.MenuUseCase,
	logger log.Logger,
	roleUseCase *biz.RoleUseCase,
	jobUseCase *biz.JobUseCase,
	dictUseCase *biz.DictUseCase,
	deptUseCase *biz.DeptUseCase) *SysService {
	return &SysService{uc: uc,
		lc:          lc,
		mc:          mc,
		log:         log.NewHelper(logger),
		roleUseCase: roleUseCase,
		jobUseCase:  jobUseCase,
		dictUseCase: dictUseCase,
		deptUseCase: deptUseCase}
}

func (s *SysService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	tokenDTO, _ := s.uc.UserLogin(ctx, &biz.UserDTO{
		UserName: req.UserName,
		Password: req.Password,
	})
	return &pb.LoginResp{
		Status:           tokenDTO.Status,
		CurrentAuthority: tokenDTO.CurrentAuthority,
		Id:               tokenDTO.Id,
		UserName:         tokenDTO.UserName,
		AccessToken:      tokenDTO.AccessToken,
		AccessExpire:     tokenDTO.AccessExpire,
		RefreshAfter:     tokenDTO.RefreshAfter,
	}, nil
}
func (s *SysService) UserInfo(ctx context.Context, req *pb.InfoReq) (*pb.InfoResp, error) {

	info := s.uc.UserInfo(ctx, req.UserId)

	rv := make([]*pb.MenuListTree, 0)
	for _, m := range info.MenuListTree {
		rv = append(rv, &pb.MenuListTree{
			Id:       m.Id,
			Name:     m.Name,
			ParentId: m.ParentId,
			Icon:     m.Icon,
			Path:     m.Url,
		})
	}

	return &pb.InfoResp{
		Avatar:         info.Avatar,
		Name:           info.Name,
		MenuListTree:   rv,
		BackgroundUrls: nil,
	}, nil
}
func (s *SysService) UserAdd(ctx context.Context, req *pb.UserAddReq) (*pb.UserAddResp, error) {
	return &pb.UserAddResp{}, nil
}
func (s *SysService) UserList(ctx context.Context, req *pb.UserListReq) (*pb.UserListResp, error) {
	listUser, _ := s.uc.ListUser(ctx, &biz.UserListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		Name:     req.Name,
		NickName: req.NickName,
		Mobile:   req.Mobile,
		Email:    req.Email,
		Status:   req.Status,
		DeptId:   req.DeptId,
	})

	list := make([]*pb.UserListData, 0)
	for _, item := range listUser {
		list = append(list, &pb.UserListData{
			Id:             item.Id,
			Name:           item.Name,
			NickName:       item.NickName,
			Avatar:         item.Avatar,
			Password:       item.Password,
			Salt:           item.Salt,
			Email:          item.Email,
			Mobile:         item.Mobile,
			Status:         int64(item.Status),
			DeptId:         item.DeptId,
			CreateBy:       item.CreateBy,
			CreateTime:     item.CreateTime,
			LastUpdateBy:   item.LastUpdateBy,
			LastUpdateTime: item.LastUpdateTime,
			DelFlag:        int64(item.DelFlag),
			JobId:          int64(item.JobId),
		})
	}

	return &pb.UserListResp{
		Total: int64(len(list)),
		List:  list,
	}, nil
}
func (s *SysService) UserUpdate(ctx context.Context, req *pb.UserUpdateReq) (*pb.UserUpdateResp, error) {

	var tempUser *biz.User
	_ = copier.Copy(&tempUser, &req)
	s.uc.UpdateUser(ctx, tempUser)

	return &pb.UserUpdateResp{
		Pong: "",
	}, nil
}
func (s *SysService) UserDelete(ctx context.Context, req *pb.UserDeleteReq) (*pb.UserDeleteResp, error) {
	return &pb.UserDeleteResp{}, nil
}
func (s *SysService) ReSetPassword(ctx context.Context, req *pb.ReSetPasswordReq) (*pb.ReSetPasswordResp, error) {
	return &pb.ReSetPasswordResp{}, nil
}
func (s *SysService) UpdateUserStatus(ctx context.Context, req *pb.UserStatusReq) (*pb.UserStatusResp, error) {
	return &pb.UserStatusResp{}, nil
}
func (s *SysService) RoleAdd(ctx context.Context, req *pb.RoleAddReq) (*pb.RoleAddResp, error) {
	return &pb.RoleAddResp{}, nil
}
func (s *SysService) RoleList(ctx context.Context, req *pb.RoleListReq) (*pb.RoleListResp, error) {
	listResp, _ := s.roleUseCase.ListRole(ctx, &biz.RoleListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.RoleListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.RoleListResp{
		Total: listResp.Total,
		List:  list,
	}, nil
}
func (s *SysService) RoleUpdate(ctx context.Context, req *pb.RoleUpdateReq) (*pb.RoleUpdateResp, error) {
	return &pb.RoleUpdateResp{}, nil
}
func (s *SysService) RoleDelete(ctx context.Context, req *pb.RoleDeleteReq) (*pb.RoleDeleteResp, error) {
	return &pb.RoleDeleteResp{}, nil
}
func (s *SysService) UpdateRoleRole(ctx context.Context, req *pb.UpdateRoleRoleReq) (*pb.UpdateRoleRoleResp, error) {
	return &pb.UpdateRoleRoleResp{}, nil
}
func (s *SysService) QueryMenuByRoleId(ctx context.Context, req *pb.QueryMenuByRoleIdReq) (*pb.QueryMenuByRoleIdResp, error) {
	return &pb.QueryMenuByRoleIdResp{}, nil
}
func (s *SysService) UpdateMenuRole(ctx context.Context, req *pb.UpdateMenuRoleReq) (*pb.UpdateMenuRoleResp, error) {
	return &pb.UpdateMenuRoleResp{}, nil
}
func (s *SysService) MenuAdd(ctx context.Context, req *pb.MenuAddReq) (*pb.MenuAddResp, error) {
	return &pb.MenuAddResp{}, nil
}
func (s *SysService) MenuList(ctx context.Context, req *pb.MenuListReq) (*pb.MenuListResp, error) {

	listMenu, _ := s.mc.ListMenu(ctx, &biz.MenuListReq{
		Name: req.Name,
		Url:  req.Url,
	})

	rv := make([]*pb.MenuListData, 0)
	for _, m := range listMenu {
		rv = append(rv, &pb.MenuListData{
			Id:             m.Id,
			Name:           m.Name,
			ParentId:       m.ParentId,
			Url:            m.Url,
			Perms:          m.Perms,
			Type:           int64(m.Type),
			Icon:           m.Icon,
			OrderNum:       int64(m.OrderNum),
			CreateBy:       m.CreateBy,
			CreateTime:     m.CreateTime,
			LastUpdateBy:   m.LastUpdateBy,
			LastUpdateTime: m.LastUpdateTime,
			DelFlag:        int64(m.DelFlag),
		})
	}

	return &pb.MenuListResp{
		Total: int64(len(rv)),
		List:  rv,
	}, nil
}
func (s *SysService) MenuUpdate(ctx context.Context, req *pb.MenuUpdateReq) (*pb.MenuUpdateResp, error) {
	return &pb.MenuUpdateResp{}, nil
}
func (s *SysService) MenuDelete(ctx context.Context, req *pb.MenuDeleteReq) (*pb.MenuDeleteResp, error) {
	return &pb.MenuDeleteResp{}, nil
}
func (s *SysService) DictAdd(ctx context.Context, req *pb.DictAddReq) (*pb.DictAddResp, error) {
	return &pb.DictAddResp{}, nil
}
func (s *SysService) DictList(ctx context.Context, req *pb.DictListReq) (*pb.DictListResp, error) {
	listResp, _ := s.dictUseCase.ListDict(ctx, &biz.DictListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.DictListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.DictListResp{
		Total: listResp.Total,
		List:  list,
	}, nil
}
func (s *SysService) DictUpdate(ctx context.Context, req *pb.DictUpdateReq) (*pb.DictUpdateResp, error) {
	return &pb.DictUpdateResp{}, nil
}
func (s *SysService) DictDelete(ctx context.Context, req *pb.DictDeleteReq) (*pb.DictDeleteResp, error) {
	return &pb.DictDeleteResp{}, nil
}
func (s *SysService) DeptAdd(ctx context.Context, req *pb.DeptAddReq) (*pb.DeptAddResp, error) {
	return &pb.DeptAddResp{}, nil
}
func (s *SysService) DeptList(ctx context.Context, req *pb.DeptListReq) (*pb.DeptListResp, error) {
	listResp, _ := s.deptUseCase.ListDept(ctx, &biz.DeptListReq{
		Name:     "",
		CreateBy: "",
	})

	list := make([]*pb.DeptListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.DeptListResp{
		Total: listResp.Total,
		List:  list,
	}, nil
}
func (s *SysService) DeptUpdate(ctx context.Context, req *pb.DeptUpdateReq) (*pb.DeptUpdateResp, error) {
	return &pb.DeptUpdateResp{}, nil
}
func (s *SysService) DeptDelete(ctx context.Context, req *pb.DeptDeleteReq) (*pb.DeptDeleteResp, error) {
	return &pb.DeptDeleteResp{}, nil
}
func (s *SysService) LoginLogAdd(ctx context.Context, req *pb.LoginLogAddReq) (*pb.LoginLogAddResp, error) {

	_ = s.lc.LoginLogAdd(ctx, &biz.LogDTO{
		UserName: req.UserName,
		Status:   req.Status,
		Ip:       req.Ip,
		CreateBy: req.CreateBy,
	})

	return &pb.LoginLogAddResp{
		Pong: "",
	}, nil
}
func (s *SysService) LoginLogList(ctx context.Context, req *pb.LoginLogListReq) (*pb.LoginLogListResp, error) {
	listResp, _ := s.lc.ListLog(ctx, &biz.LogListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.LoginLogListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.LoginLogListResp{
		Total: listResp.Total,
		List:  list,
	}, nil
}
func (s *SysService) LoginLogDelete(ctx context.Context, req *pb.LoginLogDeleteReq) (*pb.LoginLogDeleteResp, error) {
	return &pb.LoginLogDeleteResp{}, nil
}
func (s *SysService) SysLogAdd(ctx context.Context, req *pb.SysLogAddReq) (*pb.SysLogAddResp, error) {
	return &pb.SysLogAddResp{}, nil
}
func (s *SysService) SysLogList(ctx context.Context, req *pb.SysLogListReq) (*pb.SysLogListResp, error) {
	return &pb.SysLogListResp{}, nil
}
func (s *SysService) SysLogDelete(ctx context.Context, req *pb.SysLogDeleteReq) (*pb.SysLogDeleteResp, error) {
	return &pb.SysLogDeleteResp{}, nil
}
func (s *SysService) ConfigAdd(ctx context.Context, req *pb.ConfigAddReq) (*pb.ConfigAddResp, error) {
	return &pb.ConfigAddResp{}, nil
}
func (s *SysService) ConfigList(ctx context.Context, req *pb.ConfigListReq) (*pb.ConfigListResp, error) {
	return &pb.ConfigListResp{}, nil
}
func (s *SysService) ConfigUpdate(ctx context.Context, req *pb.ConfigUpdateReq) (*pb.ConfigUpdateResp, error) {
	return &pb.ConfigUpdateResp{}, nil
}
func (s *SysService) ConfigDelete(ctx context.Context, req *pb.ConfigDeleteReq) (*pb.ConfigDeleteResp, error) {
	return &pb.ConfigDeleteResp{}, nil
}
func (s *SysService) UpdateConfigRole(ctx context.Context, req *pb.UpdateConfigRoleReq) (*pb.UpdateConfigRoleResp, error) {
	return &pb.UpdateConfigRoleResp{}, nil
}
func (s *SysService) JobAdd(ctx context.Context, req *pb.JobAddReq) (*pb.JobAddResp, error) {
	return &pb.JobAddResp{}, nil
}
func (s *SysService) JobList(ctx context.Context, req *pb.JobListReq) (*pb.JobListResp, error) {
	listResp, _ := s.jobUseCase.ListJob(ctx, &biz.JobListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.JobListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.JobListResp{
		Total: listResp.Total,
		List:  list,
	}, nil
}
func (s *SysService) JobUpdate(ctx context.Context, req *pb.JobUpdateReq) (*pb.JobUpdateResp, error) {
	return &pb.JobUpdateResp{}, nil
}
func (s *SysService) JobDelete(ctx context.Context, req *pb.JobDeleteReq) (*pb.JobDeleteResp, error) {
	return &pb.JobDeleteResp{}, nil
}
