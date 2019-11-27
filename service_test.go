package mic_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"gitlab.51baibao.com/server/mic"
)

var _ = Describe("Service", func() {
	It("new web", func() {
		_, cleanup, err := mic.DefaultWeb(mic.Opt{
			Name:       "foo",
			TracerAddr: "bj-etcd-dev-host-001.51baibao.com:6831",
		})
		Expect(err).ShouldNot(HaveOccurred())
		defer cleanup()
	})
})
