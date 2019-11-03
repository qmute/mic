package mic

import (
	"log"

	"github.com/micro/go-micro/server"
	"github.com/quexer/go-plugins/broker/rabbitmq"
)

// RabbitMQDurableQueue 开启rabbitMQ的持久化订阅
//
// Example:
//
//   micro.RegisterSubscriber("foo", server, hdl, mic.RabbitMQDurableQueue("bar"));
func RabbitMQDurableQueue(name string) server.SubscriberOption {
	if name == "" {
		log.Fatal("DurableQueue doesn't work with empty name")
	}
	fName := server.SubscriberQueue(name)
	fDurable := rabbitmq.ServerDurableQueue()
	return func(o *server.SubscriberOptions) {
		fName(o)
		fDurable(o)
	}

}

// RabbitMQDurableMessageContext
// 发布消息时使用此context， 可以确保消费是持久化的
//
// Example:
//	ctx := mic.RabbitMQDurableMessageContext(context.Background())
//	publisher.Publish(ctx, msg)
var RabbitMQDurableMessageContext = rabbitmq.DurableMessageContext
