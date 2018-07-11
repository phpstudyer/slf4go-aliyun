package aliyun

import (
	config "github.com/dynamicgo/go-config"
	"github.com/dynamicgo/slf4go"
)

type loggerFactory struct {
}

func newLoggerFactory(config config.Config) (slf4go.LoggerFactory, error) {
	return &loggerFactory{}, nil
}

func (factory *loggerFactory) GetLogger(name string) slf4go.Logger {
	return nil
}

func init() {
	println("[slf4go] register aliyun backend")
	slf4go.RegisterBackend("aliyun", func(config config.Config) (slf4go.LoggerFactory, error) {
		return newLoggerFactory(config)
	})
}
