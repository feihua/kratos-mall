package service

import (
	"context"
	"github.com/feihua/kratos-mall/app/front/admin/internal/biz"
	"github.com/feihua/kratos-mall/app/front/admin/internal/biz/sms"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"

	pb "github.com/feihua/kratos-mall/api/front/admin/v1"
)

type SmsService struct {
	pb.UnimplementedSmsServer
	uc                           *biz.GreeterUsecase
	couponUseCase                *sms.CouponUseCase
	couponHistoryUseCase         *sms.CouponHistoryUseCase
	flashPromotionUseCase        *sms.FlashPromotionUseCase
	flashPromotionHistoryUseCase *sms.FlashPromotionHistoryUseCase
	flashPromotionSessionUseCase *sms.FlashPromotionSessionUseCase
	homeAdvertiseUseCase         *sms.HomeAdvertiseUseCase
	homeBrandUseCase             *sms.HomeBrandUseCase
	homeNewProductUseCase        *sms.HomeNewProductUseCase
	homeRecommendProductUseCase  *sms.HomeRecommendProductUseCase
	homeRecommendSubjectUseCase  *sms.HomeRecommendSubjectUseCase
	log                          *log.Helper
}

func NewSmsService(uc *biz.GreeterUsecase, logger log.Logger,
	couponUseCase *sms.CouponUseCase,
	couponHistoryUseCase *sms.CouponHistoryUseCase,
	flashPromotionUseCase *sms.FlashPromotionUseCase,
	flashPromotionHistoryUseCase *sms.FlashPromotionHistoryUseCase,
	flashPromotionSessionUseCase *sms.FlashPromotionSessionUseCase,
	homeAdvertiseUseCase *sms.HomeAdvertiseUseCase,
	homeBrandUseCase *sms.HomeBrandUseCase,
	homeNewProductUseCase *sms.HomeNewProductUseCase,
	homeRecommendProductUseCase *sms.HomeRecommendProductUseCase,
	homeRecommendSubjectUseCase *sms.HomeRecommendSubjectUseCase) *SmsService {
	return &SmsService{uc: uc, log: log.NewHelper(logger),
		couponUseCase:                couponUseCase,
		couponHistoryUseCase:         couponHistoryUseCase,
		flashPromotionHistoryUseCase: flashPromotionHistoryUseCase,
		flashPromotionSessionUseCase: flashPromotionSessionUseCase,
		flashPromotionUseCase:        flashPromotionUseCase,
		homeAdvertiseUseCase:         homeAdvertiseUseCase,
		homeBrandUseCase:             homeBrandUseCase,
		homeNewProductUseCase:        homeNewProductUseCase,
		homeRecommendProductUseCase:  homeRecommendProductUseCase,
		homeRecommendSubjectUseCase:  homeRecommendSubjectUseCase}
}

