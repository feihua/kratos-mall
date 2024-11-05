// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/feihua/kratos-mall/app/pms/internal/biz"
	"github.com/feihua/kratos-mall/app/pms/internal/conf"
	"github.com/feihua/kratos-mall/app/pms/internal/data"
	"github.com/feihua/kratos-mall/app/pms/internal/server"
	"github.com/feihua/kratos-mall/app/pms/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewDB(confData, logger)
	dataData, cleanup, err := data.NewData(db, logger)
	if err != nil {
		return nil, nil, err
	}
	greeterRepo := data.NewGreeterRepo(dataData, logger)
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo, logger)
	brandRepo := data.NewBrandRepo(dataData, logger)
	brandUseCase := biz.NewBrandUseCase(brandRepo, logger)
	categoryRepo := data.NewCategoryRepo(dataData, logger)
	categoryUseCase := biz.NewCategoryUseCase(categoryRepo, logger)
	commentRepo := data.NewCommentRepo(dataData, logger)
	commentUseCase := biz.NewCommentUseCase(commentRepo, logger)
	commentReplayRepo := data.NewCommentReplayRepo(dataData, logger)
	commentReplayUseCase := biz.NewCommentReplayUseCase(commentReplayRepo, logger)
	feightTemplateRepo := data.NewFeightTemplateRepo(dataData, logger)
	feightTemplateUseCase := biz.NewFeightTemplateUseCase(feightTemplateRepo, logger)
	memberPriceRepo := data.NewMemberPriceRepo(dataData, logger)
	memberPriceUseCase := biz.NewMemberPriceUseCase(memberPriceRepo, logger)
	operateHistoryRepo := data.NewOperateHistoryRepo(dataData, logger)
	operateHistoryUseCase := biz.NewOperateHistoryUseCase(operateHistoryRepo, logger)
	productRepo := data.NewProductRepo(dataData, logger)
	productUseCase := biz.NewProductUseCase(productRepo, logger)
	skuStockRepo := data.NewSkuStockRepo(dataData, logger)
	skuStockUseCase := biz.NewSkuStockUseCase(skuStockRepo, logger)
	vertifyRecordRepo := data.NewVertifyRecordRepo(dataData, logger)
	vertifyRecordUseCase := biz.NewVertifyRecordUseCase(vertifyRecordRepo, logger)
	pmsService := service.NewPmsService(greeterUsecase, logger, brandUseCase, categoryUseCase, commentUseCase, commentReplayUseCase, feightTemplateUseCase, memberPriceUseCase, operateHistoryUseCase, productUseCase, skuStockUseCase, vertifyRecordUseCase)
	grpcServer := server.NewGRPCServer(confServer, pmsService, logger)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
