package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/pay/internal/biz"

	pb "kratos-mall/api/pay/v1"
)

type PayService struct {
	pb.UnimplementedPayServer
	uc  *biz.GreeterUsecase
	log *log.Helper
}

func NewPayService(uc *biz.GreeterUsecase, logger log.Logger) *PayService {
	return &PayService{uc: uc, log: log.NewHelper(logger)}
}

func (s *PayService) AppUnifiedOrder(ctx context.Context, req *pb.UnifiedOrderReq) (*pb.UnifiedOrderResp, error) {
	return &pb.UnifiedOrderResp{}, nil
}
func (s *PayService) H5UnifiedOrder(ctx context.Context, req *pb.UnifiedOrderReq) (*pb.H5UnifiedOrderResp, error) {
	return &pb.H5UnifiedOrderResp{}, nil
}
func (s *PayService) JsUnifiedOrder(ctx context.Context, req *pb.UnifiedOrderReq) (*pb.UnifiedOrderResp, error) {
	return &pb.UnifiedOrderResp{}, nil
}
func (s *PayService) OrderQuery(ctx context.Context, req *pb.OrderQueryReq) (*pb.OrderQueryResp, error) {
	return &pb.OrderQueryResp{}, nil
}
