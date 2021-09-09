package pay

import (
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewMerchantRepo, NewRecordRepo)
