package server

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	v1 "kratos-mall/api/front/admin/v1"
	"kratos-mall/app/front/admin/internal/conf"
	jwt "kratos-mall/app/front/admin/internal/pkg/middleware"
	"kratos-mall/app/front/admin/internal/service"
)

func getOperation(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		if tr, ok := transport.FromServerContext(ctx); ok {
			fmt.Println(tr.Operation())
		}
		return handler(ctx, req)
	}
}

func MatchFunc(x string) bool {
	return true
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, oms *service.OmsService, pay *service.PayService, pms *service.PmsService, sms *service.SmsService, sys *service.SysService, ums *service.UmsService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
			selector.Server(jwt.AuthMiddleware()).Match(MatchFunc).
				Build(),
			getOperation,
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
	openApiHandler := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", openApiHandler)
	v1.RegisterGreeterHTTPServer(srv, greeter)
	v1.RegisterOmsHTTPServer(srv, oms)
	v1.RegisterPayHTTPServer(srv, pay)
	v1.RegisterPmsHTTPServer(srv, pms)
	v1.RegisterSmsHTTPServer(srv, sms)
	v1.RegisterSysHTTPServer(srv, sys)
	v1.RegisterUmsHTTPServer(srv, ums)
	return srv
}
