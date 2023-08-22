package mic

import (
	"log"

	"go-micro.dev/v4/server"
)

// RabbitMQDurableQueue 开启rabbitMQ的持久化订阅
//
// Example:
//
//	micro.RegisterSubscriber("foo", server, hdl, mic.RabbitMQDurableQueue("bar"));
func RabbitMQDurableQueue(name string) server.SubscriberOption {
	if name == "" {
		log.Fatal("DurableQueue doesn't work with empty name")
	}
	// 为实现可靠订阅， 以下几项必须同时使用
	fName := server.SubscriberQueue(name) // 固定名字
	// fDurable := rmq.ServerDurableQueue()       // 队列持久化
	fDisableAutoAck := server.DisableAutoAck() // 禁用自动ack（同时影响mq connection 和 broker 的处理逻辑）
	// fAckOnSuccess := rmq.ServerAckOnSuccess()  // 确认成功后才ack
	return func(o *server.SubscriberOptions) {
		fName(o)
		// fDurable(o)
		fDisableAutoAck(o)
		// fAckOnSuccess(o)
	}
}

// RabbitMQDurableMessageContext 发布消息时使用此context， 可以确保消费是持久化的
//
// Example:
//
//	ctx := mic.RabbitMQDurableMessageContext(context.Background())
//	publisher.Publish(ctx, msg)
// var RabbitMQDurableMessageContext = rmq.DurableMessageContext
