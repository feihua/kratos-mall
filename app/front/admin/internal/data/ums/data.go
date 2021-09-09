package ums

import (
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewChangeHistoryRepo,
	NewIntegrationChangeHistoryRepo,
	NewConsumeSettingRepo,
	NewLoginLogRepo,
	NewMemberRepo,
	NewMemberAddressRepo,
	NewMemberLevelRepo,
	NewRuleSettingRepo,
	NewTagRepo,
	NewTaskRepo,
)
