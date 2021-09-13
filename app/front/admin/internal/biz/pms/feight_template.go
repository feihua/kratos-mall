package pms

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
type FeightTemplateListResp struct {
	Total int64
	List  []*FeightTemplate
}

type FeightTemplateRepo interface {
	CreateFeightTemplate(context.Context, *FeightTemplate) error
	GetFeightTemplate(ctx context.Context, id int64) error
	UpdateFeightTemplate(context.Context, *FeightTemplate) error
	ListFeightTemplate(ctx context.Context, req *FeightTemplateListReq) (*FeightTemplateListResp, error)
	DeleteFeightTemplate(ctx context.Context, id int64) error
}

type FeightTemplateUseCase struct {
	repo FeightTemplateRepo
	log  *log.Helper
}

func NewFeightTemplateUseCase(repo FeightTemplateRepo, logger log.Logger) *FeightTemplateUseCase {
	return &FeightTemplateUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (f FeightTemplateUseCase) CreateFeightTemplate(ctx context.Context, template *FeightTemplate) error {
	panic("implement me")
}

func (f FeightTemplateUseCase) GetFeightTemplate(ctx context.Context, id int64) error {
	panic("implement me")
}

func (f FeightTemplateUseCase) UpdateFeightTemplate(ctx context.Context, template *FeightTemplate) error {
	panic("implement me")
}

func (f FeightTemplateUseCase) ListFeightTemplate(ctx context.Context, req *FeightTemplateListReq) (*FeightTemplateListResp, error) {
	return f.repo.ListFeightTemplate(ctx, req)
}

func (f FeightTemplateUseCase) DeleteFeightTemplate(ctx context.Context, id int64) error {
	panic("implement me")
}
