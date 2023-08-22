package mic

// type traceOpt struct {
// 	Name       string // name
// 	TracerAddr string // tracer address
// }

// // init global tracer, 创建并设置全局tracer
// func initGlobalTracer(opt traceOpt) (opentracing.Tracer, func(), error) {
// 	cfg := config.Configuration{
// 		ServiceName: opt.Name, // tracer name
// 		Sampler: &config.SamplerConfig{
// 			Type:  jaeger.SamplerTypeConst,
// 			Param: 1,
// 		},
// 		Reporter: &config.ReporterConfig{
// 			LogSpans:            true,
// 			BufferFlushInterval: 1 * time.Second,
// 		},
// 	}
// 	sender, err := jaeger.NewUDPTransport(opt.TracerAddr, 0) // set Jaeger report receive address
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	reporter := jaeger.NewRemoteReporter(sender) // create Jaeger reporter
// 	// Initialize Opentracing tracer with Jaeger Reporter
// 	tracer, closer, err := cfg.NewTracer(
// 		config.Reporter(reporter),
// 	)
//
// 	// 这样就不必四处传递tracer了
// 	opentracing.SetGlobalTracer(tracer)
//
// 	return tracer, func() {
// 		if err := closer.Close(); err != nil {
// 			log.Error(err)
// 		}
// 	}, err
// }

// // Trace 创建trace所需 context 和  trace span
// func Trace(ctx context.Context, name string) (context.Context, opentracing.Span) {
// 	span, ctx := opentracing.StartSpanFromContext(ctx, name)
// 	md, ok := metadata.FromContext(ctx)
// 	if !ok {
// 		md = make(map[string]string)
// 	}
// 	if err := opentracing.GlobalTracer().Inject(span.Context(), opentracing.TextMap, opentracing.TextMapCarrier(md)); err != nil {
// 		log.Error(err)
// 	}
// 	ctx = opentracing.ContextWithSpan(ctx, span)
// 	ctx = metadata.NewContext(ctx, md)
// 	return ctx, span
// }
