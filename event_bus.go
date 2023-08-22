package mic

import (
	"context"
	"fmt"
	"sync"

	"go-micro.dev/v4"
)

// EventPub事件发布接口
type EventPub interface {
	Pub(ctx context.Context, topic string, msg interface{}) error
}

// EventSub 事件订阅接口
type EventSub interface {
	// queue，可选的 订阅队列名称。
	// 当不同业务方订阅同一个topic时，需要各自指定不同队列名
	Sub(topic string, msg interface{}, queue ...string) error
}

// MicroEventBus micro 事件总线，提供对 micro 消息系统的简单封装
type MicroEventBus struct {
	sync.Mutex
	pubMap       map[string]micro.Event
	microService micro.Service
}

// NewMicroEventBus 新建micro事件总线
func NewMicroEventBus(service micro.Service) *MicroEventBus {
	return &MicroEventBus{
		pubMap:       map[string]micro.Event{},
		microService: service,
	}
}

// Pub 发布消息
func (p *MicroEventBus) Pub(ctx context.Context, topic string, msg interface{}) error {
	// 目前的应用场景全是持久化的。  若以后遇到其它情况，到时再扩展接口
	// return p.getPublisher(topic).Publish(RabbitMQDurableMessageContext(ctx), msg)
	return p.getPublisher(topic).Publish(ctx, msg)
}

func (p *MicroEventBus) getPublisher(topic string) micro.Event {
	p.Lock()
	defer p.Unlock()

	pub := p.pubMap[topic]
	if pub == nil {
		pub = micro.NewEvent(topic, p.microService.Client())
		p.pubMap[topic] = pub
	}
	return pub
}

// Sub 订阅消息
func (p *MicroEventBus) Sub(topic string, hdl interface{}, queue ...string) error {
	name := topic
	if len(queue) > 0 {
		name = fmt.Sprintf("%s:%s", topic, queue[0])
	}
	return micro.RegisterSubscriber(topic, p.microService.Server(), hdl, RabbitMQDurableQueue(name))
}
