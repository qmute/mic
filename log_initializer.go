package mic

import (
	"github.com/gin-gonic/gin"
	logPlugin "github.com/go-micro/plugins/v4/logger/logrus"
	"github.com/sirupsen/logrus"
	"go-micro.dev/v4/logger"
)

// InitLogger 日志初始化
func InitLogger() {
	const tsFmt = "2006-01-02 15:04:05"
	if gin.Mode() == gin.DebugMode {
		logrus.SetFormatter(&logrus.TextFormatter{TimestampFormat: tsFmt, FullTimestamp: true})
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: tsFmt})
	}
	logger.DefaultLogger = logPlugin.NewLogger(logPlugin.WithLogger(logrus.StandardLogger()))
}
