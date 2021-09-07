package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/oms/internal/biz"

	pb "kratos-mall/api/oms/v1"
)

type OmsService struct {
	pb.UnimplementedOmsServer
	uc                    *biz.GreeterUsecase
	cartItemUseCase       *biz.CartItemUseCase
	companyAddressUseCase *biz.CompanyAddressUseCase
	operateHistoryUseCase *biz.OperateHistoryUseCase
	orderUseCase          *biz.OrderUseCase
	returnApplyUseCase    *biz.ReturnApplyUseCase
	returnReasonUseCase   *biz.ReturnReasonUseCase
	settingUseCase        *biz.SettingUseCase
	log                   *log.Helper
}

func NewOmsService(uc *biz.GreeterUsecase, logger log.Logger,
	cartItemUseCase *biz.CartItemUseCase,
	companyAddressUseCase *biz.CompanyAddressUseCase,
	operateHistoryUseCase *biz.OperateHistoryUseCase,
	orderUseCase *biz.OrderUseCase,
	returnApplyUseCase *biz.ReturnApplyUseCase,
	returnReasonUseCase *biz.ReturnReasonUseCase,
	settingUseCase *biz.SettingUseCase) *OmsService {
	return &OmsService{uc: uc, log: log.NewHelper(logger),
		cartItemUseCase:       cartItemUseCase,
		companyAddressUseCase: companyAddressUseCase,
		operateHistoryUseCase: operateHistoryUseCase,
		orderUseCase:          orderUseCase,
		returnApplyUseCase:    returnApplyUseCase,
		returnReasonUseCase:   returnReasonUseCase,
		settingUseCase:        settingUseCase}
}

