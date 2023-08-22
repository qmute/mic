package internal

import (
	"context"
	"fmt"
	"runtime/debug"

	"github.com/sirupsen/logrus"
	"go-micro.dev/v4/server"
)

// SubscribePanicWrapper 包装事件订阅，防止panic
func SubscribePanicWrapper(next server.SubscriberFunc) server.SubscriberFunc {
	return func(ctx context.Context, msg server.Message) (err error) {
		defer func() {
			if e := recover(); e != nil {
				err = fmt.Errorf("panic %+v\n", e)
				logrus.WithField("stack", string(debug.Stack())).Errorf("subscribe panic recovered from %s \n %+v\n", msg.Topic(), err)
				fmt.Printf("panic %+v\n%s\n", e, string(debug.Stack()))
			}
		}()
		return next(ctx, msg)
	}
}

// SubscribeErrLogWrapper 包装事件订阅，出错时打panic日志(带stack)
func SubscribeErrLogWrapper(next server.SubscriberFunc) server.SubscriberFunc {
	return func(ctx context.Context, msg server.Message) error {
		err := next(ctx, msg)
		if err != nil {
			logrus.WithField("stack", string(debug.Stack())).Errorf("subscribe panic err from %s \n %+v\n", msg.Topic(), err)
			fmt.Printf("panic %+v\n%s\n", err, string(debug.Stack()))
		}
		return err
	}
}
