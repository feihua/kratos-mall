package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/front/admin/internal/biz"
	"kratos-mall/app/front/admin/internal/biz/pay"

	pb "kratos-mall/api/front/admin/v1"
)

type PayService struct {
	pb.UnimplementedPayServer
	merchantUseCase *pay.MerchantUseCase
	recordUseCase   *pay.RecordUseCase
	log             *log.Helper
}

func NewPayService(uc *biz.GreeterUsecase, logger log.Logger, merchantUseCase *pay.MerchantUseCase,
	recordUseCase *pay.RecordUseCase) *PayService {
	return &PayService{log: log.NewHelper(logger),
		merchantUseCase: merchantUseCase,
		recordUseCase:   recordUseCase}
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
