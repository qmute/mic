# mic

wrapper and helper for [go-micro v4](https://github.com/micro/go-micro)


see DefaultService for main entry

install:  `go get github.com/qmute/mic/v4`

## 环境准备

```bash
# 安装单元测试命令工具
go install github.com/onsi/ginkgo/v2/ginkgo@latest
# 安装mockgen
go install go.uber.org/mock/mockgen@latest
# 安装swag
go install github.com/swaggo/swag/cmd/swag@latest
# 安装wire
go install github.com/google/wire/cmd/wire@latest
# 安装protoc
brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# 安装 micro protoc generator(v4)
go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest

# pprof 性能调优工具
go install github.com/google/pprof@latest
```

## 常见命令

- 格式化 `make fmt`
- 编译mock `make mock`
- 运行测试 `make test`

