package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type RecordListReq struct {
	Current  int64
	PageSize int64
}

type Record struct {
	Id         int    // 主键
	BusinessId string // 业务编号
	Amount     string // 金额
	PayType    string // 支付类型(APP:APP支付 JSAPI:小程序,公众号 MWEB:H5支付)
	Remarks    string // 备注
	ReturnCode string
	ReturnMsg  string
	ResultCode string
	ResultMsg  string
	PayStatus  int    // 0：初始化 1：已发送 2：成功 3：失败
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
}

type RecordRepo interface {
	CreateRecord(context.Context, *Record) error
	GetRecord(ctx context.Context, id int64) error
	UpdateRecord(context.Context, *Record) error
	ListRecord(ctx context.Context, req *RecordListReq) ([]*Record, error)
	DeleteRecord(ctx context.Context, id int64) error
}

type RecordUseCase struct {
	repo RecordRepo
	log  *log.Helper
}

func NewRecordUseCase(repo RecordRepo, logger log.Logger) *RecordUseCase {
	return &RecordUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (r *RecordUseCase) CreateRecord(ctx context.Context, user *Record) error {
	panic("implement me")
}

func (r *RecordUseCase) GetRecord(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r *RecordUseCase) UpdateRecord(ctx context.Context, user *Record) error {
	panic("implement me")
}

func (r *RecordUseCase) ListRecord(ctx context.Context, pageNum, pageSize int64) ([]*Record, error) {
	panic("implement me")
}

func (r *RecordUseCase) DeleteRecord(ctx context.Context, id int64) error {
	panic("implement me")
}
