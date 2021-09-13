package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	"kratos-mall/app/front/admin/internal/biz"
	"kratos-mall/app/front/admin/internal/biz/sys"

	pb "kratos-mall/api/front/admin/v1"
)

type SysService struct {
	pb.UnimplementedSysServer
	uc          *biz.UserUseCase
	logUseCase  *sys.LogUseCase
	menuUseCase *sys.MenuUseCase
	roleUseCase *sys.RoleUseCase
	jobUseCase  *sys.JobUseCase
	dictUseCase *sys.DictUseCase
	deptUseCase *sys.DeptUseCase
	log         *log.Helper
}

func NewSysService(uc *biz.UserUseCase,
	logUseCase *sys.LogUseCase,
	mc *sys.MenuUseCase,
	logger log.Logger,
	roleUseCase *sys.RoleUseCase,
	jobUseCase *sys.JobUseCase,
	dictUseCase *sys.DictUseCase,
	deptUseCase *sys.DeptUseCase) *SysService {
	return &SysService{uc: uc,
		logUseCase:  logUseCase,
		menuUseCase: mc,
		log:         log.NewHelper(logger),
		roleUseCase: roleUseCase,
		jobUseCase:  jobUseCase,
		dictUseCase: dictUseCase,
		deptUseCase: deptUseCase}
}

