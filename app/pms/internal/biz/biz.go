package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewGreeterUsecase,
	NewBrandUseCase,
	NewCategoryUseCase,
	NewCommentUseCase,
	NewCommentReplayUseCase,
	NewFeightTemplateUseCase,
	NewMemberPriceUseCase,
	NewOperateHistoryUseCase,
	NewProductUseCase,
	NewSkuStockUseCase,
	NewVertifyRecordUseCase,
)
