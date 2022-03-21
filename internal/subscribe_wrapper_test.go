package internal_test

import (
	"context"
	"errors"

	"github.com/micro/go-micro/v2/server"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"gitlab.51baibao.com/server/mic/internal"
	"gitlab.51baibao.com/server/mic/internal/mocks/mserver"
)

var _ = Describe("SubscribeWrapper", func() {
	var mMsg *mserver.MockMessage
	BeforeEach(func() {
		mMsg = mserver.NewMockMessage(ctl)
		mMsg.EXPECT().Topic().Return("topic")
	})
	It("SubscribeRecoveryWrapper", func() {
		h2 := internal.SubscribePanicWrapper(func(ctx context.Context, msg server.Message) error {
			panic("foo")
		})
		f := func() {
			_ = h2(ctx, mMsg)
		}
		Ω(f).NotTo(Panic())
	})
	It("SubscribeErrLogWrapper", func() {
		h2 := internal.SubscribeErrLogWrapper(func(ctx context.Context, msg server.Message) error {
			return errors.New("haha")
		})
		err := h2(ctx, mMsg)
		Ω(err).To(HaveOccurred())
	})
})
