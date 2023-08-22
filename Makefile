
.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: mock
mock:
	go generate ./... # mock+wire

.PHONY: test # 测试全部
test: fmt
	# 增加编译参数，解决protoregistry冲突问题
	go test -ldflags "-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn"