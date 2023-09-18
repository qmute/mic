package internal_test

import (
	"context"
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go-micro.dev/v4/server"

	"gitlab.51baibao.com/server/mic/v4/internal"
	"gitlab.51baibao.com/server/mic/v4/internal/mocks/mserver"
)

var _ = Describe("GrpcWrapper", func() {
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
