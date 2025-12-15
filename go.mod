module github.com/qmute/mic/v4

go 1.23.0

toolchain go1.23.3

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/cockroachdb/errors v1.12.0
	github.com/gin-gonic/gin v1.11.0
	github.com/go-micro/plugins/v4/client/grpc v1.2.1
	github.com/go-micro/plugins/v4/logger/logrus v1.2.0
	github.com/go-micro/plugins/v4/server/grpc v1.2.0
	github.com/go-micro/plugins/v4/wrapper/breaker/hystrix v1.2.0
	github.com/go-micro/plugins/v4/wrapper/monitoring/prometheus v1.2.0
	github.com/go-micro/plugins/v4/wrapper/ratelimiter/uber v1.2.0
	github.com/onsi/ginkgo/v2 v2.27.3
	github.com/onsi/gomega v1.38.3
	github.com/quexer/rmq/v4 v4.0.3
	github.com/quexer/syncr v0.0.0-20231027072233-6a5824425470
	github.com/robfig/cron/v3 v3.0.1
	github.com/sirupsen/logrus v1.9.3
	go-micro.dev/v4 v4.10.2
	go.uber.org/mock v0.6.0
)

require (
	dario.cat/mergo v1.0.0 // indirect
	github.com/Masterminds/semver/v3 v3.4.0 // indirect
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/ProtonMail/go-crypto v0.0.0-20230717121422-5aa5874ade95 // indirect
	github.com/acomagu/bufpipe v1.0.4 // indirect
	github.com/armon/go-metrics v0.4.1 // indirect
	github.com/benbjohnson/clock v1.3.5 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bitly/go-simplejson v0.5.1 // indirect
	github.com/bytedance/sonic v1.14.0 // indirect
	github.com/bytedance/sonic/loader v0.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/cloudflare/circl v1.3.5 // indirect
	github.com/cloudwego/base64x v0.1.6 // indirect
	github.com/cockroachdb/logtags v0.0.0-20230118201751-21c54148d20b // indirect
	github.com/cockroachdb/redact v1.1.5 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/fatih/color v1.15.0 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.8 // indirect
	github.com/getsentry/sentry-go v0.27.0 // indirect
	github.com/gin-contrib/sse v1.1.0 // indirect
	github.com/go-git/gcfg v1.5.1-0.20230307220236-3a3c6141e376 // indirect
	github.com/go-git/go-billy/v5 v5.4.1 // indirect
	github.com/go-git/go-git/v5 v5.8.1 // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.27.0 // indirect
	github.com/go-task/slim-sprig/v3 v3.0.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/goccy/go-yaml v1.18.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/google/pprof v0.0.0-20250403155104-27863c87afa6 // indirect
	github.com/google/uuid v1.4.0 // indirect
	github.com/hashicorp/consul/api v1.25.1 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.5.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/hashicorp/serf v0.10.1 // indirect
	github.com/imdario/mergo v0.3.16 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kevinburke/ssh_config v1.2.0 // indirect
	github.com/klauspost/cpuid/v2 v2.3.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/miekg/dns v1.1.55 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/oxtoacart/bpool v0.0.0-20190530202638-03653db5a59c // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pelletier/go-toml/v2 v2.2.4 // indirect
	github.com/pjbgf/sha1cd v0.3.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v1.19.1 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.48.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/quic-go/qpack v0.5.1 // indirect
	github.com/quic-go/quic-go v0.54.0 // indirect
	github.com/rogpeppe/go-internal v1.13.1 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/sergi/go-diff v1.3.1 // indirect
	github.com/skeema/knownhosts v1.2.0 // indirect
	github.com/smartystreets/goconvey v1.8.1 // indirect
	github.com/streadway/amqp v1.0.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.3.0 // indirect
	github.com/urfave/cli/v2 v2.25.7 // indirect
	github.com/xanzy/ssh-agent v0.3.3 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	go.uber.org/ratelimit v0.3.0 // indirect
	go.yaml.in/yaml/v3 v3.0.4 // indirect
	golang.org/x/arch v0.20.0 // indirect
	golang.org/x/crypto v0.41.0 // indirect
	golang.org/x/exp v0.0.0-20231006140011-7918f672742d // indirect
	golang.org/x/mod v0.27.0 // indirect
	golang.org/x/net v0.43.0 // indirect
	golang.org/x/sync v0.16.0 // indirect
	golang.org/x/sys v0.35.0 // indirect
	golang.org/x/text v0.28.0 // indirect
	golang.org/x/tools v0.36.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
	google.golang.org/grpc v1.56.3 // indirect
	google.golang.org/protobuf v1.36.9 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
)
