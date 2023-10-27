package mic

import (
	"log"
	"os"
	"strings"
	"time"

	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/sync"

	"github.com/quexer/syncr"
)

type LockOptions struct {
	wait time.Duration // 等待时间
	ttl  time.Duration // 过期时间
}

type LockOption func(*LockOptions)

// LockWait 设置等待时间. 小于0表示永远等待，0表示不等待，大于0表示等待指定时间. 默认0
func LockWait(wait time.Duration) LockOption {
	return func(o *LockOptions) {
		o.wait = wait
	}
}

// LockTTL 过期时间，大于0表示到期自动解锁，否则表示永不自动解锁. 默认15s
func LockTTL(ttl time.Duration) LockOption {
	return func(o *LockOptions) {
		o.ttl = ttl
	}
}

/*
定义内存分区锁和分布式锁两个interface，剪裁 micro sync.Sync，只保留 Lock 和 Unlock 两个方法
这样可以在使用时明确区分，也方便依赖注入。注意：二者ttl语义不同
*/

// MemSync 内存分区锁，是 micro sync.Sync 的剪裁版本
// 仅在Lock超时时返回 sync.ErrLockTimeout，其它场景不会报错
// ttl 从锁定时算起
type MemSync interface {
	// Lock acquires a lock
	Lock(id string, opts ...LockOption) error
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
// 读取micro 注册中心环境变量，创建分布式锁实例。如果注册中心为空，则创建内存锁
// 目前支持 consul
func NewSync() Sync {
	tp := os.Getenv("MICRO_REGISTRY")
	addr := os.Getenv("MICRO_REGISTRY_ADDRESS")

	switch tp {
	case "consul":
		return &syncAdapter{mutex: syncr.NewConsulSync(
			sync.Nodes(strings.Split(addr, ",")...),
			sync.Prefix("dlock/"), // 写死前缀，避免与其它值冲突
		)}
	case "":
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

func (p *syncAdapter) Lock(id string, opts ...LockOption) error {
	o := &LockOptions{
		wait: 0,
		ttl:  15 * time.Second,
	}
	for _, opt := range opts {
		opt(o)
	}

	return p.mutex.Lock(id, sync.LockTTL(o.ttl), sync.LockWait(o.wait))
}
func (p *syncAdapter) Unlock(id string) {
	if err := p.mutex.Unlock(id); err != nil {
		logger.Error("unlock err", err)
	}
}
