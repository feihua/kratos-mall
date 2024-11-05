package data

import (
	"context"
	omsV1 "github.com/feihua/kratos-mall/api/oms/v1"
	payV1 "github.com/feihua/kratos-mall/api/pay/v1"
	pmsV1 "github.com/feihua/kratos-mall/api/pms/v1"
	smsV1 "github.com/feihua/kratos-mall/api/sms/v1"
	sysV1 "github.com/feihua/kratos-mall/api/sys/v1"
	umsV1 "github.com/feihua/kratos-mall/api/ums/v1"
	"github.com/feihua/kratos-mall/app/front/admin/internal/conf"
	consul "github.com/go-kratos/consul/registry"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewGreeterRepo,
	NewDiscovery,
	NewRegistrar,
	NewUmsServiceClient,
	NewPmsServiceClient,
	NewPayServiceClient,
	NewOmsServiceClient,
	NewSmsServiceClient,
	NewSysServiceClient,
)

// Data .
type Data struct {
	log       *log.Helper
	UmsClient umsV1.UmsClient
	PayClient payV1.PayClient
	PmsClient pmsV1.PmsClient
	SmsClient smsV1.SmsClient
	SysClient sysV1.SysClient
	OmsClient omsV1.OmsClient
}

// NewData .
func NewData(logger log.Logger, umsClient umsV1.UmsClient, payClient payV1.PayClient, pmsClient pmsV1.PmsClient, smsClient smsV1.SmsClient, sysClient sysV1.SysClient, omsClient omsV1.OmsClient) (*Data, error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	return &Data{
			log:       l,
			UmsClient: umsClient,
			PayClient: payClient,
			PmsClient: pmsClient,
			SmsClient: smsClient,
			SysClient: sysClient,
			OmsClient: omsClient},
		nil
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewUmsServiceClient(r registry.Discovery) umsV1.UmsClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///mall.ums.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := umsV1.NewUmsClient(conn)
	return c
}

func NewOmsServiceClient(r registry.Discovery) omsV1.OmsClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///mall.oms.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := omsV1.NewOmsClient(conn)
	return c
}

func NewPmsServiceClient(r registry.Discovery) pmsV1.PmsClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///mall.pms.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := pmsV1.NewPmsClient(conn)
	return c
}

func NewPayServiceClient(r registry.Discovery) payV1.PayClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///mall.pay.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := payV1.NewPayClient(conn)
	return c
}

func NewSysServiceClient(r registry.Discovery) sysV1.SysClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///mall.sys.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := sysV1.NewSysClient(conn)
	return c
}

func NewSmsServiceClient(r registry.Discovery) smsV1.SmsClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///mall.sms.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := smsV1.NewSmsClient(conn)
	return c
}
