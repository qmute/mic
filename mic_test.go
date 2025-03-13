package mic_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/qmute/mic/v5"
)

var _ = Describe("Mic", func() {
	It("DefaultConfig", func() {
		_, err := DefaultConfig()
		Expect(err).ShouldNot(HaveOccurred())
	})
})
