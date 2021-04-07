package internal

import (
	"context"
	"fmt"
	"runtime/debug"

	"github.com/micro/go-micro/v2/server"
	"github.com/sirupsen/logrus"
)

// SubscribePanicWrapper 包装事件订阅，防止panic
func SubscribePanicWrapper(next server.SubscriberFunc) server.SubscriberFunc {
	return func(ctx context.Context, msg server.Message) (err error) {
		defer func() {
			if e := recover(); e != nil {
				err = fmt.Errorf("panic %+v\n", e)
				logrus.WithError(err).WithField("stack", string(debug.Stack())).Errorf("subscribe panic recovered from %s \n %+v\n", msg.Topic(), err)
				fmt.Printf("panic %+v\n%s\n", e, string(debug.Stack()))
			}
		}()
		return next(ctx, msg)
	}
}

// SubscribeErrLogWrapper 包装事件订阅，出错时打日志(带stack)
func SubscribeErrLogWrapper(next server.SubscriberFunc) server.SubscriberFunc {
	return func(ctx context.Context, msg server.Message) error {
		err := next(ctx, msg)
		if err != nil {
			logrus.WithError(err).Errorf("subscribe err from %s \n %+v\n", msg.Topic(), err)
			fmt.Printf("%+v\n", err)
		}
		return err
	}
}
