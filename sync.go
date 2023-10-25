package mic

import (
	"log"

	"github.com/go-micro/plugins/v4/sync/consul"
	"github.com/go-micro/plugins/v4/sync/memory"
	"go-micro.dev/v4"
	"go-micro.dev/v4/sync"
)

// NewSync 根据当前 service 创建 Sync
func NewSync(service micro.Service) sync.Sync {
	registry := service.Options().Registry
	switch registry.String() {
	case "consul":
		// 如果是consul，那么就重用服务地址作consul sync
		return consul.NewSync(
			sync.Nodes(registry.Options().Addrs...),
			sync.Prefix("dlock"),
		)
	case "mdns":
		// 内存sync，用于本地开发
		return memory.NewSync()
	default:
		log.Panicf("sync not support %s", registry.String())
	}
	return nil
}
