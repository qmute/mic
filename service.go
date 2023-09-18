package mic

import (
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	"go-micro.dev/v4/server"

	grpcClient "github.com/go-micro/plugins/v4/client/grpc"
	grpcServer "github.com/go-micro/plugins/v4/server/grpc"
	hystrixPlugin "github.com/go-micro/plugins/v4/wrapper/breaker/hystrix"
	"github.com/go-micro/plugins/v4/wrapper/monitoring/prometheus"
	limiter "github.com/go-micro/plugins/v4/wrapper/ratelimiter/uber"
	"go-micro.dev/v4/web"

	"gitlab.51baibao.com/server/mic/v4/internal"
)

// Opt grpc server 初始化选项
type Opt struct {
	Name string // name

	// optional
	TracerAddr     string        // tracer address
	Version        string        // service version
	Addr           string        // 监听地址
	HystrixTimeout time.Duration // 熔断时限, 默认 1s
	Limit          int           // 限流阈值, 默认 5000 qps
}

// GetLimit 限流阈值
func (p Opt) GetLimit() int {
	if p.Limit == 0 {
		return 5000
	}
	return p.Limit
}

// GetHystrixTimeout 获取熔断时限
func (p Opt) GetHystrixTimeout() time.Duration {
	if p.HystrixTimeout == 0 {
		return time.Second
	}
	return p.HystrixTimeout
}

// 当地址为空时，不作处理，框架会自动填充随机地址。 主动填空会报错
func optionalAddress(addr string) micro.Option {
	return func(o *micro.Options) {
		if addr == "" {
			return
		}
		if err := o.Server.Init(server.Address(addr)); err != nil {
			log.Fatal(err)
		}

	}
}

func optionalVersion(v string) micro.Option {
	return func(o *micro.Options) {
		if v == "" {
			return
		}
		if err := o.Server.Init(server.Version(v)); err != nil {
			log.Fatal(err)
		}
	}
}

// func optionalServerTrace(tracer opentracing.Tracer) micro.Option {
// 	if tracer == nil {
// 		return func(o *micro.Options) {}
// 	}
// 	return micro.WrapHandler(otplugin.NewHandlerWrapper(tracer))
// }
//
// func optionalSubscribeTrace(tracer opentracing.Tracer) micro.Option {
// 	if tracer == nil {
// 		return func(o *micro.Options) {}
// 	}
// 	return micro.WrapSubscriber(otplugin.NewSubscriberWrapper(tracer))
// }
//
// func optionalClientTrace(tracer opentracing.Tracer) micro.Option {
// 	if tracer == nil {
// 		return func(o *micro.Options) {}
// 	}
// 	return micro.WrapClient(otplugin.NewClientWrapper(tracer))
// }

// DefaultService 创建默认 micro.Service ，适用于 grpc server 绝大多数场景
// 如果想覆盖默认行为，可以后续在service.Init()中追加（例如version, addr等）
func DefaultService(opt Opt) (micro.Service, func(), error) {
	// var tracer opentracing.Tracer
	cleanup := func() {} // 默认啥也不干
	//
	// if opt.TracerAddr != "" {
	// 	var err error
	// 	tracer, cleanup, err = initGlobalTracer(traceOpt{Name: opt.Name, TracerAddr: opt.TracerAddr})
	// 	if err != nil {
	// 		return nil, nil, err
	// 	}
	// }

	service := micro.NewService(
		micro.Server(grpcServer.NewServer()),
		micro.Client(grpcClient.NewClient()),
		// common
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Name(opt.Name),
		micro.AfterStart(func() error {
			log.Infof("Service started %s %s %s", opt.Name, opt.Version, opt.Addr)
			return nil
		}),
		optionalVersion(opt.Version),
		optionalAddress(opt.Addr),

		// server 相关。执行顺序：正序。 先设置先执行
		micro.WrapHandler(internal.GrpcErrLogWrapper), // 错误日志
		// optionalServerTrace(tracer),                                  // server trace
		micro.WrapHandler(prometheus.NewHandlerWrapper()),            // 监控
		micro.WrapHandler(limiter.NewHandlerWrapper(opt.GetLimit())), // 限流

		// sub 相关
		micro.WrapSubscriber(internal.SubscribeErrLogWrapper), // 错误日志
		// optionalSubscribeTrace(tracer),                        // subscribe trace

		// client 相关。执行顺序：倒序。 最后设置的最先执行
		micro.WrapClient(hystrixPlugin.NewClientWrapper()), // 熔断
		// optionalClientTrace(tracer),                        // client trace， 包含 mq pub trace
	)

	// rpc server: graceful shutdown
	if err := service.Server().Init(server.Wait(nil)); err != nil {
		return nil, nil, err
	}

	hystrix.DefaultTimeout = int(opt.GetHystrixTimeout() / time.Millisecond)

	return service, cleanup, nil
}

// WebOpt web server 初始化选项
type WebOpt struct {
	Name    string
	Addr    string
	Service micro.Service
}

// DefaultWeb 用micro.Service创建默认 web.Service ，适用于 web server
func DefaultWeb(opt WebOpt) web.Service {
	if opt.Name == "" {
		// 名称不应为空
		opt.Name = opt.Service.Name() + ".auto_web"
	}
	version := opt.Service.Server().Options().Version
	return web.NewService(
		web.Address(opt.Addr),
		web.MicroService(opt.Service),
		web.Name(opt.Name),
		web.Version(version),
		web.AfterStart(func() error {
			log.Infof("Web started %s %s %s", opt.Name, version, opt.Addr)
			return nil
		}),
	)
}
