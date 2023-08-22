package mic_test

import (
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	"go-micro.dev/v4/logger"

	"gitlab.51baibao.com/server/mic/v4"
)

var _ = Describe("LogInitializer", func() {
	It("debug", func() {
		gin.SetMode(gin.DebugMode)
		_ = (&mic.LogInitializer{}).Initialize(ctx)
		logger.Info("debug")
	})
	It("prod", func() {
		gin.SetMode(gin.ReleaseMode)
		_ = (&mic.LogInitializer{}).Initialize(ctx)
		logger.Info("release")
	})
})
