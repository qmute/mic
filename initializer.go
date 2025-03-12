package mic

import (
	"context"

	"github.com/pkg/errors"
	"go-micro.dev/v4/logger"
)

// Initializer 初始化器
type Initializer interface {
	// Name 初始化器的名称
	Name() string
	// IsNeedInit 是否需要初始化
	IsNeedInit(ctx context.Context) (bool, error)
	// Initialize 初始化数据
	Initialize(ctx context.Context) error
}

// InitAll 初始化传入的Initializer
func InitAll(ctx context.Context, it ...Initializer) error {
	for _, v := range it {
		name := v.Name()
		need, err := v.IsNeedInit(ctx)
		if err != nil {
			return err
		}

		log := logger.Fields(map[string]interface{}{"name": name})

		if !need {
			log.Log(logger.InfoLevel, "init not need")
			continue
		}

		err = v.Initialize(ctx)
		if err != nil {
			return errors.WithMessagef(err, "%s init", name)
		}
		log.Log(logger.InfoLevel, "init completed")
	}
	return nil
}
