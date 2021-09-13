package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	"kratos-mall/app/pms/internal/biz"

	pb "kratos-mall/api/pms/v1"
)

type PmsService struct {
	pb.UnimplementedPmsServer
	uc                    *biz.GreeterUsecase
	brandUseCase          *biz.BrandUseCase
	categoryUseCase       *biz.CategoryUseCase
	commentUseCase        *biz.CommentUseCase
	commentReplayUseCase  *biz.CommentReplayUseCase
	feightTemplateUseCase *biz.FeightTemplateUseCase
	memberPriceUseCase    *biz.MemberPriceUseCase
	operateHistoryUseCase *biz.OperateHistoryUseCase
	productUseCase        *biz.ProductUseCase
	skuStockUseCase       *biz.SkuStockUseCase
	vertifyRecordUseCase  *biz.VertifyRecordUseCase
	log                   *log.Helper
}

func NewPmsService(uc *biz.GreeterUsecase, logger log.Logger,
	brandUseCase *biz.BrandUseCase,
	categoryUseCase *biz.CategoryUseCase,
	commentUseCase *biz.CommentUseCase,
	commentReplayUseCase *biz.CommentReplayUseCase,
	feightTemplateUseCase *biz.FeightTemplateUseCase,
	memberPriceUseCase *biz.MemberPriceUseCase,
	operateHistoryUseCase *biz.OperateHistoryUseCase,
	productUseCase *biz.ProductUseCase,
	skuStockUseCase *biz.SkuStockUseCase,
	vertifyRecordUseCase *biz.VertifyRecordUseCase) *PmsService {
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
	listResp, _ := s.productUseCase.ListProduct(ctx, &biz.ProductListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.ProductListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.ProductListResp{
		Total: listResp.Total,
		List:  list,
	}, nil
}
func (s *PmsService) ProductUpdate(ctx context.Context, req *pb.ProductUpdateReq) (*pb.ProductUpdateResp, error) {
	return &pb.ProductUpdateResp{}, nil
}
func (s *PmsService) ProductDelete(ctx context.Context, req *pb.ProductDeleteReq) (*pb.ProductDeleteResp, error) {
	return &pb.ProductDeleteResp{}, nil
}
func (s *PmsService) AlbumAdd(ctx context.Context, req *pb.AlbumAddReq) (*pb.AlbumAddResp, error) {
	return &pb.AlbumAddResp{}, nil
}
func (s *PmsService) AlbumList(ctx context.Context, req *pb.AlbumListReq) (*pb.AlbumListResp, error) {
	return &pb.AlbumListResp{}, nil
}
func (s *PmsService) AlbumUpdate(ctx context.Context, req *pb.AlbumUpdateReq) (*pb.AlbumUpdateResp, error) {
	return &pb.AlbumUpdateResp{}, nil
}
func (s *PmsService) AlbumDelete(ctx context.Context, req *pb.AlbumDeleteReq) (*pb.AlbumDeleteResp, error) {
	return &pb.AlbumDeleteResp{}, nil
}
func (s *PmsService) AlbumPicAdd(ctx context.Context, req *pb.AlbumPicAddReq) (*pb.AlbumPicAddResp, error) {
	return &pb.AlbumPicAddResp{}, nil
}
func (s *PmsService) AlbumPicList(ctx context.Context, req *pb.AlbumPicListReq) (*pb.AlbumPicListResp, error) {
	return &pb.AlbumPicListResp{}, nil
}
func (s *PmsService) AlbumPicUpdate(ctx context.Context, req *pb.AlbumPicUpdateReq) (*pb.AlbumPicUpdateResp, error) {
	return &pb.AlbumPicUpdateResp{}, nil
}
func (s *PmsService) AlbumPicDelete(ctx context.Context, req *pb.AlbumPicDeleteReq) (*pb.AlbumPicDeleteResp, error) {
	return &pb.AlbumPicDeleteResp{}, nil
}
func (s *PmsService) BrandAdd(ctx context.Context, req *pb.BrandAddReq) (*pb.BrandAddResp, error) {
	return &pb.BrandAddResp{}, nil
}
func (s *PmsService) BrandList(ctx context.Context, req *pb.BrandListReq) (*pb.BrandListResp, error) {
	listResp, _ := s.brandUseCase.ListBrand(ctx, &biz.BrandListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.BrandListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.BrandListResp{
		Total: listResp.Total,
		List:  list,
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
	listResp, _ := s.commentUseCase.ListComment(ctx, &biz.CommentListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.CommentListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.CommentListResp{
		Total: listResp.Total,
		List:  list,
	}, nil
}
func (s *PmsService) CommentUpdate(ctx context.Context, req *pb.CommentUpdateReq) (*pb.CommentUpdateResp, error) {
	return &pb.CommentUpdateResp{}, nil
}
func (s *PmsService) CommentDelete(ctx context.Context, req *pb.CommentDeleteReq) (*pb.CommentDeleteResp, error) {
	return &pb.CommentDeleteResp{}, nil
}
func (s *PmsService) CommentReplayAdd(ctx context.Context, req *pb.CommentReplayAddReq) (*pb.CommentReplayAddResp, error) {
	return &pb.CommentReplayAddResp{}, nil
}
func (s *PmsService) CommentReplayList(ctx context.Context, req *pb.CommentReplayListReq) (*pb.CommentReplayListResp, error) {
	return &pb.CommentReplayListResp{}, nil
}
func (s *PmsService) CommentReplayUpdate(ctx context.Context, req *pb.CommentReplayUpdateReq) (*pb.CommentReplayUpdateResp, error) {
	return &pb.CommentReplayUpdateResp{}, nil
}
func (s *PmsService) CommentReplayDelete(ctx context.Context, req *pb.CommentReplayDeleteReq) (*pb.CommentReplayDeleteResp, error) {
	return &pb.CommentReplayDeleteResp{}, nil
}
func (s *PmsService) FeightTemplateAdd(ctx context.Context, req *pb.FeightTemplateAddReq) (*pb.FeightTemplateAddResp, error) {
	return &pb.FeightTemplateAddResp{}, nil
}
func (s *PmsService) FeightTemplateList(ctx context.Context, req *pb.FeightTemplateListReq) (*pb.FeightTemplateListResp, error) {
	listResp, _ := s.feightTemplateUseCase.ListFeightTemplate(ctx, &biz.FeightTemplateListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.FeightTemplateListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.FeightTemplateListResp{
		Total: listResp.Total,
		List:  list,
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
	listResp, _ := s.memberPriceUseCase.ListMemberPrice(ctx, &biz.MemberPriceListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.MemberPriceListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.MemberPriceListResp{
		Total: listResp.Total,
		List:  list,
	}, nil
}
func (s *PmsService) MemberPriceUpdate(ctx context.Context, req *pb.MemberPriceUpdateReq) (*pb.MemberPriceUpdateResp, error) {
	return &pb.MemberPriceUpdateResp{}, nil
}
func (s *PmsService) MemberPriceDelete(ctx context.Context, req *pb.MemberPriceDeleteReq) (*pb.MemberPriceDeleteResp, error) {
	return &pb.MemberPriceDeleteResp{}, nil
}
func (s *PmsService) ProductAttributeCategoryAdd(ctx context.Context, req *pb.ProductAttributeCategoryAddReq) (*pb.ProductAttributeCategoryAddResp, error) {
	return &pb.ProductAttributeCategoryAddResp{}, nil
}
func (s *PmsService) ProductAttributeCategoryList(ctx context.Context, req *pb.ProductAttributeCategoryListReq) (*pb.ProductAttributeCategoryListResp, error) {
	return &pb.ProductAttributeCategoryListResp{}, nil
}
func (s *PmsService) ProductAttributeCategoryUpdate(ctx context.Context, req *pb.ProductAttributeCategoryUpdateReq) (*pb.ProductAttributeCategoryUpdateResp, error) {
	return &pb.ProductAttributeCategoryUpdateResp{}, nil
}
func (s *PmsService) ProductAttributeCategoryDelete(ctx context.Context, req *pb.ProductAttributeCategoryDeleteReq) (*pb.ProductAttributeCategoryDeleteResp, error) {
	return &pb.ProductAttributeCategoryDeleteResp{}, nil
}
func (s *PmsService) ProductAttributeAdd(ctx context.Context, req *pb.ProductAttributeAddReq) (*pb.ProductAttributeAddResp, error) {
	return &pb.ProductAttributeAddResp{}, nil
}
func (s *PmsService) ProductAttributeList(ctx context.Context, req *pb.ProductAttributeListReq) (*pb.ProductAttributeListResp, error) {
	return &pb.ProductAttributeListResp{}, nil
}
func (s *PmsService) ProductAttributeUpdate(ctx context.Context, req *pb.ProductAttributeUpdateReq) (*pb.ProductAttributeUpdateResp, error) {
	return &pb.ProductAttributeUpdateResp{}, nil
}
func (s *PmsService) ProductAttributeDelete(ctx context.Context, req *pb.ProductAttributeDeleteReq) (*pb.ProductAttributeDeleteResp, error) {
	return &pb.ProductAttributeDeleteResp{}, nil
}
func (s *PmsService) ProductAttributeValueAdd(ctx context.Context, req *pb.ProductAttributeValueAddReq) (*pb.ProductAttributeValueAddResp, error) {
	return &pb.ProductAttributeValueAddResp{}, nil
}
func (s *PmsService) ProductAttributeValueList(ctx context.Context, req *pb.ProductAttributeValueListReq) (*pb.ProductAttributeValueListResp, error) {
	return &pb.ProductAttributeValueListResp{}, nil
}
func (s *PmsService) ProductAttributeValueUpdate(ctx context.Context, req *pb.ProductAttributeValueUpdateReq) (*pb.ProductAttributeValueUpdateResp, error) {
	return &pb.ProductAttributeValueUpdateResp{}, nil
}
func (s *PmsService) ProductAttributeValueDelete(ctx context.Context, req *pb.ProductAttributeValueDeleteReq) (*pb.ProductAttributeValueDeleteResp, error) {
	return &pb.ProductAttributeValueDeleteResp{}, nil
}
func (s *PmsService) ProductCategoryAttributeRelationAdd(ctx context.Context, req *pb.ProductCategoryAttributeRelationAddReq) (*pb.ProductCategoryAttributeRelationAddResp, error) {
	return &pb.ProductCategoryAttributeRelationAddResp{}, nil
}
func (s *PmsService) ProductCategoryAttributeRelationList(ctx context.Context, req *pb.ProductCategoryAttributeRelationListReq) (*pb.ProductCategoryAttributeRelationListResp, error) {
	return &pb.ProductCategoryAttributeRelationListResp{}, nil
}
func (s *PmsService) ProductCategoryAttributeRelationUpdate(ctx context.Context, req *pb.ProductCategoryAttributeRelationUpdateReq) (*pb.ProductCategoryAttributeRelationUpdateResp, error) {
	return &pb.ProductCategoryAttributeRelationUpdateResp{}, nil
}
func (s *PmsService) ProductCategoryAttributeRelationDelete(ctx context.Context, req *pb.ProductCategoryAttributeRelationDeleteReq) (*pb.ProductCategoryAttributeRelationDeleteResp, error) {
	return &pb.ProductCategoryAttributeRelationDeleteResp{}, nil
}
func (s *PmsService) ProductCategoryAdd(ctx context.Context, req *pb.ProductCategoryAddReq) (*pb.ProductCategoryAddResp, error) {
	return &pb.ProductCategoryAddResp{}, nil
}
func (s *PmsService) ProductCategoryList(ctx context.Context, req *pb.ProductCategoryListReq) (*pb.ProductCategoryListResp, error) {
	listResp, _ := s.categoryUseCase.ListCategory(ctx, &biz.CategoryListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.ProductCategoryListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.ProductCategoryListResp{
		Total: listResp.Total,
		List:  list,
	}, nil
}
func (s *PmsService) ProductCategoryUpdate(ctx context.Context, req *pb.ProductCategoryUpdateReq) (*pb.ProductCategoryUpdateResp, error) {
	return &pb.ProductCategoryUpdateResp{}, nil
}
func (s *PmsService) ProductCategoryDelete(ctx context.Context, req *pb.ProductCategoryDeleteReq) (*pb.ProductCategoryDeleteResp, error) {
	return &pb.ProductCategoryDeleteResp{}, nil
}
func (s *PmsService) ProductFullReductionAdd(ctx context.Context, req *pb.ProductFullReductionAddReq) (*pb.ProductFullReductionAddResp, error) {
	return &pb.ProductFullReductionAddResp{}, nil
}
func (s *PmsService) ProductFullReductionList(ctx context.Context, req *pb.ProductFullReductionListReq) (*pb.ProductFullReductionListResp, error) {
	return &pb.ProductFullReductionListResp{}, nil
}
func (s *PmsService) ProductFullReductionUpdate(ctx context.Context, req *pb.ProductFullReductionUpdateReq) (*pb.ProductFullReductionUpdateResp, error) {
	return &pb.ProductFullReductionUpdateResp{}, nil
}
func (s *PmsService) ProductFullReductionDelete(ctx context.Context, req *pb.ProductFullReductionDeleteReq) (*pb.ProductFullReductionDeleteResp, error) {
	return &pb.ProductFullReductionDeleteResp{}, nil
}
func (s *PmsService) ProductLadderAdd(ctx context.Context, req *pb.ProductLadderAddReq) (*pb.ProductLadderAddResp, error) {
	return &pb.ProductLadderAddResp{}, nil
}
func (s *PmsService) ProductLadderList(ctx context.Context, req *pb.ProductLadderListReq) (*pb.ProductLadderListResp, error) {
	return &pb.ProductLadderListResp{}, nil
}
func (s *PmsService) ProductLadderUpdate(ctx context.Context, req *pb.ProductLadderUpdateReq) (*pb.ProductLadderUpdateResp, error) {
	return &pb.ProductLadderUpdateResp{}, nil
}
func (s *PmsService) ProductLadderDelete(ctx context.Context, req *pb.ProductLadderDeleteReq) (*pb.ProductLadderDeleteResp, error) {
	return &pb.ProductLadderDeleteResp{}, nil
}
func (s *PmsService) ProductOperateLogAdd(ctx context.Context, req *pb.ProductOperateLogAddReq) (*pb.ProductOperateLogAddResp, error) {
	return &pb.ProductOperateLogAddResp{}, nil
}
func (s *PmsService) ProductOperateLogList(ctx context.Context, req *pb.ProductOperateLogListReq) (*pb.ProductOperateLogListResp, error) {
	return &pb.ProductOperateLogListResp{}, nil
}
func (s *PmsService) ProductOperateLogUpdate(ctx context.Context, req *pb.ProductOperateLogUpdateReq) (*pb.ProductOperateLogUpdateResp, error) {
	return &pb.ProductOperateLogUpdateResp{}, nil
}
func (s *PmsService) ProductOperateLogDelete(ctx context.Context, req *pb.ProductOperateLogDeleteReq) (*pb.ProductOperateLogDeleteResp, error) {
	return &pb.ProductOperateLogDeleteResp{}, nil
}
func (s *PmsService) ProductVertifyRecordAdd(ctx context.Context, req *pb.ProductVertifyRecordAddReq) (*pb.ProductVertifyRecordAddResp, error) {
	return &pb.ProductVertifyRecordAddResp{}, nil
}
func (s *PmsService) ProductVertifyRecordList(ctx context.Context, req *pb.ProductVertifyRecordListReq) (*pb.ProductVertifyRecordListResp, error) {
	return &pb.ProductVertifyRecordListResp{}, nil
}
func (s *PmsService) ProductVertifyRecordUpdate(ctx context.Context, req *pb.ProductVertifyRecordUpdateReq) (*pb.ProductVertifyRecordUpdateResp, error) {
	return &pb.ProductVertifyRecordUpdateResp{}, nil
}
func (s *PmsService) ProductVertifyRecordDelete(ctx context.Context, req *pb.ProductVertifyRecordDeleteReq) (*pb.ProductVertifyRecordDeleteResp, error) {
	return &pb.ProductVertifyRecordDeleteResp{}, nil
}
func (s *PmsService) SkuStockAdd(ctx context.Context, req *pb.SkuStockAddReq) (*pb.SkuStockAddResp, error) {
	return &pb.SkuStockAddResp{}, nil
}
func (s *PmsService) SkuStockList(ctx context.Context, req *pb.SkuStockListReq) (*pb.SkuStockListResp, error) {
	listResp, _ := s.skuStockUseCase.ListSkuStock(ctx, &biz.SkuStockListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.SkuStockListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.SkuStockListResp{
		Total: listResp.Total,
		List:  list,
	}, nil
}
func (s *PmsService) SkuStockUpdate(ctx context.Context, req *pb.SkuStockUpdateReq) (*pb.SkuStockUpdateResp, error) {
	return &pb.SkuStockUpdateResp{}, nil
}
func (s *PmsService) SkuStockDelete(ctx context.Context, req *pb.SkuStockDeleteReq) (*pb.SkuStockDeleteResp, error) {
	return &pb.SkuStockDeleteResp{}, nil
}
