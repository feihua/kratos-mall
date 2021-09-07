package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewGreeterUsecase,
	NewCouponUseCase,
	NewCouponHistoryUseCase,
	NewFlashPromotionUseCase,
	NewFlashPromotionHistoryUseCase,
	NewFlashPromotionSessionUseCase,
	NewHomeAdvertiseUseCase,
	NewHomeBrandUseCase,
	NewHomeNewProductUseCase,
	NewHomeRecommendProductUseCase,
	NewHomeRecommendSubjectUseCase,
)
