package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type DictListReq struct {
	Current  int64
	PageSize int64
	Value    string
	Label    string
	DelFlag  int64
	Type     string
}

type Dict struct {
	Id             int64  // 编号
	Value          string // 数据值
	Label          string // 标签名
	Type           string // 类型
	Description    string // 描述
	Sort           string // 排序（升序）
	CreateBy       string // 创建人
	CreateTime     string // 创建时间
	LastUpdateBy   string // 更新人
	LastUpdateTime string // 更新时间
	Remarks        string // 备注信息
	DelFlag        int    // 是否删除  -1：已删除  0：正常
}
type DictListResp struct {
	Total int64
	List  []*Dict
}

type DictRepo interface {
	CreateDict(context.Context, *Dict) error
	GetDict(ctx context.Context, id int64) error
	UpdateDict(context.Context, *Dict) error
	ListDict(ctx context.Context, req *DictListReq) (*DictListResp, error)
	DeleteDict(ctx context.Context, id int64) error
}

type DictUseCase struct {
	repo DictRepo
	log  *log.Helper
}

func NewDictUseCase(repo DictRepo, logger log.Logger) *DictUseCase {
	return &DictUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (d *DictUseCase) CreateDict(ctx context.Context, user *Dict) error {
	panic("implement me")
}

func (d *DictUseCase) GetDict(ctx context.Context, id int64) error {
	panic("implement me")
}

func (d *DictUseCase) UpdateDict(ctx context.Context, user *Dict) error {
	panic("implement me")
}

func (d *DictUseCase) ListDict(ctx context.Context, req *DictListReq) (*DictListResp, error) {
	return d.repo.ListDict(ctx, req)
}

func (d *DictUseCase) DeleteDict(ctx context.Context, id int64) error {
	panic("implement me")
}
