package sms

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
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
