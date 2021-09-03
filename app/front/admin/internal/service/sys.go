package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/front/admin/internal/biz"

	pb "kratos-mall/api/front/admin/v1"
)

type SysService struct {
	pb.UnimplementedSysServer
	uc  *biz.UserUseCase
	log *log.Helper
}

func NewSysService(uc *biz.UserUseCase, logger log.Logger) *SysService {
	return &SysService{uc: uc, log: log.NewHelper(logger)}
}

func (s *SysService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	fmt.Printf("%s", "hello")
	login, _ := s.uc.Login(ctx, &biz.Login{
		UserName: req.UserName,
		Password: req.Password,
	})
	return &pb.LoginResp{
		Status:           login.Status,
		CurrentAuthority: login.CurrentAuthority,
		Id:               login.Id,
		UserName:         login.UserName,
		AccessToken:      login.AccessToken,
		AccessExpire:     0,
		RefreshAfter:     0,
	}, nil
}
func (s *SysService) UserInfo(ctx context.Context, req *pb.InfoReq) (*pb.InfoResp, error) {
	return &pb.InfoResp{}, nil
}
func (s *SysService) UserAdd(ctx context.Context, req *pb.UserAddReq) (*pb.UserAddResp, error) {
	return &pb.UserAddResp{}, nil
}
func (s *SysService) UserList(ctx context.Context, req *pb.UserListReq) (*pb.UserListResp, error) {
	return &pb.UserListResp{}, nil
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
	return &pb.RoleListResp{}, nil
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
	return &pb.MenuListResp{}, nil
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
	return &pb.DictListResp{}, nil
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
	return &pb.DeptListResp{}, nil
}
func (s *SysService) DeptUpdate(ctx context.Context, req *pb.DeptUpdateReq) (*pb.DeptUpdateResp, error) {
	return &pb.DeptUpdateResp{}, nil
}
func (s *SysService) DeptDelete(ctx context.Context, req *pb.DeptDeleteReq) (*pb.DeptDeleteResp, error) {
	return &pb.DeptDeleteResp{}, nil
}
func (s *SysService) LoginLogList(ctx context.Context, req *pb.LoginLogListReq) (*pb.LoginLogListResp, error) {
	return &pb.LoginLogListResp{}, nil
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
	return &pb.JobListResp{}, nil
}
func (s *SysService) JobUpdate(ctx context.Context, req *pb.JobUpdateReq) (*pb.JobUpdateResp, error) {
	return &pb.JobUpdateResp{}, nil
}
func (s *SysService) JobDelete(ctx context.Context, req *pb.JobDeleteReq) (*pb.JobDeleteResp, error) {
	return &pb.JobDeleteResp{}, nil
}
