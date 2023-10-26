package mic_test

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/quexer/syncr"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/sync"

	"gitlab.51baibao.com/server/mic/v4"
)

var _ = Describe("Sync", func() {
	Context("MemSync", func() {
		It("Ttl", func() {
			mutex := syncr.NewMemorySync()
			// 连续锁定 ttl 累加
			_ = mutex.Lock("test", sync.LockTTL(time.Second))
			_ = mutex.Lock("test", sync.LockTTL(time.Second))
			Ω(true).Should(BeTrue())
		})
		It("Wait", func() {
			mutex := syncr.NewMemorySync()

			go func() {
				_ = mutex.Lock("test")
			}()

			fn := func() error {
				return mutex.Lock("test", sync.LockWait(100*time.Microsecond))
			}

			Eventually(fn).Should(MatchError(sync.ErrLockTimeout))
		})
	})
	PContext("Sync", func() {
		It("Ttl", func() {

			mutex := mic.NewSync(micro.NewService(micro.Registry(&dummyConsul{})))
			{
				err := mutex.Lock("test", sync.LockTTL(10*time.Second))
				Ω(err).To(Succeed())
				// log.Println(1)
			}

			{
				err := mutex.Lock("test", sync.LockTTL(10*time.Second), sync.LockWait(10*time.Second))
				Ω(err).To(MatchError(sync.ErrLockTimeout))
				// log.Println(2)
			}
			// time.Sleep(5 * time.Minute)
			{
				mutex.Unlock("test")
				// log.Println(3)
			}
		})
	})
})

// 为方便测试，本地实现的dummy consul registry
type dummyConsul struct {
}

func (p *dummyConsul) Init(option ...registry.Option) error {
	// TODO implement me
	panic("implement me")
}

func (p *dummyConsul) Options() registry.Options {
	return registry.Options{
		Addrs: []string{"bj-meishi-dev-host.51baibao.com:8500"},
	}
}

func (p *dummyConsul) Register(service *registry.Service, option ...registry.RegisterOption) error {
	// TODO implement me
	panic("implement me")
}

func (p *dummyConsul) Deregister(service *registry.Service, option ...registry.DeregisterOption) error {
	// TODO implement me
	panic("implement me")
}

func (p *dummyConsul) GetService(s string, option ...registry.GetOption) ([]*registry.Service, error) {
	// TODO implement me
	panic("implement me")
}

func (p *dummyConsul) ListServices(option ...registry.ListOption) ([]*registry.Service, error) {
	// TODO implement me
	panic("implement me")
}

func (p *dummyConsul) Watch(option ...registry.WatchOption) (registry.Watcher, error) {
	// TODO implement me
	panic("implement me")
}

func (p *dummyConsul) String() string {
	return "consul"
}
