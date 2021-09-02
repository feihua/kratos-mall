package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Dict struct {
	Hello string
}

type DictRepo interface {
	CreateDict(context.Context, *Dict) error
	GetDict(ctx context.Context, id int64) error
	UpdateDict(context.Context, *Dict) error
	ListDict(ctx context.Context, pageNum, pageSize int64) ([]*Beer, error)
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

func (d *DictUseCase) ListDict(ctx context.Context, pageNum, pageSize int64) ([]*Beer, error) {
	panic("implement me")
}

func (d *DictUseCase) DeleteDict(ctx context.Context, id int64) error {
	panic("implement me")
}
