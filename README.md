# mic

针对go-micro框架的进一步封装及相关工具库， 降低使用门槛和学习成本

主入口见 DefaultService

install:  `go get gitlab.51baibao.com/server/mic`

注: go get 目前不支持gitlab subgroup。 问题修改后， 此库将移到 server/lib subgroup下

    https://gitlab.com/gitlab-org/gitlab-foss/issues/32149
    https://gitlab.com/gitlab-org/gitlab-foss/issues/37832
    https://github.com/golang/go/issues/34094
    
    
change log:

- v0.7.0 breaking change, upgrade micro to v2.9.1; replace `quexer/go-plugins/broker/rabbitmq` with `quexer/rmq`
- v0.5.6 support micro.Service / web.Service  
- v0.2 gin trace middleware
- v0.1 init  