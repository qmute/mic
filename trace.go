package mic

import (
	"context"
	"fmt"
	"time"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/util/log"
	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

type TraceOpt struct {
	Name       string // name
	TracerAddr string // tracer address
}

// init global tracer
func initGlobalTracer(opt TraceOpt) (opentracing.Tracer, func(), error) {
	cfg := jaegercfg.Configuration{
		ServiceName: opt.Name, // tracer name
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
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
		jaegercfg.Reporter(reporter),
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
			ctx, span, err := ocplugin.StartSpanFromContext(ctx, t, name)
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

// micro client wrapper, auto trace any call to any remote endpoint
func clientTraceWrapper(t opentracing.Tracer) client.Wrapper {
	return ocplugin.NewClientWrapper(t)
}

// server 端 trace
func ServerTrace(ctx context.Context, name string) opentracing.Span {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = make(map[string]string)
	}
	var sp opentracing.Span
	wireContext, _ := opentracing.GlobalTracer().Extract(opentracing.TextMap, opentracing.TextMapCarrier(md))
	sp = opentracing.StartSpan(name, opentracing.ChildOf(wireContext))
	return sp
}

// client 端 trace
func ClientTrace(ctx context.Context, name string) (context.Context, opentracing.Span) {
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
