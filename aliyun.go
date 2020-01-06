package aliyun

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"

	sls "github.com/aliyun/aliyun-log-go-sdk"
	config "github.com/dynamicgo/go-config"
	"github.com/dynamicgo/slf4go"
	"github.com/gogo/protobuf/proto"
)

var (
  	PutLogInterval = time.Second * 5
	PutLogMaxCount = 100
	MaxCount = 10000
)

type loggerFactory struct {
	source         string
	project        *sls.LogProject
	logstore       *sls.LogStore
	cached         int
	putLogInterval time.Duration
	putLogMaxCount int
}

func newLoggerFactory(config config.Config) (slf4go.LoggerFactory, error) {

	project := &sls.LogProject{
		Name:            config.Get("project").String(""),
		Endpoint:        config.Get("endpoint").String(""),
		AccessKeyID:     config.Get("accesskey", "id").String(""),
		AccessKeySecret: config.Get("accesskey", "secret").String(""),
	}

	logstore, err := project.GetLogStore(config.Get("logstore").String(""))

	if err != nil {
		return nil, err
	}

	return &loggerFactory{
		project:        project,
		logstore:       logstore,
		source:         config.Get("source").String(""),
		cached:         MaxCount,
		putLogInterval: PutLogInterval,
		putLogMaxCount: PutLogMaxCount,
	}, nil
}

func (factory *loggerFactory) GetLogger(name string) slf4go.Logger {
	log := &aliyunLog{
		topic:          name,
		source:         factory.source,
		logstore:       factory.logstore,
		mq:             make(chan []*sls.LogContent, factory.cached),
		codelevel:      3,
		putLogInterval: factory.putLogInterval,
		putLogMaxCount: factory.putLogMaxCount,
	}

	go log.runLoop()

	return log
}

type aliyunLog struct {
	topic          string
	source         string
	logstore       *sls.LogStore
	mq             chan []*sls.LogContent
	codelevel      int
	putLogInterval time.Duration
	putLogMaxCount int
}

func (l *aliyunLog) runLoop() {
	now := time.Now()
	group := &sls.LogGroup{
		Topic:  proto.String(l.topic),
		Source: proto.String(l.source),
		Logs:   []*sls.Log{},
	}
	(now.Add(l.putLogInterval).Before(time.Now()))

	ticker := time.NewTicker(l.putLogInterval)
	defer ticker.Stop()
	go func() {
		for range ticker.C {
			if len(group.Logs) > 0 {
				err := l.logstore.PutLogs(group)
				if err != nil {
					fmt.Printf("logstore put logs err, %s\n", err)
				} else {
					group.Logs = []*sls.Log{}
				}
			}
		}
	}()

	for content := range l.mq {
		count := len(group.Logs)
		if count >= l.putLogMaxCount {
			log.Printf("group %+v", group)
			if err := l.logstore.PutLogs(group); err != nil {
				fmt.Printf("logstore put logs err, %s\n", err)
				continue
			}
			group.Logs = []*sls.Log{}
		}

		group.Logs = append(group.Logs, &sls.Log{
			Contents: content,
			Time:     proto.Uint32(uint32(time.Now().Unix())),
		})

	}
}

func (logger *aliyunLog) SourceCodeLevel(level int) {
	logger.codelevel = level
}

func (logger *aliyunLog) GetName() string {
	return logger.topic
}

func (logger *aliyunLog) Source() string {
	_, filename, line, _ := runtime.Caller(logger.codelevel)

	return fmt.Sprintf("%s:%d", filepath.Base(filename), line)
}

func (logger *aliyunLog) Trace(args ...interface{}) {
	logger.mq <- []*sls.LogContent{

		&sls.LogContent{
			Key:   proto.String("Level"),
			Value: proto.String("Trace"),
		},
		&sls.LogContent{
			Key:   proto.String("File"),
			Value: proto.String(logger.Source()),
		},
		&sls.LogContent{
			Key:   proto.String("Content"),
			Value: proto.String(fmt.Sprint(args...)),
		},
	}
}

func (logger *aliyunLog) TraceF(format string, args ...interface{}) {
	logger.mq <- []*sls.LogContent{

		&sls.LogContent{
			Key:   proto.String("Level"),
			Value: proto.String("Trace"),
		},
		&sls.LogContent{
			Key:   proto.String("File"),
			Value: proto.String(logger.Source()),
		},
		&sls.LogContent{
			Key:   proto.String("Content"),
			Value: proto.String(fmt.Sprintf(format, args...)),
		},
	}
}

func (logger *aliyunLog) Debug(args ...interface{}) {
	logger.mq <- []*sls.LogContent{

		&sls.LogContent{
			Key:   proto.String("Level"),
			Value: proto.String("Debug"),
		},
		&sls.LogContent{
			Key:   proto.String("File"),
			Value: proto.String(logger.Source()),
		},
		&sls.LogContent{
			Key:   proto.String("Content"),
			Value: proto.String(fmt.Sprint(args...)),
		},
	}
}

