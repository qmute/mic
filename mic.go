package mic

import (
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/env"
)

// 针对go-micro 框架的工具库和进一步封装， 降低使用门槛和学习成本

// 主入口见 DefaultService

// DefaultConfig 最常用的config方式， 从环境变量中读取
func DefaultConfig() (config.Config, error) {
	conf := config.NewConfig()
	err := conf.Load(env.NewSource())
	return conf, err
}
