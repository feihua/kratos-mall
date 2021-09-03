package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "kratos-mall/api/front/admin/v1"
	"kratos-mall/app/front/admin/internal/conf"
	"kratos-mall/app/front/admin/internal/service"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, greeter *service.GreeterService, oms *service.OmsService, pay *service.PayService, pms *service.PmsService, sms *service.SmsService, sys *service.SysService, ums *service.UmsService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterGreeterServer(srv, greeter)
	v1.RegisterOmsServer(srv, oms)
	v1.RegisterPayServer(srv, pay)
	v1.RegisterPmsServer(srv, pms)
	v1.RegisterSmsServer(srv, sms)
	v1.RegisterSysServer(srv, sys)
	v1.RegisterUmsServer(srv, ums)
	return srv
}
