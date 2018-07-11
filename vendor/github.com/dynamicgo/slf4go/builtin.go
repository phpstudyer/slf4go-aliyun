package slf4go

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"

	config "github.com/dynamicgo/go-config"
	"github.com/fatih/color"
)

var fatalp = color.New(color.FgRed).PrintFunc()
var fatalf = color.New(color.FgRed).PrintfFunc()

var errorp = color.New(color.FgRed).PrintFunc()
var errorf = color.New(color.FgRed).PrintfFunc()

var warnp = color.New(color.FgYellow).PrintFunc()
var warnf = color.New(color.FgYellow).PrintfFunc()

var infop = color.New(color.FgWhite).PrintFunc()
var infof = color.New(color.FgWhite).PrintfFunc()

var debugp = color.New(color.FgCyan).PrintFunc()
var debugf = color.New(color.FgCyan).PrintfFunc()

var tracep = color.New(color.FgBlue).PrintFunc()
var tracef = color.New(color.FgBlue).PrintfFunc()

type colorConsole struct {
}

func newColorConsole() LoggerFactory {
	return &colorConsole{}
}

func (console *colorConsole) GetLogger(name string) Logger {
	return &colorConsoleLogger{name: name}
}

type colorConsoleLogger struct {
	name string
}

func (logger *colorConsoleLogger) GetName() string {
	return logger.name
}

func source() string {
	_, filename, line, _ := runtime.Caller(3)

	return fmt.Sprintf("%s:%d", filepath.Base(filename), line)
}

func (logger *colorConsoleLogger) Trace(args ...interface{}) {
	tracef("[%s][%s][%s] TRACE ", time.Now().Format("2006-01-02 15:04:05"), logger.name, source())
	tracep(args...)
	tracep("\n")
}

func (logger *colorConsoleLogger) TraceF(format string, args ...interface{}) {
	tracef("[%s][%s][%s] TRACE ", time.Now().Format("2006-01-02 15:04:05"), logger.name, source())
	tracef(format, args...)
	tracep("\n")
}

func (logger *colorConsoleLogger) Debug(args ...interface{}) {
	debugf("[%s][%s][%s] DEBUG ", time.Now().Format("2006-01-02 15:04:05"), logger.name, source())
	debugp(args...)
	debugp("\n")
}

func (logger *colorConsoleLogger) DebugF(format string, args ...interface{}) {
	debugf("[%s][%s][%s] DEBUG ", time.Now().Format("2006-01-02 15:04:05"), logger.name, source())
	debugf(format, args...)
	debugp("\n")
}

func (logger *colorConsoleLogger) Info(args ...interface{}) {
	infof("[%s][%s][%s] INFO  ", time.Now().Format("2006-01-02 15:04:05"), logger.name, source())
	infop(args...)
	infop("\n")
}

func (logger *colorConsoleLogger) InfoF(format string, args ...interface{}) {
	infof("[%s][%s][%s] INFO  ", time.Now().Format("2006-01-02 15:04:05"), logger.name, source())
	infof(format, args...)
	infop("\n")
}

func (logger *colorConsoleLogger) Warn(args ...interface{}) {
	warnf("[%s][%s][%s] WARN  ", time.Now().Format("2006-01-02 15:04:05"), logger.name, source())
	warnp(args...)
	warnp("\n")
}

func (logger *colorConsoleLogger) WarnF(format string, args ...interface{}) {
	warnf("[%s][%s][%s] WARN  ", time.Now().Format("2006-01-02 15:04:05"), logger.name, source())
	warnf(format, args...)
	warnp("\n")
}

func (logger *colorConsoleLogger) Error(args ...interface{}) {
	errorf("[%s][%s][%s] ERROR ", time.Now().Format("2006-01-02 15:04:05"), logger.name, source())
	errorp(args...)
	errorp("\n")
}

func (logger *colorConsoleLogger) ErrorF(format string, args ...interface{}) {
	errorf("[%s][%s][%s] ERROR ", time.Now().Format("2006-01-02 15:04:05"), logger.name, source())
	errorf(format, args...)
	errorp("\n")
}

func (logger *colorConsoleLogger) Fatal(args ...interface{}) {
	fatalf("[%s][%s][%s] FATAL ", time.Now().Format("2006-01-02 15:04:05"), logger.name, source())
	fatalp(args...)
	fatalp("\n")
}

func (logger *colorConsoleLogger) FatalF(format string, args ...interface{}) {
	fatalf("[%s][%s][%s] FATAL ", time.Now().Format("2006-01-02 15:04:05"), logger.name, source())
	fatalf(format, args...)
	fatalp("\n")
}

func init() {
	println("[slf4go] register console backend")
	RegisterBackend("console", func(config config.Config) (LoggerFactory, error) {
		return newColorConsole(), nil
	})
}
