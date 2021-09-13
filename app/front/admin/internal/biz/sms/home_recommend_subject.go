package sms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type HomeRecommendSubjectListReq struct {
	Current  int64
	PageSize int64
}

type HomeRecommendSubject struct {
	Id              int64
	SubjectId       int64
	SubjectName     string
	RecommendStatus int
	Sort            int
}
type HomeRecommendSubjectListResp struct {
	Total int64
	List  []*HomeRecommendSubject
}

type HomeRecommendSubjectRepo interface {
	CreateHomeRecommendSubject(context.Context, *HomeRecommendSubject) error
	GetHomeRecommendSubject(ctx context.Context, id int64) error
	UpdateHomeRecommendSubject(context.Context, *HomeRecommendSubject) error
	ListHomeRecommendSubject(ctx context.Context, req *HomeRecommendSubjectListReq) (*HomeRecommendSubjectListResp, error)
	DeleteHomeRecommendSubject(ctx context.Context, id int64) error
}

type HomeRecommendSubjectUseCase struct {
	repo HomeRecommendSubjectRepo
	log  *log.Helper
}

func NewHomeRecommendSubjectUseCase(repo HomeRecommendSubjectRepo, logger log.Logger) *HomeRecommendSubjectUseCase {
	return &HomeRecommendSubjectUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (h HomeRecommendSubjectUseCase) CreateHomeRecommendSubject(ctx context.Context, subject *HomeRecommendSubject) error {
	panic("implement me")
}

func (h HomeRecommendSubjectUseCase) GetHomeRecommendSubject(ctx context.Context, id int64) error {
	panic("implement me")
}

func (h HomeRecommendSubjectUseCase) UpdateHomeRecommendSubject(ctx context.Context, subject *HomeRecommendSubject) error {
	panic("implement me")
}

func (h HomeRecommendSubjectUseCase) ListHomeRecommendSubject(ctx context.Context, req *HomeRecommendSubjectListReq) (*HomeRecommendSubjectListResp, error) {
	return h.repo.ListHomeRecommendSubject(ctx, req)
}

func (h HomeRecommendSubjectUseCase) DeleteHomeRecommendSubject(ctx context.Context, id int64) error {
	panic("implement me")
}
