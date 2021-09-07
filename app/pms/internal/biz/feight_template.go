package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type FeightTemplateListReq struct {
	Current  int64
	PageSize int64
}

type FeightTemplate struct {
	Id             int64
	Name           string
	ChargeType     int    // 计费类型:0->按重量；1->按件数
	FirstWeight    string // 首重kg
	FirstFee       string // 首费（元）
	ContinueWeight string
	ContinmeFee    string
	Dest           string // 目的地（省、市）
}

type FeightTemplateRepo interface {
	CreateFeightTemplate(context.Context, *FeightTemplate) error
	GetFeightTemplate(ctx context.Context, id int64) error
	UpdateFeightTemplate(context.Context, *FeightTemplate) error
	ListFeightTemplate(ctx context.Context, req *FeightTemplateListReq) ([]*FeightTemplate, error)
	DeleteFeightTemplate(ctx context.Context, id int64) error
}

type FeightTemplateUseCase struct {
	repo FeightTemplateRepo
	log  *log.Helper
}

func NewFeightTemplateUseCase(repo FeightTemplateRepo, logger log.Logger) *FeightTemplateUseCase {
	return &FeightTemplateUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (r *FeightTemplateUseCase) CreateFeightTemplate(ctx context.Context, user *FeightTemplate) error {
	panic("implement me")
}

func (r *FeightTemplateUseCase) GetFeightTemplate(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r *FeightTemplateUseCase) UpdateFeightTemplate(ctx context.Context, user *FeightTemplate) error {
	panic("implement me")
}

func (r *FeightTemplateUseCase) ListFeightTemplate(ctx context.Context, pageNum, pageSize int64) ([]*FeightTemplate, error) {
	panic("implement me")
}

func (r *FeightTemplateUseCase) DeleteFeightTemplate(ctx context.Context, id int64) error {
	panic("implement me")
}
