package mic_test

import (
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	"go-micro.dev/v5/logger"

	"github.com/qmute/mic/v5"
)

var _ = Describe("LogInitializer", func() {
	It("debug", func() {
		gin.SetMode(gin.DebugMode)
		mic.InitLogger()
		logger.Info("debug")
	})
	It("prod", func() {
		gin.SetMode(gin.ReleaseMode)
		mic.InitLogger()
		logger.Info("release")
	})
})
