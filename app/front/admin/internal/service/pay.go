package service

import (
	"context"

	pb "kratos-mall/api/front/admin/v1"
)

type PayService struct {
	pb.UnimplementedPayServer
}

func NewPayService() *PayService {
	return &PayService{}
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
