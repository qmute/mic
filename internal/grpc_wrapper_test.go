package internal_test

import (
	"context"
	"errors"

	"github.com/micro/go-micro/v2/server"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"gitlab.51baibao.com/server/mic/internal"
	"gitlab.51baibao.com/server/mic/internal/mocks/mserver"
)

var _ = Describe("GrpcWrapper", func() {
	It("GrpcRecoveryWrapper", func() {
		h2 := internal.GrpcRecoveryWrapper(func(ctx context.Context, req server.Request, rsp interface{}) error {
			panic("foo")
		})

		f := func() {
			mReq := mserver.NewMockRequest(ctl)
			mReq.EXPECT().Service().Return("foo")
			mReq.EXPECT().Endpoint().Return("bar")
			_ = h2(ctx, mReq, nil)
		}
		Ω(f).NotTo(Panic())
	})
	It("GrpcErrLogWrapper", func() {
		h2 := internal.GrpcErrLogWrapper(func(ctx context.Context, req server.Request, rsp interface{}) error {
			return errors.New("haha")
		})

		mReq := mserver.NewMockRequest(ctl)
		mReq.EXPECT().Service().Return("foo")
		mReq.EXPECT().Endpoint().Return("bar")
		err := h2(ctx, mReq, nil)

		Ω(err).To(HaveOccurred())
	})
})
