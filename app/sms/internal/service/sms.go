package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/sms/internal/biz"

	pb "kratos-mall/api/sms/v1"
)

type SmsService struct {
	pb.UnimplementedSmsServer
	uc                           *biz.GreeterUsecase
	couponUseCase                *biz.CouponUseCase
	couponHistoryUseCase         *biz.CouponHistoryUseCase
	flashPromotionUseCase        *biz.FlashPromotionUseCase
	flashPromotionHistoryUseCase *biz.FlashPromotionHistoryUseCase
	flashPromotionSessionUseCase *biz.FlashPromotionSessionUseCase
	homeAdvertiseUseCase         *biz.HomeAdvertiseUseCase
	homeBrandUseCase             *biz.HomeBrandUseCase
	homeNewProductUseCase        *biz.HomeNewProductUseCase
	homeRecommendProductUseCase  *biz.HomeRecommendProductUseCase
	homeRecommendSubjectUseCase  *biz.HomeRecommendSubjectUseCase
	log                          *log.Helper
}

func NewSmsService(uc *biz.GreeterUsecase, logger log.Logger,
	couponUseCase *biz.CouponUseCase,
	couponHistoryUseCase *biz.CouponHistoryUseCase,
	flashPromotionUseCase *biz.FlashPromotionUseCase,
	flashPromotionHistoryUseCase *biz.FlashPromotionHistoryUseCase,
	flashPromotionSessionUseCase *biz.FlashPromotionSessionUseCase,
	homeAdvertiseUseCase *biz.HomeAdvertiseUseCase,
	homeBrandUseCase *biz.HomeBrandUseCase,
	homeNewProductUseCase *biz.HomeNewProductUseCase,
	homeRecommendProductUseCase *biz.HomeRecommendProductUseCase,
	homeRecommendSubjectUseCase *biz.HomeRecommendSubjectUseCase) *SmsService {
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
	return &pb.CouponListResp{}, nil
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
	return &pb.CouponHistoryListResp{}, nil
}
func (s *SmsService) CouponHistoryUpdate(ctx context.Context, req *pb.CouponHistoryUpdateReq) (*pb.CouponHistoryUpdateResp, error) {
	return &pb.CouponHistoryUpdateResp{}, nil
}
func (s *SmsService) CouponHistoryDelete(ctx context.Context, req *pb.CouponHistoryDeleteReq) (*pb.CouponHistoryDeleteResp, error) {
	return &pb.CouponHistoryDeleteResp{}, nil
}
func (s *SmsService) CouponProductCategoryRelationAdd(ctx context.Context, req *pb.CouponProductCategoryRelationAddReq) (*pb.CouponProductCategoryRelationAddResp, error) {
	return &pb.CouponProductCategoryRelationAddResp{}, nil
}
func (s *SmsService) CouponProductCategoryRelationList(ctx context.Context, req *pb.CouponProductCategoryRelationListReq) (*pb.CouponProductCategoryRelationListResp, error) {
	return &pb.CouponProductCategoryRelationListResp{}, nil
}
func (s *SmsService) CouponProductCategoryRelationUpdate(ctx context.Context, req *pb.CouponProductCategoryRelationUpdateReq) (*pb.CouponProductCategoryRelationUpdateResp, error) {
	return &pb.CouponProductCategoryRelationUpdateResp{}, nil
}
func (s *SmsService) CouponProductCategoryRelationDelete(ctx context.Context, req *pb.CouponProductCategoryRelationDeleteReq) (*pb.CouponProductCategoryRelationDeleteResp, error) {
	return &pb.CouponProductCategoryRelationDeleteResp{}, nil
}
func (s *SmsService) CouponProductRelationAdd(ctx context.Context, req *pb.CouponProductRelationAddReq) (*pb.CouponProductRelationAddResp, error) {
	return &pb.CouponProductRelationAddResp{}, nil
}
func (s *SmsService) CouponProductRelationList(ctx context.Context, req *pb.CouponProductRelationListReq) (*pb.CouponProductRelationListResp, error) {
	return &pb.CouponProductRelationListResp{}, nil
}
func (s *SmsService) CouponProductRelationUpdate(ctx context.Context, req *pb.CouponProductRelationUpdateReq) (*pb.CouponProductRelationUpdateResp, error) {
	return &pb.CouponProductRelationUpdateResp{}, nil
}
func (s *SmsService) CouponProductRelationDelete(ctx context.Context, req *pb.CouponProductRelationDeleteReq) (*pb.CouponProductRelationDeleteResp, error) {
	return &pb.CouponProductRelationDeleteResp{}, nil
}
func (s *SmsService) FlashPromotionLogAdd(ctx context.Context, req *pb.FlashPromotionLogAddReq) (*pb.FlashPromotionLogAddResp, error) {
	return &pb.FlashPromotionLogAddResp{}, nil
}
func (s *SmsService) FlashPromotionLogList(ctx context.Context, req *pb.FlashPromotionLogListReq) (*pb.FlashPromotionLogListResp, error) {
	return &pb.FlashPromotionLogListResp{}, nil
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
	return &pb.FlashPromotionListResp{}, nil
}
func (s *SmsService) FlashPromotionUpdate(ctx context.Context, req *pb.FlashPromotionUpdateReq) (*pb.FlashPromotionUpdateResp, error) {
	return &pb.FlashPromotionUpdateResp{}, nil
}
func (s *SmsService) FlashPromotionDelete(ctx context.Context, req *pb.FlashPromotionDeleteReq) (*pb.FlashPromotionDeleteResp, error) {
	return &pb.FlashPromotionDeleteResp{}, nil
}
func (s *SmsService) FlashPromotionProductRelationAdd(ctx context.Context, req *pb.FlashPromotionProductRelationAddReq) (*pb.FlashPromotionProductRelationAddResp, error) {
	return &pb.FlashPromotionProductRelationAddResp{}, nil
}
func (s *SmsService) FlashPromotionProductRelationList(ctx context.Context, req *pb.FlashPromotionProductRelationListReq) (*pb.FlashPromotionProductRelationListResp, error) {
	return &pb.FlashPromotionProductRelationListResp{}, nil
}
func (s *SmsService) FlashPromotionProductRelationUpdate(ctx context.Context, req *pb.FlashPromotionProductRelationUpdateReq) (*pb.FlashPromotionProductRelationUpdateResp, error) {
	return &pb.FlashPromotionProductRelationUpdateResp{}, nil
}
func (s *SmsService) FlashPromotionProductRelationDelete(ctx context.Context, req *pb.FlashPromotionProductRelationDeleteReq) (*pb.FlashPromotionProductRelationDeleteResp, error) {
	return &pb.FlashPromotionProductRelationDeleteResp{}, nil
}
func (s *SmsService) FlashPromotionSessionAdd(ctx context.Context, req *pb.FlashPromotionSessionAddReq) (*pb.FlashPromotionSessionAddResp, error) {
	return &pb.FlashPromotionSessionAddResp{}, nil
}
func (s *SmsService) FlashPromotionSessionList(ctx context.Context, req *pb.FlashPromotionSessionListReq) (*pb.FlashPromotionSessionListResp, error) {
	return &pb.FlashPromotionSessionListResp{}, nil
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
	return &pb.HomeAdvertiseListResp{}, nil
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
	return &pb.HomeBrandListResp{}, nil
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
	return &pb.HomeNewProductListResp{}, nil
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
	return &pb.HomeRecommendProductListResp{}, nil
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
	return &pb.HomeRecommendSubjectListResp{}, nil
}
func (s *SmsService) HomeRecommendSubjectUpdate(ctx context.Context, req *pb.HomeRecommendSubjectUpdateReq) (*pb.HomeRecommendSubjectUpdateResp, error) {
	return &pb.HomeRecommendSubjectUpdateResp{}, nil
}
func (s *SmsService) HomeRecommendSubjectDelete(ctx context.Context, req *pb.HomeRecommendSubjectDeleteReq) (*pb.HomeRecommendSubjectDeleteResp, error) {
	return &pb.HomeRecommendSubjectDeleteResp{}, nil
}
