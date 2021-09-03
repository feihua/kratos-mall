package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "kratos-mall/api/front/admin/v1"
	"kratos-mall/app/front/admin/internal/biz"
)

type GreeterService struct {
	v1.UnimplementedGreeterServer
	uc  *biz.GreeterUsecase
	log *log.Helper
}

func NewGreeterService(uc *biz.GreeterUsecase, logger log.Logger) *GreeterService {
	return &GreeterService{uc: uc, log: log.NewHelper(logger)}
}

func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	s.log.WithContext(ctx).Infof("SayHello Received: %v", in.GetName())

	if in.GetName() == "error" {
		return nil, v1.ErrorUserNotFound("user not found: %s", in.GetName())
	}

	s.uc.Create(ctx, &biz.Greeter{
		Hello: in.GetName(),
	})
	return &v1.HelloReply{Message: "Hello " + in.GetName()}, nil
}
