package pms

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewBrandRepo,
	NewCategoryRepo,
	NewCommentRepo,
	NewCommentReplayRepo,
	NewFeightTemplateRepo,
	NewMemberPriceRepo,
	NewOperateHistoryRepo,
	NewProductRepo,
	NewSkuStockRepo,
	NewVertifyRecordRepo,
)
