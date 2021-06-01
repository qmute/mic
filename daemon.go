package mic

import (
	"fmt"
	"runtime/debug"

	log "github.com/micro/go-micro/v2/logger"
	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
)

// Daemon Daemon接口
type Daemon interface {
	// Job 定义job，返回 cron 字符串和执行函数。
	// 如果cron字符串为空，则由用户自己控制运行频率。 框架只负责新启一个线程把"死循环"运行起来
	Job() (string, func())
}

// InitDaemon 初始化传入的 Daemon
func InitDaemon(it ...Daemon) error {
	c := newCron()
	for _, v := range it {
		spec, fn := v.Job()
		// 定义了cron， 则加入执行引擎
		if spec != "" {
			if _, err := c.AddFunc(spec, fn); err != nil {
				return errors.WithStack(err)
			}
			continue
		}

		// 没有定义cron， 则启动单独线程运行， 并作panic保护
		go func(f func()) {
			defer func() {
				if e := recover(); e != nil {
					stack := string(debug.Stack())
					log.Fields(map[string]interface{}{"stack": stack}).Log(log.ErrorLevel, "job panic", e)
					fmt.Printf("job panic %+v\n%s\n", e, stack)
				}
			}()
			f()
		}(fn)
	}
	c.Start()
	return nil
}

// 创建Cron
func newCron() *cron.Cron {
	logger := &daemonLogger{}
	chain := cron.WithChain(
		cron.Recover(logger),
		cron.DelayIfStillRunning(logger),
	)
	return cron.New(cron.WithLogger(logger), chain)
}

type daemonLogger struct {
}

func (p *daemonLogger) Info(msg string, keysAndValues ...interface{}) {
	arg := []interface{}{msg}
	log.Info(append(arg, keysAndValues...))
}

func (p *daemonLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	arg := []interface{}{msg}
	log.Error(append(arg, keysAndValues...))
	log.Errorf("%+v", err)
}
