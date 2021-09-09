package sys

import (
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewUserRepo,
	NewLogRepo,
	NewMenuRepo,
	NewDeptRepo,
	NewDictRepo,
	NewJobRepo,
	NewRoleRepo,
)
