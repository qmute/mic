package mic

import (
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/util/log"
	hystrixPlugin "github.com/micro/go-plugins/wrapper/breaker/hystrix"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus"
	limiter "github.com/micro/go-plugins/wrapper/ratelimiter/uber"
	otplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"
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
		log.Info()
		o.Server.Init(server.Address(addr))
	}
}

func optionalVersion(ver string) micro.Option {
	return func(o *micro.Options) {
		if ver == "" {
			return
		}
		o.Server.Init(server.Version(ver))
	}
}

// 创建默认 micro.Service ，适用于绝大多数场景
// 如果想覆盖默认行为不够，可以后续在service.Init()中追加（例如version, port）
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

	hystrix.DefaultTimeout = int(opt.GetHystrixTimeout() / time.Millisecond)

	return service, cleanup, nil
}
