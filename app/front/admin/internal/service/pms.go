package service

import (
	"context"

	pb "kratos-mall/api/front/admin/v1"
)

type PmsService struct {
	pb.UnimplementedPmsServer
}

func NewPmsService() *PmsService {
	return &PmsService{}
}

func (s *PmsService) ProductAdd(ctx context.Context, req *pb.ProductAddReq) (*pb.ProductAddResp, error) {
	return &pb.ProductAddResp{}, nil
}
func (s *PmsService) ProductList(ctx context.Context, req *pb.ProductListReq) (*pb.ProductListResp, error) {
	return &pb.ProductListResp{}, nil
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
	return &pb.BrandListResp{}, nil
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
	return &pb.CommentListResp{}, nil
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
	return &pb.FeightTemplateListResp{}, nil
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
	return &pb.MemberPriceListResp{}, nil
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
	return &pb.ProductCategoryListResp{}, nil
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
	return &pb.SkuStockListResp{}, nil
}
func (s *PmsService) SkuStockUpdate(ctx context.Context, req *pb.SkuStockUpdateReq) (*pb.SkuStockUpdateResp, error) {
	return &pb.SkuStockUpdateResp{}, nil
}
func (s *PmsService) SkuStockDelete(ctx context.Context, req *pb.SkuStockDeleteReq) (*pb.SkuStockDeleteResp, error) {
	return &pb.SkuStockDeleteResp{}, nil
}
