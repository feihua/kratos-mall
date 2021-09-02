package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewGreeterUsecase, NewUserUseCase, NewRoleUseCase, NewMenuUseCase, NewLogUseCase, NewJobUseCase, NewDictUseCase, NewDeptUseCase)