func (s *SmsService) CouponAdd(ctx context.Context, req *pb.CouponAddReq) (*pb.CouponAddResp, error) {
	return &pb.CouponAddResp{}, nil
}
func (s *SmsService) CouponList(ctx context.Context, req *pb.CouponListReq) (*pb.CouponListResp, error) {
	listResp, _ := s.couponUseCase.ListCoupon(ctx, &sms.CouponListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.CouponListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.CouponListResp{
		Code:     "000000",
		Message:  "查询订单信息成功",
		Current:  int32(req.Current),
		PageSize: int32(req.PageSize),
		Total:    int32(listResp.Total),
		Data:     list,
		Success:  true,
	}, nil
}
func (s *SmsService) CouponUpdate(ctx context.Context, req *pb.CouponUpdateReq) (*pb.CouponUpdateResp, error) {
	return &pb.CouponUpdateResp{}, nil
}
func (s *SmsService) CouponDelete(ctx context.Context, req *pb.CouponDeleteReq) (*pb.CouponDeleteResp, error) {
	return &pb.CouponDeleteResp{}, nil
}
func (s *SmsService) CouponHistoryAdd(ctx context.Context, req *pb.CouponHistoryAddReq) (*pb.CouponHistoryAddResp, error) {
	return &pb.CouponHistoryAddResp{}, nil
}
func (s *SmsService) CouponHistoryList(ctx context.Context, req *pb.CouponHistoryListReq) (*pb.CouponHistoryListResp, error) {
	listResp, _ := s.couponHistoryUseCase.ListCouponHistory(ctx, &sms.CouponHistoryListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.CouponHistoryListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.CouponHistoryListResp{
		Code:     "000000",
		Message:  "查询订单信息成功",
		Current:  int32(req.Current),
		PageSize: int32(req.PageSize),
		Total:    int32(listResp.Total),
		Data:     list,
		Success:  true,
	}, nil
}
func (s *SmsService) CouponHistoryUpdate(ctx context.Context, req *pb.CouponHistoryUpdateReq) (*pb.CouponHistoryUpdateResp, error) {
	return &pb.CouponHistoryUpdateResp{}, nil
}
func (s *SmsService) CouponHistoryDelete(ctx context.Context, req *pb.CouponHistoryDeleteReq) (*pb.CouponHistoryDeleteResp, error) {
	return &pb.CouponHistoryDeleteResp{}, nil
}
func (s *SmsService) FlashPromotionLogAdd(ctx context.Context, req *pb.FlashPromotionLogAddReq) (*pb.FlashPromotionLogAddResp, error) {
	return &pb.FlashPromotionLogAddResp{}, nil
}
func (s *SmsService) FlashPromotionLogList(ctx context.Context, req *pb.FlashPromotionLogListReq) (*pb.FlashPromotionLogListResp, error) {
	listResp, _ := s.flashPromotionHistoryUseCase.ListFlashPromotionHistory(ctx, &sms.FlashPromotionHistoryListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.FlashPromotionLogListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.FlashPromotionLogListResp{
		Code:     "000000",
		Message:  "查询订单信息成功",
		Current:  int32(req.Current),
		PageSize: int32(req.PageSize),
		Total:    int32(listResp.Total),
		Data:     list,
		Success:  true,
	}, nil
}
func (s *SmsService) FlashPromotionLogUpdate(ctx context.Context, req *pb.FlashPromotionLogUpdateReq) (*pb.FlashPromotionLogUpdateResp, error) {
	return &pb.FlashPromotionLogUpdateResp{}, nil
}
func (s *SmsService) FlashPromotionLogDelete(ctx context.Context, req *pb.FlashPromotionLogDeleteReq) (*pb.FlashPromotionLogDeleteResp, error) {
	return &pb.FlashPromotionLogDeleteResp{}, nil
}
func (s *SmsService) FlashPromotionAdd(ctx context.Context, req *pb.FlashPromotionAddReq) (*pb.FlashPromotionAddResp, error) {
	return &pb.FlashPromotionAddResp{}, nil
}
func (s *SmsService) FlashPromotionList(ctx context.Context, req *pb.FlashPromotionListReq) (*pb.FlashPromotionListResp, error) {
	listResp, _ := s.flashPromotionUseCase.ListFlashPromotion(ctx, &sms.FlashPromotionListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.FlashPromotionListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.FlashPromotionListResp{
		Code:     "000000",
		Message:  "查询订单信息成功",
		Current:  int32(req.Current),
		PageSize: int32(req.PageSize),
		Total:    int32(listResp.Total),
		Data:     list,
		Success:  true,
	}, nil
}
func (s *SmsService) FlashPromotionUpdate(ctx context.Context, req *pb.FlashPromotionUpdateReq) (*pb.FlashPromotionUpdateResp, error) {
	return &pb.FlashPromotionUpdateResp{}, nil
}
func (s *SmsService) FlashPromotionDelete(ctx context.Context, req *pb.FlashPromotionDeleteReq) (*pb.FlashPromotionDeleteResp, error) {
	return &pb.FlashPromotionDeleteResp{}, nil
}
func (s *SmsService) FlashPromotionSessionAdd(ctx context.Context, req *pb.FlashPromotionSessionAddReq) (*pb.FlashPromotionSessionAddResp, error) {
	return &pb.FlashPromotionSessionAddResp{}, nil
}
func (s *SmsService) FlashPromotionSessionList(ctx context.Context, req *pb.FlashPromotionSessionListReq) (*pb.FlashPromotionSessionListResp, error) {
	listResp, _ := s.flashPromotionSessionUseCase.ListFlashPromotionSession(ctx, &sms.FlashPromotionSessionListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.FlashPromotionSessionListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.FlashPromotionSessionListResp{
		Code:     "000000",
		Message:  "查询订单信息成功",
		Current:  int32(req.Current),
		PageSize: int32(req.PageSize),
		Total:    int32(listResp.Total),
		Data:     list,
		Success:  true,
	}, nil
}
func (s *SmsService) FlashPromotionSessionUpdate(ctx context.Context, req *pb.FlashPromotionSessionUpdateReq) (*pb.FlashPromotionSessionUpdateResp, error) {
	return &pb.FlashPromotionSessionUpdateResp{}, nil
}
func (s *SmsService) FlashPromotionSessionDelete(ctx context.Context, req *pb.FlashPromotionSessionDeleteReq) (*pb.FlashPromotionSessionDeleteResp, error) {
	return &pb.FlashPromotionSessionDeleteResp{}, nil
}
func (s *SmsService) HomeAdvertiseAdd(ctx context.Context, req *pb.HomeAdvertiseAddReq) (*pb.HomeAdvertiseAddResp, error) {
	return &pb.HomeAdvertiseAddResp{}, nil
}
func (s *SmsService) HomeAdvertiseList(ctx context.Context, req *pb.HomeAdvertiseListReq) (*pb.HomeAdvertiseListResp, error) {
	listResp, _ := s.homeAdvertiseUseCase.ListHomeAdvertise(ctx, &sms.HomeAdvertiseListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.HomeAdvertiseListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.HomeAdvertiseListResp{
		Code:     "000000",
		Message:  "查询订单信息成功",
		Current:  int32(req.Current),
		PageSize: int32(req.PageSize),
		Total:    int32(listResp.Total),
		Data:     list,
		Success:  true,
	}, nil
}
func (s *SmsService) HomeAdvertiseUpdate(ctx context.Context, req *pb.HomeAdvertiseUpdateReq) (*pb.HomeAdvertiseUpdateResp, error) {
	return &pb.HomeAdvertiseUpdateResp{}, nil
}
func (s *SmsService) HomeAdvertiseDelete(ctx context.Context, req *pb.HomeAdvertiseDeleteReq) (*pb.HomeAdvertiseDeleteResp, error) {
	return &pb.HomeAdvertiseDeleteResp{}, nil
}
func (s *SmsService) HomeBrandAdd(ctx context.Context, req *pb.HomeBrandAddReq) (*pb.HomeBrandAddResp, error) {
	return &pb.HomeBrandAddResp{}, nil
}
func (s *SmsService) HomeBrandList(ctx context.Context, req *pb.HomeBrandListReq) (*pb.HomeBrandListResp, error) {
	listResp, _ := s.homeBrandUseCase.ListHomeBrand(ctx, &sms.HomeBrandListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.HomeBrandListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.HomeBrandListResp{
		Code:     "000000",
		Message:  "查询订单信息成功",
		Current:  int32(req.Current),
		PageSize: int32(req.PageSize),
		Total:    int32(listResp.Total),
		Data:     list,
		Success:  true,
	}, nil
}
func (s *SmsService) HomeBrandUpdate(ctx context.Context, req *pb.HomeBrandUpdateReq) (*pb.HomeBrandUpdateResp, error) {
	return &pb.HomeBrandUpdateResp{}, nil
}
func (s *SmsService) HomeBrandDelete(ctx context.Context, req *pb.HomeBrandDeleteReq) (*pb.HomeBrandDeleteResp, error) {
	return &pb.HomeBrandDeleteResp{}, nil
}
func (s *SmsService) HomeNewProductAdd(ctx context.Context, req *pb.HomeNewProductAddReq) (*pb.HomeNewProductAddResp, error) {
	return &pb.HomeNewProductAddResp{}, nil
}
func (s *SmsService) HomeNewProductList(ctx context.Context, req *pb.HomeNewProductListReq) (*pb.HomeNewProductListResp, error) {
	listResp, _ := s.homeNewProductUseCase.ListHomeNewProduct(ctx, &sms.HomeNewProductListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.HomeNewProductListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.HomeNewProductListResp{
		Code:     "000000",
		Message:  "查询订单信息成功",
		Current:  int32(req.Current),
		PageSize: int32(req.PageSize),
		Total:    int32(listResp.Total),
		Data:     list,
		Success:  true,
	}, nil
}
func (s *SmsService) HomeNewProductUpdate(ctx context.Context, req *pb.HomeNewProductUpdateReq) (*pb.HomeNewProductUpdateResp, error) {
	return &pb.HomeNewProductUpdateResp{}, nil
}
func (s *SmsService) HomeNewProductDelete(ctx context.Context, req *pb.HomeNewProductDeleteReq) (*pb.HomeNewProductDeleteResp, error) {
	return &pb.HomeNewProductDeleteResp{}, nil
}
func (s *SmsService) HomeRecommendProductAdd(ctx context.Context, req *pb.HomeRecommendProductAddReq) (*pb.HomeRecommendProductAddResp, error) {
	return &pb.HomeRecommendProductAddResp{}, nil
}
func (s *SmsService) HomeRecommendProductList(ctx context.Context, req *pb.HomeRecommendProductListReq) (*pb.HomeRecommendProductListResp, error) {
	listResp, _ := s.homeRecommendProductUseCase.ListHomeRecommendProduct(ctx, &sms.HomeRecommendProductListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.HomeRecommendProductListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.HomeRecommendProductListResp{
		Code:     "000000",
		Message:  "查询订单信息成功",
		Current:  int32(req.Current),
		PageSize: int32(req.PageSize),
		Total:    int32(listResp.Total),
		Data:     list,
		Success:  true,
	}, nil
}
func (s *SmsService) HomeRecommendProductUpdate(ctx context.Context, req *pb.HomeRecommendProductUpdateReq) (*pb.HomeRecommendProductUpdateResp, error) {
	return &pb.HomeRecommendProductUpdateResp{}, nil
}
func (s *SmsService) HomeRecommendProductDelete(ctx context.Context, req *pb.HomeRecommendProductDeleteReq) (*pb.HomeRecommendProductDeleteResp, error) {
	return &pb.HomeRecommendProductDeleteResp{}, nil
}
func (s *SmsService) HomeRecommendSubjectAdd(ctx context.Context, req *pb.HomeRecommendSubjectAddReq) (*pb.HomeRecommendSubjectAddResp, error) {
	return &pb.HomeRecommendSubjectAddResp{}, nil
}
func (s *SmsService) HomeRecommendSubjectList(ctx context.Context, req *pb.HomeRecommendSubjectListReq) (*pb.HomeRecommendSubjectListResp, error) {
	listResp, _ := s.homeRecommendSubjectUseCase.ListHomeRecommendSubject(ctx, &sms.HomeRecommendSubjectListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.HomeRecommendSubjectListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.HomeRecommendSubjectListResp{
		Code:     "000000",
		Message:  "查询订单信息成功",
		Current:  int32(req.Current),
		PageSize: int32(req.PageSize),
		Total:    int32(listResp.Total),
		Data:     list,
		Success:  true,
	}, nil
}
func (s *SmsService) HomeRecommendSubjectUpdate(ctx context.Context, req *pb.HomeRecommendSubjectUpdateReq) (*pb.HomeRecommendSubjectUpdateResp, error) {
	return &pb.HomeRecommendSubjectUpdateResp{}, nil
}
func (s *SmsService) HomeRecommendSubjectDelete(ctx context.Context, req *pb.HomeRecommendSubjectDeleteReq) (*pb.HomeRecommendSubjectDeleteResp, error) {
	return &pb.HomeRecommendSubjectDeleteResp{}, nil
}
