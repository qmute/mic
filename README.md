# mic

针对go-micro框架的进一步封装及相关工具库， 降低使用门槛和学习成本

主入口见 DefaultService

install:  `go get github.com/qmute/mic`
    
change log:

- v0.7.0 breaking change, upgrade micro to v2.9.1; replace `quexer/go-plugins/broker/rabbitmq` with `quexer/rmq`
- v0.5.6 support micro.Service / web.Service  
- v0.2 gin trace middleware
- v0.1 init  