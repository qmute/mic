package mic

import (
	"log"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/sync"

	"github.com/quexer/syncr"
)

/*
定义内存分区锁和分布式锁两个interface，剪裁 micro sync.Sync，只保留 Lock 和 Unlock 两个方法
这样可以在使用时明确区分，也方便依赖注入。注意：二者ttl语义不同
*/

// MemSync 内存分区锁，是 micro sync.Sync 的剪裁版本
// 仅在Lock超时时返回 sync.ErrLockTimeout，其它场景不会报错
// ttl 从锁定时算起
type MemSync interface {
	// Lock acquires a lock
	Lock(id string, opts ...sync.LockOption) error
	// Unlock releases a lock
	Unlock(id string)
}

// Sync 分布式锁，是 micro sync.Sync 的剪裁版本
// 注意：ttl 从程序时退出算起，相当于重启前一直有效，所以一定要 Unlock
// wait 默认15s
type Sync interface {
	MemSync
}

// NewMemSync 创建内存锁
func NewMemSync() MemSync {
	return &syncAdapter{mutex: syncr.NewMemorySync()}
}

// NewSync 创建分布式锁
// 根据当前注册中心创建对应实现。已支持的注册中心有：本地mdns、consul
// 注：本地开发时会回退到内存锁，不具备跨进程能力
func NewSync(service micro.Service) Sync {
	registry := service.Options().Registry
	switch registry.String() {
	case "consul":
		// 如果是consul，那么就重用服务地址作consul sync
		return &syncAdapter{mutex: syncr.NewConsulSync(
			sync.Nodes(registry.Options().Addrs...),
			sync.Prefix("dlock/"), // 写死前缀，避免与其它值
		)}
	case "mdns":
		// 内存锁，用于本地开发
		return &syncAdapter{mutex: syncr.NewMemorySync()}
	default:
		log.Panicf("sync not support %s", registry.String())
	}
	return nil
}

// syncAdapter 适配 sync.Sync
type syncAdapter struct {
	mutex sync.Sync
}

func (p *syncAdapter) Lock(id string, opts ...sync.LockOption) error {
	return p.mutex.Lock(id, opts...)
}
func (p *syncAdapter) Unlock(id string) {
	if err := p.mutex.Unlock(id); err != nil {
		logger.Error("unlock err", err)
	}
}
