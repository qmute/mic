package mic

import (
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
	hystrixPlugin "github.com/quexer/go-plugins/wrapper/breaker/hystrix"
	"github.com/quexer/go-plugins/wrapper/monitoring/prometheus"
	limiter "github.com/quexer/go-plugins/wrapper/ratelimiter/uber"
	otplugin "github.com/quexer/go-plugins/wrapper/trace/opentracing"
)

type Opt struct {
	Name       string // name
	TracerAddr string // tracer address

	// optional
	Version        string        // service version
	Addr           string        // 监听地址
	HystrixTimeout time.Duration // 熔断时限, 默认 1s
	Limit          int           // 限流阈值, 默认 5000 qps
}

func (p Opt) GetLimit() int {
	if p.Limit == 0 {
		return 5000
	}
	return p.Limit
}

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
		o.Server.Init(server.Address(addr))
	}
}

func optionalVersion(v string) micro.Option {
	return func(o *micro.Options) {
		if v == "" {
			return
		}
		log.Info("Version ", v)
		o.Server.Init(server.Version(v))
	}
}

func optionalWebAddress(addr string) web.Option {
	return func(o *web.Options) {
		if addr == "" {
			return
		}
		o.Address = addr
	}
}

func optionalWebVersion(v string) web.Option {
	return func(o *web.Options) {
		if v == "" {
			return
		}
		log.Info("Version ", v)
		o.Version = v
	}
}

// 创建默认 micro.Service ，适用于 grpc server 绝大多数场景
// 如果想覆盖默认行为，可以后续在service.Init()中追加（例如version, addr等）
func DefaultService(opt Opt) (micro.Service, func(), error) {
	tracer, cleanup, err := initGlobalTracer(traceOpt{Name: opt.Name, TracerAddr: opt.TracerAddr})
	if err != nil {
		return nil, nil, err
	}

	service := micro.NewService(
		// common
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Name(opt.Name),
		micro.AfterStart(func() error {
			log.Info("Service started")
			return nil
		}),
		optionalVersion(opt.Version),
		optionalAddress(opt.Addr),

		// server 相关。执行顺序：正序。 先设置先执行
		micro.WrapHandler(serverTraceWrapper(tracer)),                // server trace
		micro.WrapHandler(prometheus.NewHandlerWrapper()),            // 监控
		micro.WrapHandler(limiter.NewHandlerWrapper(opt.GetLimit())), // 限流

		// sub 相关
		micro.WrapSubscriber(subTraceWrapper(tracer)), // subscribe trace

		// client 相关。执行顺序：倒序。 最后设置的最先执行
		micro.WrapClient(hystrixPlugin.NewClientWrapper()),  // 熔断
		micro.WrapClient(otplugin.NewClientWrapper(tracer)), // client trace， 包含 mq pub trace
	)

	// rpc server: graceful shutdown
	if err := service.Server().Init(server.Wait(nil)); err != nil {
		return nil, nil, err
	}

	hystrix.DefaultTimeout = int(opt.GetHystrixTimeout() / time.Millisecond)

	return service, cleanup, nil
}

// 创建默认 web.Service ，适用于 web server
// 如果想覆盖默认行为，可以后续在service.Init()中追加（例如version, addr等）
func DefaultWeb(opt Opt) (web.Service, func(), error) {
	tracer, cleanup, err := initGlobalTracer(traceOpt{Name: opt.Name, TracerAddr: opt.TracerAddr})
	if err != nil {
		return nil, nil, err
	}

	// 此service 仅用作 client call， 不启动 grpc server
	service := micro.NewService(
		// common
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),

		// server 相关。执行顺序：正序。 先设置先执行
		micro.WrapHandler(prometheus.NewHandlerWrapper()), // 监控

		// sub 相关
		micro.WrapSubscriber(subTraceWrapper(tracer)), // subscribe trace

		// client 相关。执行顺序：倒序。 最后设置的最先执行
		micro.WrapClient(hystrixPlugin.NewClientWrapper()),  // 熔断
		micro.WrapClient(otplugin.NewClientWrapper(tracer)), // client trace， 包含 mq pub trace
	)

	hystrix.DefaultTimeout = int(opt.GetHystrixTimeout() / time.Millisecond)

	webService := web.NewService(
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*10),
		optionalWebAddress(opt.Addr),
		optionalWebVersion(opt.Version),
		web.Name(opt.Name),
		web.MicroService(service),
		web.AfterStart(func() error {
			log.Info("Service started")
			return nil
		}),
	)
	return webService, cleanup, nil
}
