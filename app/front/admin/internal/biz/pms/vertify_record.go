package pms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type VertifyRecordListReq struct {
	Current  int64
	PageSize int64
}

type VertifyRecord struct {
	Id         int64
	ProductId  int64
	CreateTime string
	VertifyMan string
	Status     int
	Detail     string
}

type VertifyRecordRepo interface {
	CreateVertifyRecord(context.Context, *VertifyRecord) error
	GetVertifyRecord(ctx context.Context, id int64) error
	UpdateVertifyRecord(context.Context, *VertifyRecord) error
	ListVertifyRecord(ctx context.Context, req *VertifyRecordListReq) ([]*VertifyRecord, error)
	DeleteVertifyRecord(ctx context.Context, id int64) error
}

type VertifyRecordUseCase struct {
	repo VertifyRecordRepo
	log  *log.Helper
}

func NewVertifyRecordUseCase(repo VertifyRecordRepo, logger log.Logger) *VertifyRecordUseCase {
	return &VertifyRecordUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (r *VertifyRecordUseCase) CreateVertifyRecord(ctx context.Context, user *VertifyRecord) error {
	panic("implement me")
}

func (r *VertifyRecordUseCase) GetVertifyRecord(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r *VertifyRecordUseCase) UpdateVertifyRecord(ctx context.Context, user *VertifyRecord) error {
	panic("implement me")
}

func (r *VertifyRecordUseCase) ListVertifyRecord(ctx context.Context, pageNum, pageSize int64) ([]*VertifyRecord, error) {
	panic("implement me")
}

func (r *VertifyRecordUseCase) DeleteVertifyRecord(ctx context.Context, id int64) error {
	panic("implement me")
}
