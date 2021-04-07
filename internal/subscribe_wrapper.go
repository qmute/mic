package internal

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2/server"
	"github.com/sirupsen/logrus"
)

// SubPanicWrapper 包装事件订阅，防止panic
func SubPanicWrapper(next server.SubscriberFunc) server.SubscriberFunc {
	return func(ctx context.Context, msg server.Message) (err error) {
		defer func() {
			if e := recover(); e != nil {
				logrus.WithError(err).Errorf("subscribe panic recovered from %s \n %+v\n", msg.Topic(), err)
				fmt.Printf("panic %+v\n", e)
				err = fmt.Errorf("panic %+v\n", e)
			}
		}()
		return next(ctx, msg)
	}
}

// SubErrLogWrapper 包装事件订阅，出错时打日志(带stack)
func SubErrLogWrapper(next server.SubscriberFunc) server.SubscriberFunc {
	return func(ctx context.Context, msg server.Message) error {
		err := next(ctx, msg)
		if err != nil {
			logrus.WithError(err).Errorf("subscribe err from %s \n %+v\n", msg.Topic(), err)
			fmt.Printf("%+v\n", err)
		}
		return err
	}
}
