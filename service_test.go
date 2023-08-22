package mic_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go-micro.dev/v4"

	"gitlab.51baibao.com/server/mic/v4"
)

var _ = Describe("Service", func() {
	It("auto web", func() {
		s := micro.NewService(micro.Version("abc"), micro.Name("foo"))
		web := mic.DefaultWeb(mic.WebOpt{
			Addr:    ":8888",
			Service: s,
		})
		Expect(web.Options().Name).Should(Equal("foo.auto_web"))
		Expect(web.Options().Version).Should(Equal("abc"))
	})
	It("assign web web", func() {
		s := micro.NewService(micro.Version("abc"), micro.Name("foo"))
		web := mic.DefaultWeb(mic.WebOpt{
			Addr:    ":8888",
			Name:    "bar",
			Service: s,
		})
		Expect(web.Options().Name).Should(Equal("bar"))
		Expect(web.Options().Version).Should(Equal("abc"))
	})
})
