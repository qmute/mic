package internal

import (
	"context"
	"fmt"
	"runtime/debug"

	"github.com/sirupsen/logrus"
	"go-micro.dev/v5/server"
)

// SubscribeErrLogWrapper 包装事件订阅，出错时打panic日志(带stack)
func SubscribeErrLogWrapper(next server.SubscriberFunc) server.SubscriberFunc {
	return func(ctx context.Context, msg server.Message) error {
		err := next(ctx, msg)
		if err != nil {
			logrus.WithField("stack", string(debug.Stack())).Errorf("subscribe panic err from %s \n %+v\n playload: %#v \n raw body: %s", msg.Topic(), err, msg.Payload(), string(msg.Body()))
			fmt.Printf("panic %+v\n%s\n", err, string(debug.Stack()))
		}
		return err
	}
}
