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

type HomeRecommendSubjectRepo interface {
	CreateHomeRecommendSubject(context.Context, *HomeRecommendSubject) error
	GetHomeRecommendSubject(ctx context.Context, id int64) error
	UpdateHomeRecommendSubject(context.Context, *HomeRecommendSubject) error
	ListHomeRecommendSubject(ctx context.Context, req *HomeRecommendSubjectListReq) ([]*HomeRecommendSubject, error)
	DeleteHomeRecommendSubject(ctx context.Context, id int64) error
}

type HomeRecommendSubjectUseCase struct {
	repo HomeRecommendSubjectRepo
	log  *log.Helper
}

func NewHomeRecommendSubjectUseCase(repo HomeRecommendSubjectRepo, logger log.Logger) *HomeRecommendSubjectUseCase {
	return &HomeRecommendSubjectUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (r *HomeRecommendSubjectUseCase) CreateHomeRecommendSubject(ctx context.Context, user *HomeRecommendSubject) error {
	panic("implement me")
}

func (r *HomeRecommendSubjectUseCase) GetHomeRecommendSubject(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r *HomeRecommendSubjectUseCase) UpdateHomeRecommendSubject(ctx context.Context, user *HomeRecommendSubject) error {
	panic("implement me")
}

func (r *HomeRecommendSubjectUseCase) ListHomeRecommendSubject(ctx context.Context, pageNum, pageSize int64) ([]*HomeRecommendSubject, error) {
	panic("implement me")
}

func (r *HomeRecommendSubjectUseCase) DeleteHomeRecommendSubject(ctx context.Context, id int64) error {
	panic("implement me")
}
