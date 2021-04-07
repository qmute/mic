package internal

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2/server"
	"github.com/sirupsen/logrus"
)

// GrpcRecoveryWrapper,  serverWrapper to avoid crash
func GrpcRecoveryWrapper(h server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) (err error) {
		defer func() {
			if e := recover(); e != nil {
				err = fmt.Errorf("panic %+v\n", e)
				logrus.WithError(err).Errorf("panic recovered %s.%s \n %+v\n", req.Service(), req.Endpoint(), err)
				fmt.Printf("panic %+v\n", e)
			}
		}()
		return h(ctx, req, rsp)
	}
}

// GrpcErrLogWrapper,  serverWrapper to print err stack
func GrpcErrLogWrapper(h server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		err := h(ctx, req, rsp)
		if err != nil {
			logrus.WithError(err).Errorf("err %s.%s \n %+v\n", req.Service(), req.Endpoint(), err)
			fmt.Printf("err %+v\n", err)
		}
		return err
	}
}
