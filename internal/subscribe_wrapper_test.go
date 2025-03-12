package internal_test

import (
	"context"
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go-micro.dev/v4/server"

	"github.com/qmute/mic/v4/internal"
	"github.com/qmute/mic/v4/internal/mocks/mserver"
)

var _ = Describe("SubscribeWrapper", func() {
	var mMsg *mserver.MockMessage
	BeforeEach(func() {
		mMsg = mserver.NewMockMessage(ctl)
		mMsg.EXPECT().Topic().Return("topic")
	})
	It("SubscribeErrLogWrapper", func() {
		h2 := internal.SubscribeErrLogWrapper(func(ctx context.Context, msg server.Message) error {
			return errors.New("haha")
		})
		err := h2(ctx, mMsg)
		Ω(err).To(HaveOccurred())
	})
})