func (s *OmsService) OrderAdd(ctx context.Context, req *pb.OrderAddReq) (*pb.OrderAddResp, error) {
	return &pb.OrderAddResp{}, nil
}
func (s *OmsService) OrderList(ctx context.Context, req *pb.OrderListReq) (*pb.OrderListResp, error) {
	return &pb.OrderListResp{}, nil
}
func (s *OmsService) OrderUpdate(ctx context.Context, req *pb.OrderUpdateReq) (*pb.OrderUpdateResp, error) {
	return &pb.OrderUpdateResp{}, nil
}
func (s *OmsService) OrderDelete(ctx context.Context, req *pb.OrderDeleteReq) (*pb.OrderDeleteResp, error) {
	return &pb.OrderDeleteResp{}, nil
}
func (s *OmsService) CartItemAdd(ctx context.Context, req *pb.CartItemAddReq) (*pb.CartItemAddResp, error) {
	return &pb.CartItemAddResp{}, nil
}
func (s *OmsService) CartItemList(ctx context.Context, req *pb.CartItemListReq) (*pb.CartItemListResp, error) {
	return &pb.CartItemListResp{}, nil
}
func (s *OmsService) CartItemUpdate(ctx context.Context, req *pb.CartItemUpdateReq) (*pb.CartItemUpdateResp, error) {
	return &pb.CartItemUpdateResp{}, nil
}
func (s *OmsService) CartItemDelete(ctx context.Context, req *pb.CartItemDeleteReq) (*pb.CartItemDeleteResp, error) {
	return &pb.CartItemDeleteResp{}, nil
}
func (s *OmsService) CompanyAddressAdd(ctx context.Context, req *pb.CompanyAddressAddReq) (*pb.CompanyAddressAddResp, error) {
	return &pb.CompanyAddressAddResp{}, nil
}
func (s *OmsService) CompanyAddressList(ctx context.Context, req *pb.CompanyAddressListReq) (*pb.CompanyAddressListResp, error) {
	return &pb.CompanyAddressListResp{}, nil
}
func (s *OmsService) CompanyAddressUpdate(ctx context.Context, req *pb.CompanyAddressUpdateReq) (*pb.CompanyAddressUpdateResp, error) {
	return &pb.CompanyAddressUpdateResp{}, nil
}
func (s *OmsService) CompanyAddressDelete(ctx context.Context, req *pb.CompanyAddressDeleteReq) (*pb.CompanyAddressDeleteResp, error) {
	return &pb.CompanyAddressDeleteResp{}, nil
}
func (s *OmsService) OrderItemAdd(ctx context.Context, req *pb.OrderItemAddReq) (*pb.OrderItemAddResp, error) {
	return &pb.OrderItemAddResp{}, nil
}
func (s *OmsService) OrderItemList(ctx context.Context, req *pb.OrderItemListReq) (*pb.OrderItemListResp, error) {
	return &pb.OrderItemListResp{}, nil
}
func (s *OmsService) OrderItemUpdate(ctx context.Context, req *pb.OrderItemUpdateReq) (*pb.OrderItemUpdateResp, error) {
	return &pb.OrderItemUpdateResp{}, nil
}
func (s *OmsService) OrderItemDelete(ctx context.Context, req *pb.OrderItemDeleteReq) (*pb.OrderItemDeleteResp, error) {
	return &pb.OrderItemDeleteResp{}, nil
}
func (s *OmsService) OrderOperateHistoryAdd(ctx context.Context, req *pb.OrderOperateHistoryAddReq) (*pb.OrderOperateHistoryAddResp, error) {
	return &pb.OrderOperateHistoryAddResp{}, nil
}
func (s *OmsService) OrderOperateHistoryList(ctx context.Context, req *pb.OrderOperateHistoryListReq) (*pb.OrderOperateHistoryListResp, error) {
	return &pb.OrderOperateHistoryListResp{}, nil
}
func (s *OmsService) OrderOperateHistoryUpdate(ctx context.Context, req *pb.OrderOperateHistoryUpdateReq) (*pb.OrderOperateHistoryUpdateResp, error) {
	return &pb.OrderOperateHistoryUpdateResp{}, nil
}
func (s *OmsService) OrderOperateHistoryDelete(ctx context.Context, req *pb.OrderOperateHistoryDeleteReq) (*pb.OrderOperateHistoryDeleteResp, error) {
	return &pb.OrderOperateHistoryDeleteResp{}, nil
}
func (s *OmsService) OrderReturnApplyAdd(ctx context.Context, req *pb.OrderReturnApplyAddReq) (*pb.OrderReturnApplyAddResp, error) {
	return &pb.OrderReturnApplyAddResp{}, nil
}
func (s *OmsService) OrderReturnApplyList(ctx context.Context, req *pb.OrderReturnApplyListReq) (*pb.OrderReturnApplyListResp, error) {
	return &pb.OrderReturnApplyListResp{}, nil
}
func (s *OmsService) OrderReturnApplyUpdate(ctx context.Context, req *pb.OrderReturnApplyUpdateReq) (*pb.OrderReturnApplyUpdateResp, error) {
	return &pb.OrderReturnApplyUpdateResp{}, nil
}
func (s *OmsService) OrderReturnApplyDelete(ctx context.Context, req *pb.OrderReturnApplyDeleteReq) (*pb.OrderReturnApplyDeleteResp, error) {
	return &pb.OrderReturnApplyDeleteResp{}, nil
}
func (s *OmsService) OrderReturnReasonAdd(ctx context.Context, req *pb.OrderReturnReasonAddReq) (*pb.OrderReturnReasonAddResp, error) {
	return &pb.OrderReturnReasonAddResp{}, nil
}
func (s *OmsService) OrderReturnReasonList(ctx context.Context, req *pb.OrderReturnReasonListReq) (*pb.OrderReturnReasonListResp, error) {
	return &pb.OrderReturnReasonListResp{}, nil
}
func (s *OmsService) OrderReturnReasonUpdate(ctx context.Context, req *pb.OrderReturnReasonUpdateReq) (*pb.OrderReturnReasonUpdateResp, error) {
	return &pb.OrderReturnReasonUpdateResp{}, nil
}
func (s *OmsService) OrderReturnReasonDelete(ctx context.Context, req *pb.OrderReturnReasonDeleteReq) (*pb.OrderReturnReasonDeleteResp, error) {
	return &pb.OrderReturnReasonDeleteResp{}, nil
}
func (s *OmsService) OrderSettingAdd(ctx context.Context, req *pb.OrderSettingAddReq) (*pb.OrderSettingAddResp, error) {
	return &pb.OrderSettingAddResp{}, nil
}
func (s *OmsService) OrderSettingList(ctx context.Context, req *pb.OrderSettingListReq) (*pb.OrderSettingListResp, error) {
	return &pb.OrderSettingListResp{}, nil
}
func (s *OmsService) OrderSettingUpdate(ctx context.Context, req *pb.OrderSettingUpdateReq) (*pb.OrderSettingUpdateResp, error) {
	return &pb.OrderSettingUpdateResp{}, nil
}
func (s *OmsService) OrderSettingDelete(ctx context.Context, req *pb.OrderSettingDeleteReq) (*pb.OrderSettingDeleteResp, error) {
	return &pb.OrderSettingDeleteResp{}, nil
}
