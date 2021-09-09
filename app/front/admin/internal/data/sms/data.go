package sms

import (
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewCouponRepo,
	NewCouponHistoryRepo,
	NewFlashPromotionRepo,
	NewFlashPromotionHistoryRepo,
	NewFlashPromotionSessionRepo,
	NewHomeAdvertiseRepo,
	NewHomeBrandRepo,
	NewHomeNewProductRepo,
	NewHomeRecommendProductRepo,
	NewHomeRecommendSubjectRepo,
)
