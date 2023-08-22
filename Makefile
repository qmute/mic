
.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: mock
mock:
	go generate ./... # mock+wire

.PHONY: test # 测试全部
test: fmt
	ginkgo -r .