func (s *SysService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {

	loginResp, _ := s.uc.Login(ctx, &biz.LoginDTO{
		UserName: req.UserName,
		Password: req.Password,
	})

	return &pb.LoginResp{
		Status:           loginResp.Status,
		CurrentAuthority: loginResp.CurrentAuthority,
		Id:               loginResp.Id,
		UserName:         loginResp.UserName,
		AccessToken:      loginResp.AccessToken,
		AccessExpire:     loginResp.AccessExpire,
		RefreshAfter:     loginResp.RefreshAfter,
		Code:             "000000",
		Message:          "登录成功",
	}, nil

}
func (s *SysService) UserInfo(ctx context.Context, req *pb.InfoReq) (*pb.InfoResp, error) {

	userInfo, _ := s.uc.UserInfo(ctx, req.UserId)

	rv := make([]*pb.MenuListTree, 0)
	for _, m := range userInfo.MenuListTree {
		rv = append(rv, &pb.MenuListTree{
			Id:       m.Id,
			Path:     m.Url,
			Name:     m.Name,
			ParentId: m.ParentId,
			Icon:     m.Icon,
		})
	}

	return &pb.InfoResp{
		Avatar:       userInfo.Avatar,
		Name:         userInfo.Name,
		MenuListTree: rv,
		Code:         "000000",
		Message:      "获取个人信息成功",
	}, nil
}
func (s *SysService) UserAdd(ctx context.Context, req *pb.UserAddReq) (*pb.UserAddResp, error) {
	return &pb.UserAddResp{}, nil
}
func (s *SysService) UserList(ctx context.Context, req *pb.UserListReq) (*pb.UserListResp, error) {

	userListResp, _ := s.uc.UserList(ctx, &biz.UserListReq{
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
	for _, item := range userListResp.List {
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
		Code:     "000000",
		Message:  "查询用户列表成功",
		Current:  req.Current,
		PageSize: req.PageSize,
		Total:    userListResp.Total,
		Data:     list,
		Success:  true,
	}, nil
}
func (s *SysService) UserUpdate(ctx context.Context, req *pb.UserUpdateReq) (*pb.UserUpdateResp, error) {
	return &pb.UserUpdateResp{}, nil
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
	listResp, _ := s.roleUseCase.ListRole(ctx, &sys.RoleListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.RoleListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.RoleListResp{
		Code:     "000000",
		Message:  "查询订单信息成功",
		Current:  req.Current,
		PageSize: req.PageSize,
		Total:    listResp.Total,
		Data:     list,
		Success:  true,
	}, nil
}
func (s *SysService) RoleUpdate(ctx context.Context, req *pb.RoleUpdateReq) (*pb.RoleUpdateResp, error) {
	return &pb.RoleUpdateResp{}, nil
}
func (s *SysService) RoleDelete(ctx context.Context, req *pb.RoleDeleteReq) (*pb.RoleDeleteResp, error) {
	return &pb.RoleDeleteResp{}, nil
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
	listResp, _ := s.menuUseCase.ListMenu(ctx, &sys.MenuListReq{
		Name: "",
		Url:  "",
	})

	list := make([]*pb.MenuListData, 0)

	copier.Copy(&list, &listResp)
	return &pb.MenuListResp{
		Code:    "000000",
		Message: "查询订单信息成功",
		Data:    list,
		Success: true,
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
	listResp, _ := s.dictUseCase.ListDict(ctx, &sys.DictListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.DictListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.DictListResp{
		Code:     "000000",
		Message:  "查询订单信息成功",
		Current:  req.Current,
		PageSize: req.PageSize,
		Total:    listResp.Total,
		Data:     list,
		Success:  true,
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
	listResp, _ := s.deptUseCase.ListDept(ctx, &sys.DeptListReq{
		Name:     "",
		CreateBy: "",
	})

	list := make([]*pb.DeptListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.DeptListResp{
		Code:    "000000",
		Message: "查询订单信息成功",
		Total:   listResp.Total,
		Data:    list,
		Success: true,
	}, nil
}
func (s *SysService) DeptUpdate(ctx context.Context, req *pb.DeptUpdateReq) (*pb.DeptUpdateResp, error) {
	return &pb.DeptUpdateResp{}, nil
}
func (s *SysService) DeptDelete(ctx context.Context, req *pb.DeptDeleteReq) (*pb.DeptDeleteResp, error) {
	return &pb.DeptDeleteResp{}, nil
}
func (s *SysService) LoginLogList(ctx context.Context, req *pb.LoginLogListReq) (*pb.LoginLogListResp, error) {
	listResp, _ := s.logUseCase.ListLog(ctx, &sys.LogListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.LoginLogListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.LoginLogListResp{
		Code:     "000000",
		Message:  "查询订单信息成功",
		Current:  req.Current,
		PageSize: req.PageSize,
		Total:    listResp.Total,
		Data:     list,
		Success:  true,
	}, nil
}
func (s *SysService) LoginLogDelete(ctx context.Context, req *pb.LoginLogDeleteReq) (*pb.LoginLogDeleteResp, error) {
	return &pb.LoginLogDeleteResp{}, nil
}
func (s *SysService) SysLogList(ctx context.Context, req *pb.SysLogListReq) (*pb.SysLogListResp, error) {

	return &pb.SysLogListResp{}, nil
}
func (s *SysService) SysLogDelete(ctx context.Context, req *pb.SysLogDeleteReq) (*pb.SysLogDeleteResp, error) {
	return &pb.SysLogDeleteResp{}, nil
}
func (s *SysService) JobAdd(ctx context.Context, req *pb.JobAddReq) (*pb.JobAddResp, error) {
	return &pb.JobAddResp{}, nil
}
func (s *SysService) JobList(ctx context.Context, req *pb.JobListReq) (*pb.JobListResp, error) {
	listResp, _ := s.jobUseCase.ListJob(ctx, &sys.JobListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.JobListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.JobListResp{
		Code:     "000000",
		Message:  "查询订单信息成功",
		Current:  req.Current,
		PageSize: req.PageSize,
		Total:    listResp.Total,
		Data:     list,
		Success:  true,
	}, nil
}
func (s *SysService) JobUpdate(ctx context.Context, req *pb.JobUpdateReq) (*pb.JobUpdateResp, error) {
	return &pb.JobUpdateResp{}, nil
}
func (s *SysService) JobDelete(ctx context.Context, req *pb.JobDeleteReq) (*pb.JobDeleteResp, error) {
	return &pb.JobDeleteResp{}, nil
}
