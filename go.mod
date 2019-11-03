module gitlab.51baibao.com/server/mic

go 1.13

replace github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.1

replace github.com/uber/jaeger-lib => github.com/uber/jaeger-lib v2.0.0+incompatible

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/gin-gonic/gin v1.4.0
	github.com/micro/go-micro v1.14.0
	github.com/micro/go-plugins v1.4.0
	github.com/onsi/ginkgo v1.8.0
	github.com/onsi/gomega v1.5.0
	github.com/opentracing/opentracing-go v1.1.0
	github.com/quexer/go-plugins v1.4.2
	github.com/uber/jaeger-client-go v2.19.0+incompatible
)
