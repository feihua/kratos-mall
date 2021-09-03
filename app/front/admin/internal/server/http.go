package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "kratos-mall/api/front/admin/v1"
	"kratos-mall/app/front/admin/internal/conf"
	"kratos-mall/app/front/admin/internal/service"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, oms *service.OmsService, pay *service.PayService, pms *service.PmsService, sms *service.SmsService, sys *service.SysService, ums *service.UmsService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterGreeterHTTPServer(srv, greeter)
	v1.RegisterOmsHTTPServer(srv, oms)
	v1.RegisterPayHTTPServer(srv, pay)
	v1.RegisterPmsHTTPServer(srv, pms)
	v1.RegisterSmsHTTPServer(srv, sms)
	v1.RegisterSysHTTPServer(srv, sys)
	v1.RegisterUmsHTTPServer(srv, ums)
	return srv
}
