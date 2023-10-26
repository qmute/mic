package mic

import (
	"log"

	"go-micro.dev/v4"
	"go-micro.dev/v4/sync"

	"github.com/quexer/syncr"
)

/*

定义内存锁和分布式锁两个interface，剪裁 micro sync.Sync，只保留 Lock 和 Unlock 两个方法

这样可以在使用时明确区分，也方便依赖注入。注意：虽然定义二者方法定义相同，但具体行为有差异，详见各自说明
*/

// MemSync 内存分区锁
// 仅在等待lock超时情况下返回 sync.ErrLockTimeout，其它场景不会报错
type MemSync interface {
	// Lock acquires a lock
	Lock(id string, opts ...sync.LockOption) error
	// Unlock releases a lock
	Unlock(id string) error
}

// Sync 分布式锁，是 micro sync.Sync 的剪裁版本
// 分布式环境下，为避免锁定后永不超时， ttl 默认值为15s，且不能小于10s
// wait 默认值为15秒
type Sync interface {
	MemSync
}

// NewMemSync 创建内存锁
func NewMemSync() MemSync {
	return syncr.NewMemorySync()
}

// NewSync 创建分布式锁
// 根据当前注册中心创建对应实现。已支持的注册中心有：本地mdns、consul
// 注：本地开发时会回退到内存锁，不具备跨进程能力
func NewSync(service micro.Service) Sync {
	registry := service.Options().Registry
	switch registry.String() {
	case "consul":
		// 如果是consul，那么就重用服务地址作consul sync
		return syncr.NewConsulSync(
			sync.Nodes(registry.Options().Addrs...),
			sync.Prefix("dlock/"), // 写死前缀，避免与其它值
		)
	case "mdns":
		// 内存锁，用于本地开发
		return syncr.NewMemorySync()
	default:
		log.Panicf("sync not support %s", registry.String())
	}
	return nil
}
