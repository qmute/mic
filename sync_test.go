package mic_test

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/quexer/syncr"
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

			mutex := mic.NewSync("consul", "bj-meishi-dev-host.51baibao.com:8500")
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
