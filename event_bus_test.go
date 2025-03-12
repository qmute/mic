package mic_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"

	"github.com/qmute/mic/v4/internal/mocks"
	"github.com/qmute/mic/v4/internal/mocks/mserver"

	"github.com/qmute/mic/v4"
)

var _ = Describe("EventBus", func() {
	var bus *mic.MicroEventBus

	var mService *mocks.MockService
	BeforeEach(func() {
		mService = mocks.NewMockService(ctl)

		bus = mic.NewMicroEventBus(mService)
	})
	It("Pub", func() {
		mMsg := mocks.NewMockMessage(ctl)

		mClient := mocks.NewMockClient(ctl)
		mClient.EXPECT().Publish(gomock.Any(), mMsg, gomock.Any()).Times(2)
		mClient.EXPECT().NewMessage("foo", "msg").Times(2).Return(mMsg)

		mService.EXPECT().Client().Return(mClient)

		err := bus.Pub(ctx, "foo", "msg")
		Ω(err).NotTo(HaveOccurred())

		// 连发两次， pub只建一次, 消息发两次
		err = bus.Pub(ctx, "foo", "msg")
		Ω(err).NotTo(HaveOccurred())
	})
	It("Sub", func() {
		mServer := mserver.NewMockServer(ctl)
		mServer.EXPECT().NewSubscriber("foo", gomock.Any(), gomock.Any()).Return(nil)
		mServer.EXPECT().Subscribe(nil)

		mService.EXPECT().Server().Return(mServer)

		err := bus.Sub("foo", func() {})
		Ω(err).NotTo(HaveOccurred())
	})
})