func (logger *aliyunLog) DebugF(format string, args ...interface{}) {
	logger.mq <- []*sls.LogContent{

		&sls.LogContent{
			Key:   proto.String("Level"),
			Value: proto.String("Debug"),
		},
		&sls.LogContent{
			Key:   proto.String("File"),
			Value: proto.String(logger.Source()),
		},
		&sls.LogContent{
			Key:   proto.String("Content"),
			Value: proto.String(fmt.Sprintf(format, args...)),
		},
	}
}

func (logger *aliyunLog) Info(args ...interface{}) {
	logger.mq <- []*sls.LogContent{

		&sls.LogContent{
			Key:   proto.String("Level"),
			Value: proto.String("Info"),
		},
		&sls.LogContent{
			Key:   proto.String("File"),
			Value: proto.String(logger.Source()),
		},
		&sls.LogContent{
			Key:   proto.String("Content"),
			Value: proto.String(fmt.Sprint(args...)),
		},
	}
}

func (logger *aliyunLog) InfoF(format string, args ...interface{}) {
	logger.mq <- []*sls.LogContent{

		&sls.LogContent{
			Key:   proto.String("Level"),
			Value: proto.String("Info"),
		},
		&sls.LogContent{
			Key:   proto.String("File"),
			Value: proto.String(logger.Source()),
		},
		&sls.LogContent{
			Key:   proto.String("Content"),
			Value: proto.String(fmt.Sprintf(format, args...)),
		},
	}
}

func (logger *aliyunLog) Warn(args ...interface{}) {
	logger.mq <- []*sls.LogContent{

		&sls.LogContent{
			Key:   proto.String("Level"),
			Value: proto.String("Warn"),
		},
		&sls.LogContent{
			Key:   proto.String("File"),
			Value: proto.String(logger.Source()),
		},
		&sls.LogContent{
			Key:   proto.String("Content"),
			Value: proto.String(fmt.Sprint(args...)),
		},
	}
}

func (logger *aliyunLog) WarnF(format string, args ...interface{}) {
	logger.mq <- []*sls.LogContent{

		&sls.LogContent{
			Key:   proto.String("Level"),
			Value: proto.String("Warn"),
		},
		&sls.LogContent{
			Key:   proto.String("File"),
			Value: proto.String(logger.Source()),
		},
		&sls.LogContent{
			Key:   proto.String("Content"),
			Value: proto.String(fmt.Sprintf(format, args...)),
		},
	}
}

func (logger *aliyunLog) Error(args ...interface{}) {
	logger.mq <- []*sls.LogContent{

		&sls.LogContent{
			Key:   proto.String("Level"),
			Value: proto.String("Error"),
		},
		&sls.LogContent{
			Key:   proto.String("File"),
			Value: proto.String(logger.Source()),
		},
		&sls.LogContent{
			Key:   proto.String("Content"),
			Value: proto.String(fmt.Sprint(args...)),
		},
	}
}

func (logger *aliyunLog) ErrorF(format string, args ...interface{}) {
	logger.mq <- []*sls.LogContent{

		&sls.LogContent{
			Key:   proto.String("Level"),
			Value: proto.String("Error"),
		},
		&sls.LogContent{
			Key:   proto.String("File"),
			Value: proto.String(logger.Source()),
		},
		&sls.LogContent{
			Key:   proto.String("Content"),
			Value: proto.String(fmt.Sprintf(format, args...)),
		},
	}
}

func (logger *aliyunLog) Fatal(args ...interface{}) {
	logger.mq <- []*sls.LogContent{

		&sls.LogContent{
			Key:   proto.String("Level"),
			Value: proto.String("Fatal"),
		},
		&sls.LogContent{
			Key:   proto.String("File"),
			Value: proto.String(logger.Source()),
		},
		&sls.LogContent{
			Key:   proto.String("Content"),
			Value: proto.String(fmt.Sprint(args...)),
		},
	}
}

func (logger *aliyunLog) FatalF(format string, args ...interface{}) {
	logger.mq <- []*sls.LogContent{

		&sls.LogContent{
			Key:   proto.String("Level"),
			Value: proto.String("Fatal"),
		},
		&sls.LogContent{
			Key:   proto.String("File"),
			Value: proto.String(logger.Source()),
		},
		&sls.LogContent{
			Key:   proto.String("Content"),
			Value: proto.String(fmt.Sprintf(format, args...)),
		},
	}
}

func init() {
	println("[slf4go] register aliyun backend")
	slf4go.RegisterBackend("aliyun", func(config config.Config) (slf4go.LoggerFactory, error) {
		return newLoggerFactory(config)
	})
}
