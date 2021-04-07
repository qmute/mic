package internal_test

import (
	"context"
	"errors"

	"github.com/micro/go-micro/v2/server"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"gitlab.51baibao.com/server/mic/internal"
	"gitlab.51baibao.com/server/mic/internal/mocks"
)

var _ = Describe("GrpcWrapper", func() {
	It("GrpcRecoveryWrapper", func() {
		h := func(ctx context.Context, req server.Request, rsp interface{}) error {
			panic("foo")
		}

		h2 := internal.GrpcRecoveryWrapper(h)

		f := func() {
			mReq := mocks.NewMockRequest(ctl)
			mReq.EXPECT().Service().Return("foo")
			mReq.EXPECT().Endpoint().Return("bar")
			_ = h2(ctx, mReq, nil)
		}
		Ω(f).NotTo(Panic())
	})
	It("GrpcErrLogWrapper", func() {
		h := func(ctx context.Context, req server.Request, rsp interface{}) error {
			return errors.New("haha")
		}

		h2 := internal.GrpcErrLogWrapper(h)

		mReq := mocks.NewMockRequest(ctl)
		mReq.EXPECT().Service().Return("foo")
		mReq.EXPECT().Endpoint().Return("bar")
		err := h2(ctx, mReq, nil)

		Ω(err).To(HaveOccurred())
	})
})
