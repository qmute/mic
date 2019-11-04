package mic

import (
	"context"
	"fmt"
	"time"

	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/util/log"
	"github.com/opentracing/opentracing-go"
	otplugin "github.com/quexer/go-plugins/wrapper/trace/opentracing"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

type traceOpt struct {
	Name       string // name
	TracerAddr string // tracer address
}

// init global tracer, 创建并设置全局tracer
func initGlobalTracer(opt traceOpt) (opentracing.Tracer, func(), error) {
	cfg := config.Configuration{
		ServiceName: opt.Name, // tracer name
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
	}
	sender, err := jaeger.NewUDPTransport(opt.TracerAddr, 0) // set Jaeger report receive address
	if err != nil {
		return nil, nil, err
	}
	reporter := jaeger.NewRemoteReporter(sender) // create Jaeger reporter
	// Initialize Opentracing tracer with Jaeger Reporter
	tracer, closer, err := cfg.NewTracer(
		config.Reporter(reporter),
	)

	// 这样就不必四处传递tracer了
	opentracing.SetGlobalTracer(tracer)

	return tracer, func() {
		if err := closer.Close(); err != nil {
			log.Error(err)
		}
	}, err
}

// micro server wrapper, auto trace any call to any local endpoint
func serverTraceWrapper(t opentracing.Tracer) server.HandlerWrapper {
	return func(h server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			name := fmt.Sprintf("%s.%s", req.Service(), req.Endpoint())
			ctx, span, err := otplugin.StartSpanFromContext(ctx, t, name)
			if err != nil {
				return err
			}
			span.SetTag("req", req.Body())
			defer span.SetTag("res", rsp).Finish()
			err = h(ctx, req, rsp)
			if err != nil {
				span.SetTag("err", err)
			}
			return err
		}
	}
}

// subscribe wrapper, auto trace message receive
// todo 改写自 otplugin.NewSubscriberWrapper，因为官方实现中把 "Sub from" 写成了 "Pub to". 等上游修正后， 即可移除
func subTraceWrapper(t opentracing.Tracer) server.SubscriberWrapper {
	return func(next server.SubscriberFunc) server.SubscriberFunc {
		return func(ctx context.Context, msg server.Message) error {
			name := "Sub from " + msg.Topic()
			ctx, span, err := otplugin.StartSpanFromContext(ctx, t, name)
			if err != nil {
				return err
			}
			defer span.Finish()
			return next(ctx, msg)
		}
	}
}

// Trace
func Trace(ctx context.Context, name string) (context.Context, opentracing.Span) {
	span, ctx := opentracing.StartSpanFromContext(ctx, name)
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = make(map[string]string)
	}
	if err := opentracing.GlobalTracer().Inject(span.Context(), opentracing.TextMap, opentracing.TextMapCarrier(md)); err != nil {
		log.Error(err)
	}
	ctx = opentracing.ContextWithSpan(ctx, span)
	ctx = metadata.NewContext(ctx, md)
	return ctx, span
}
