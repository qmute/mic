package mic_test

import (
	"github.com/micro/go-micro"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"gitlab.51baibao.com/server/mic"
)

var _ = Describe("Service", func() {
	It("new web", func() {
		s := micro.NewService(micro.Version("abc"), micro.Name("foo"))
		web := mic.DefaultWeb(mic.WebOpt{
			Addr:    ":8888",
			Service: s,
		})
		Expect(web.Options().Name).Should(Equal("foo.web"))
		Expect(web.Options().Version).Should(Equal("abc"))
	})
})
