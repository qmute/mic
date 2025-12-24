package mic

import (
	"context"
	"fmt"
	"runtime/debug"

	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
	log "go-micro.dev/v5/logger"
)

// Daemon Daemon接口
type Daemon interface {
	// Job 定义job，返回 cron 字符串和执行函数。
	// 如果cron字符串为空，则由用户自己控制运行频率。 框架只负责新启一个线程把"死循环"运行起来
	Job() (string, func(context.Context) error)
}

// InitDaemon 初始化传入的 Daemon
func InitDaemon(it ...Daemon) error {
	c := newCron()
	for _, v := range it {
		spec, fn := v.Job()
		// 定义了cron， 则加入执行引擎
		if spec != "" {
			if _, err := c.AddFunc(spec, func() {
				if err := fn(context.Background()); err != nil {
					log.Logf(log.ErrorLevel, "job panic %+v", err)
				}
			}); err != nil {
				return errors.WithStack(err)
			}
			continue
		}

		// 没有定义cron， 则启动单独线程运行， 并作panic保护
		go func(f func(context.Context) error) {
			defer daemonRecoverGuard()
			if err := f(context.Background()); err != nil {
				log.Logf(log.ErrorLevel, "job panic %+v", err)
			}
		}(fn)
	}
	c.Start()
	return nil
}

// 创建Cron
func newCron() *cron.Cron {
	logger := &daemonLogger{}
	chain := cron.WithChain(
		recoverWrapper(logger),
		cron.DelayIfStillRunning(logger),
	)
	return cron.New(cron.WithLogger(logger), chain)
}

// Recover panics in wrapped jobs and log them with the provided logger.
func recoverWrapper(logger cron.Logger) cron.JobWrapper {
	return func(job cron.Job) cron.Job {
		return cron.FuncJob(func() {
			defer daemonRecoverGuard()
			job.Run()
		})
	}
}

func daemonRecoverGuard() {
	if e := recover(); e != nil {
		err, ok := e.(error)
		if !ok {
			err = errors.Errorf("%+v", err)
		}
		stack := string(debug.Stack())
		log.Fields(map[string]interface{}{"stack": stack}).Logf(log.ErrorLevel, "job panic %+v", err)
		fmt.Printf("job panic %+v\n%s\n", err, stack)
	}
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
