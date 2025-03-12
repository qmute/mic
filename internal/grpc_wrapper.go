package internal

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"go-micro.dev/v4/server"
)

// GrpcErrLogWrapper serverWrapper to print err stack
func GrpcErrLogWrapper(h server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		// 不能在包将错误返回，使用方有可能直接使用detail字段给用户返回
		err := h(ctx, req, rsp)
		if err != nil {
			// me, ok := ut.ParseMicError(err)
			// if ok && me.Code != 500 {
			// 	return err
			// }

			logrus.WithError(err).Errorf("rpc panic err %s.%s \n %+v\n", req.Service(), req.Endpoint(), err)
			fmt.Printf("err %+v\n", err)
		}
		return err
	}
}
