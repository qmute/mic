package internal

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"go-micro.dev/v4/server"

	"gitlab.51baibao.com/server/mic/v4/ut"
)

// GrpcErrLogWrapper serverWrapper to print err stack
func GrpcErrLogWrapper(h server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		err := h(ctx, req, rsp)
		if err != nil {
			me, ok := ut.ParseMicError(err)
			if ok && me.Code != 500 {
				return err
			}

			logrus.WithError(err).Errorf("rpc panic err %s.%s \n %+v\n", req.Service(), req.Endpoint(), err)
			fmt.Printf("err %+v\n", err)
		}
		return err
	}
}
