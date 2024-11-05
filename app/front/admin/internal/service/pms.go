package service

import (
	"context"
	"github.com/feihua/kratos-mall/app/front/admin/internal/biz"
	"github.com/feihua/kratos-mall/app/front/admin/internal/biz/pms"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"

	pb "github.com/feihua/kratos-mall/api/front/admin/v1"
)

type PmsService struct {
	pb.UnimplementedPmsServer
	uc                    *biz.GreeterUsecase
	brandUseCase          *pms.BrandUseCase
	categoryUseCase       *pms.CategoryUseCase
	commentUseCase        *pms.CommentUseCase
	commentReplayUseCase  *pms.CommentReplayUseCase
	feightTemplateUseCase *pms.FeightTemplateUseCase
	memberPriceUseCase    *pms.MemberPriceUseCase
	operateHistoryUseCase *pms.OperateHistoryUseCase
	productUseCase        *pms.ProductUseCase
	skuStockUseCase       *pms.SkuStockUseCase
	vertifyRecordUseCase  *pms.VertifyRecordUseCase
	log                   *log.Helper
}

func NewPmsService(uc *biz.GreeterUsecase, logger log.Logger,
	brandUseCase *pms.BrandUseCase,
	categoryUseCase *pms.CategoryUseCase,
	commentUseCase *pms.CommentUseCase,
	commentReplayUseCase *pms.CommentReplayUseCase,
	feightTemplateUseCase *pms.FeightTemplateUseCase,
	memberPriceUseCase *pms.MemberPriceUseCase,
	operateHistoryUseCase *pms.OperateHistoryUseCase,
	productUseCase *pms.ProductUseCase,
	skuStockUseCase *pms.SkuStockUseCase,
	vertifyRecordUseCase *pms.VertifyRecordUseCase) *PmsService {
	return &PmsService{uc: uc, log: log.NewHelper(logger),
		brandUseCase:          brandUseCase,
		categoryUseCase:       categoryUseCase,
		commentUseCase:        commentUseCase,
		commentReplayUseCase:  commentReplayUseCase,
		feightTemplateUseCase: feightTemplateUseCase,
		memberPriceUseCase:    memberPriceUseCase,
		operateHistoryUseCase: operateHistoryUseCase,
		productUseCase:        productUseCase,
		skuStockUseCase:       skuStockUseCase,
		vertifyRecordUseCase:  vertifyRecordUseCase}
}

