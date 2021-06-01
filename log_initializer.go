package mic

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/logger"
	logPlugin "github.com/micro/go-plugins/logger/logrus/v2"
	"github.com/sirupsen/logrus"
)

// LogInitializer 日志初始化器
type LogInitializer struct {
}

func (p *LogInitializer) Name() string {
	return "log_initializer"
}

func (p *LogInitializer) IsNeedInit(ctx context.Context) (bool, error) {
	return true, nil
}

func (p *LogInitializer) Initialize(ctx context.Context) error {
	const tsFmt = "2006-01-02 15:04:05"
	if gin.Mode() == gin.DebugMode {
		logrus.SetFormatter(&logrus.TextFormatter{TimestampFormat: tsFmt, FullTimestamp: true})
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: tsFmt})
	}

	logger.DefaultLogger = logPlugin.NewLogger(logPlugin.WithLogger(logrus.StandardLogger()))
	return nil
}
