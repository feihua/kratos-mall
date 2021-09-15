package sys

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	"strconv"
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

type UserListResp struct {
	Total int64
	List  []*User
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

type SelectDataReq struct {
	Current  int64
	PageSize int64
}

type RoleAll struct {
	Id     int32
	Name   string
	Remark string
}

type DeptAll struct {
	Id       int32
	Value    string
	Title    string
	ParentId int32
}

type JobAll struct {
	Id      int32
	JobName string
}

type SelectDataResp struct {
	RoleAll []*RoleAll
	DeptAll []*DeptAll
	JobAll  []*JobAll
}

type UserRepo interface {
	Login(context.Context, *LoginDTO) (*LoginRespDTO, error)
	LoginLogAdd(context.Context, *LoginDTO) error
	UserInfo(context.Context, int64) (*UserInfoDTO, error)
	UserList(context.Context, *UserListReq) (*UserListResp, error)
	UserUpdate(ctx context.Context, req *User) error
}

type UserUseCase struct {
	repo     UserRepo
	roleRepo RoleRepo
	deptRepo DeptRepo
	jobRepo  JobRepo
	log      *log.Helper
}

func NewUserUseCase(repo UserRepo, roleRepo RoleRepo, deptRepo DeptRepo, jobRepo JobRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo,
		log:      log.NewHelper(logger),
		roleRepo: roleRepo,
		deptRepo: deptRepo,
		jobRepo:  jobRepo,
	}
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

func (uc *UserUseCase) UserList(ctx context.Context, req *UserListReq) (*UserListResp, error) {

	return uc.repo.UserList(ctx, req)

}
func (uc *UserUseCase) SelectAllData(ctx context.Context, req *SelectDataReq) (*SelectDataResp, error) {

	listRole, _ := uc.roleRepo.ListRole(ctx, &RoleListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})
	roleAlls := make([]*RoleAll, 0)
	_ = copier.Copy(&roleAlls, &listRole.List)

	listDept, _ := uc.deptRepo.ListDept(ctx, &DeptListReq{})
	deptAlls := make([]*DeptAll, 0)
	//_ = copier.Copy(&deptAlls, &listDept.List)

	for _, dept := range listDept.List {
		deptAlls = append(deptAlls, &DeptAll{
			Id:       int32(dept.Id),
			Value:    strconv.FormatInt(dept.Id, 10),
			Title:    dept.Name,
			ParentId: int32(dept.ParentId),
		})
	}

	listJob, _ := uc.jobRepo.ListJob(ctx, &JobListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})
	jobAlls := make([]*JobAll, 0)
	_ = copier.Copy(&jobAlls, &listJob.List)

	return &SelectDataResp{
		RoleAll: roleAlls,
		DeptAll: deptAlls,
		JobAll:  jobAlls,
	}, nil
}
func (uc *UserUseCase) UserUpdate(ctx context.Context, req *User) error {

	return uc.repo.UserUpdate(ctx, req)

}
