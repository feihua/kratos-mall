package ums

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewChangeHistoryUseCase,
	NewIntegrationChangeHistoryUseCase,
	NewConsumeSettingUseCase,
	NewLoginLogUseCase,
	NewMemberUseCase,
	NewMemberAddressUseCase,
	NewMemberLevelUseCase,
	NewRuleSettingUseCase,
	NewTagUseCase,
	NewTaskUseCase,
)
