package mic_test

import (
	"fmt"
	"log"
	"os"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go-micro.dev/v4/sync"

	"github.com/qmute/mic/v4"
)

var _ = Describe("Sync", func() {
	Context("MemSync", func() {

		PIt("Default Ttl", func() {
			mutex := mic.NewMemSync()
			_ = mutex.Lock("test")
			time.Sleep(16 * time.Second)
			err := mutex.Lock("test")
			Ω(err).To(Succeed())
		})

		It("Ttl", func() {
			mutex := mic.NewMemSync()
			// 连续锁定 ttl 累加
			_ = mutex.Lock("test", mic.LockTTL(time.Second))
			_ = mutex.Lock("test", mic.LockTTL(time.Second))
			Ω(true).Should(BeTrue())
		})
		It("Wait", func() {
			mutex := mic.NewMemSync()

			go func() {
				_ = mutex.Lock("test")
			}()

			fn := func() error {
				return mutex.Lock("test", mic.LockWait(100*time.Microsecond))
			}

			Eventually(fn).Should(MatchError(sync.ErrLockTimeout))
		})
	})
	Context("Sync", func() {

		PIt("Ttl", func() {
			err := os.Setenv("MICRO_REGISTRY", "consul")
			Ω(err).To(Succeed())
			err = os.Setenv("MICRO_REGISTRY_ADDRESS", "bj-meishi-dev-host.51baibao.com:8500")
			Ω(err).To(Succeed())

			mutex := mic.NewSync()

			id := fmt.Sprintf("test:%d", time.Now().UnixMilli())
			{
				err := mutex.Lock(id)
				Ω(err).To(Succeed())
				log.Println(0)
				// 立即重试会报错
				err = mutex.Lock(id)
				Ω(err).To(MatchError(sync.ErrLockTimeout))
				log.Println(1)
			}

			{
				// 歇到超时
				time.Sleep(16 * time.Second)
				err := mutex.Lock(id, mic.LockTTL(10*time.Second), mic.LockWait(10*time.Second))
				// 会成功
				Ω(err).To(Succeed())
				log.Println(2)
			}
			// time.Sleep(5 * time.Minute)
			{
				mutex.Unlock(id)
				log.Println(3)
			}
		})
	})
})
