package internal

import (
	"context"
	"fmt"
	"runtime/debug"

	"github.com/micro/go-micro/v2/server"
	"github.com/sirupsen/logrus"

	"gitlab.51baibao.com/server/mic/ut"
)

// GrpcRecoveryWrapper,  serverWrapper to avoid crash
func GrpcRecoveryWrapper(h server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) (err error) {
		defer func() {
			if e := recover(); e != nil {
				err = fmt.Errorf("panic %+v\n", e)
				logrus.WithError(err).WithField("stack", string(debug.Stack())).Errorf("rpc panic recovered %s.%s \n %+v\n", req.Service(), req.Endpoint(), err)
				fmt.Printf("panic %+v\n%s\n", e, string(debug.Stack()))
			}
		}()
		return h(ctx, req, rsp)
	}
}

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