func (s *PmsService) ProductAdd(ctx context.Context, req *pb.ProductAddReq) (*pb.ProductAddResp, error) {
	return &pb.ProductAddResp{}, nil
}
func (s *PmsService) ProductList(ctx context.Context, req *pb.ProductListReq) (*pb.ProductListResp, error) {
	listResp, _ := s.productUseCase.ListProduct(ctx, &pms.ProductListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.ProductListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.ProductListResp{
		Code:     "000000",
		Message:  "查询订单信息成功",
		Current:  int32(req.Current),
		PageSize: int32(req.PageSize),
		Total:    int32(listResp.Total),
		Data:     list,
		Success:  true,
	}, nil
}
func (s *PmsService) ProductUpdate(ctx context.Context, req *pb.ProductUpdateReq) (*pb.ProductUpdateResp, error) {
	return &pb.ProductUpdateResp{}, nil
}
func (s *PmsService) ProductDelete(ctx context.Context, req *pb.ProductDeleteReq) (*pb.ProductDeleteResp, error) {
	return &pb.ProductDeleteResp{}, nil
}
func (s *PmsService) BrandAdd(ctx context.Context, req *pb.BrandAddReq) (*pb.BrandAddResp, error) {
	return &pb.BrandAddResp{}, nil
}
func (s *PmsService) BrandList(ctx context.Context, req *pb.BrandListReq) (*pb.BrandListResp, error) {
	listResp, _ := s.brandUseCase.ListBrand(ctx, &pms.BrandListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.BrandListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.BrandListResp{
		Code:     "000000",
		Message:  "查询订单信息成功",
		Current:  int32(req.Current),
		PageSize: int32(req.PageSize),
		Total:    int32(listResp.Total),
		Data:     list,
		Success:  true,
	}, nil
}
func (s *PmsService) BrandUpdate(ctx context.Context, req *pb.BrandUpdateReq) (*pb.BrandUpdateResp, error) {
	return &pb.BrandUpdateResp{}, nil
}
func (s *PmsService) BrandDelete(ctx context.Context, req *pb.BrandDeleteReq) (*pb.BrandDeleteResp, error) {
	return &pb.BrandDeleteResp{}, nil
}
func (s *PmsService) CommentAdd(ctx context.Context, req *pb.CommentAddReq) (*pb.CommentAddResp, error) {
	return &pb.CommentAddResp{}, nil
}
func (s *PmsService) CommentList(ctx context.Context, req *pb.CommentListReq) (*pb.CommentListResp, error) {
	listResp, _ := s.commentUseCase.ListComment(ctx, &pms.CommentListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.CommentListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.CommentListResp{
		Code:     "000000",
		Message:  "查询订单信息成功",
		Current:  int32(req.Current),
		PageSize: int32(req.PageSize),
		Total:    int32(listResp.Total),
		Data:     list,
		Success:  true,
	}, nil
}
func (s *PmsService) CommentUpdate(ctx context.Context, req *pb.CommentUpdateReq) (*pb.CommentUpdateResp, error) {
	return &pb.CommentUpdateResp{}, nil
}
func (s *PmsService) CommentDelete(ctx context.Context, req *pb.CommentDeleteReq) (*pb.CommentDeleteResp, error) {
	return &pb.CommentDeleteResp{}, nil
}
func (s *PmsService) FeightTemplateAdd(ctx context.Context, req *pb.FeightTemplateAddReq) (*pb.FeightTemplateAddResp, error) {
	return &pb.FeightTemplateAddResp{}, nil
}
func (s *PmsService) FeightTemplateList(ctx context.Context, req *pb.FeightTemplateListReq) (*pb.FeightTemplateListResp, error) {
	listResp, _ := s.feightTemplateUseCase.ListFeightTemplate(ctx, &pms.FeightTemplateListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.FeightTemplateListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.FeightTemplateListResp{
		Code:     "000000",
		Message:  "查询订单信息成功",
		Current:  int32(req.Current),
		PageSize: int32(req.PageSize),
		Total:    int32(listResp.Total),
		Data:     list,
		Success:  true,
	}, nil
}
func (s *PmsService) FeightTemplateUpdate(ctx context.Context, req *pb.FeightTemplateUpdateReq) (*pb.FeightTemplateUpdateResp, error) {
	return &pb.FeightTemplateUpdateResp{}, nil
}
func (s *PmsService) FeightTemplateDelete(ctx context.Context, req *pb.FeightTemplateDeleteReq) (*pb.FeightTemplateDeleteResp, error) {
	return &pb.FeightTemplateDeleteResp{}, nil
}
func (s *PmsService) MemberPriceAdd(ctx context.Context, req *pb.MemberPriceAddReq) (*pb.MemberPriceAddResp, error) {
	return &pb.MemberPriceAddResp{}, nil
}
func (s *PmsService) MemberPriceList(ctx context.Context, req *pb.MemberPriceListReq) (*pb.MemberPriceListResp, error) {
	listResp, _ := s.memberPriceUseCase.ListMemberPrice(ctx, &pms.MemberPriceListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.MemberPriceListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.MemberPriceListResp{
		Code:     "000000",
		Message:  "查询订单信息成功",
		Current:  int32(req.Current),
		PageSize: int32(req.PageSize),
		Total:    int32(listResp.Total),
		Data:     list,
		Success:  true,
	}, nil
}
func (s *PmsService) MemberPriceUpdate(ctx context.Context, req *pb.MemberPriceUpdateReq) (*pb.MemberPriceUpdateResp, error) {
	return &pb.MemberPriceUpdateResp{}, nil
}
func (s *PmsService) MemberPriceDelete(ctx context.Context, req *pb.MemberPriceDeleteReq) (*pb.MemberPriceDeleteResp, error) {
	return &pb.MemberPriceDeleteResp{}, nil
}
func (s *PmsService) ProductCategoryAdd(ctx context.Context, req *pb.ProductCategoryAddReq) (*pb.ProductCategoryAddResp, error) {
	return &pb.ProductCategoryAddResp{}, nil
}
func (s *PmsService) ProductCategoryList(ctx context.Context, req *pb.ProductCategoryListReq) (*pb.ProductCategoryListResp, error) {
	listResp, _ := s.categoryUseCase.ListCategory(ctx, &pms.CategoryListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.ProductCategoryListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.ProductCategoryListResp{
		Code:     "000000",
		Message:  "查询订单信息成功",
		Current:  int32(req.Current),
		PageSize: int32(req.PageSize),
		Total:    int32(listResp.Total),
		Data:     list,
		Success:  true,
	}, nil
}
func (s *PmsService) ProductCategoryUpdate(ctx context.Context, req *pb.ProductCategoryUpdateReq) (*pb.ProductCategoryUpdateResp, error) {
	return &pb.ProductCategoryUpdateResp{}, nil
}
func (s *PmsService) ProductCategoryDelete(ctx context.Context, req *pb.ProductCategoryDeleteReq) (*pb.ProductCategoryDeleteResp, error) {
	return &pb.ProductCategoryDeleteResp{}, nil
}
func (s *PmsService) SkuStockAdd(ctx context.Context, req *pb.SkuStockAddReq) (*pb.SkuStockAddResp, error) {
	return &pb.SkuStockAddResp{}, nil
}
func (s *PmsService) SkuStockList(ctx context.Context, req *pb.SkuStockListReq) (*pb.SkuStockListResp, error) {
	listResp, _ := s.skuStockUseCase.ListSkuStock(ctx, &pms.SkuStockListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.SkuStockListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.SkuStockListResp{
		Code:     "000000",
		Message:  "查询订单信息成功",
		Current:  req.Current,
		PageSize: req.PageSize,
		Total:    listResp.Total,
		Data:     list,
		Success:  true,
	}, nil
}
func (s *PmsService) SkuStockUpdate(ctx context.Context, req *pb.SkuStockUpdateReq) (*pb.SkuStockUpdateResp, error) {
	return &pb.SkuStockUpdateResp{}, nil
}
func (s *PmsService) SkuStockDelete(ctx context.Context, req *pb.SkuStockDeleteReq) (*pb.SkuStockDeleteResp, error) {
	return &pb.SkuStockDeleteResp{}, nil
}
