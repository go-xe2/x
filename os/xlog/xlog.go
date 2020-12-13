package xlog

import (
	"github.com/go-xe2/x/core/cmdenv"
	"github.com/go-xe2/x/os/xrunPool"
	"io"
)

const (
	LEVEL_ALL  = LEVEL_DEBU | LEVEL_INFO | LEVEL_NOTI | LEVEL_WARN | LEVEL_ERRO | LEVEL_CRIT
	LEVEL_DEV  = LEVEL_ALL
	LEVEL_PROD = LEVEL_WARN | LEVEL_ERRO | LEVEL_CRIT
	LEVEL_DEBU = 1 << iota
	LEVEL_INFO
	LEVEL_WARN
	LEVEL_ERRO
	LEVEL_NOTI
	LEVEL_CRIT
)

var (
	logger    = New()
	asyncPool = xrunPool.New(1)
)

func init() {
	SetDebug(cmdenv.Get("x.xlog.debug", true).Bool())
}

func Logger() *TLogger {
	return logger
}

func Clone() *TLogger {
	return logger.Clone()
}

// 设置日志路径
func SetPath(path string) error {
	return logger.SetPath(path)
}

// 获取日志路径
func GetPath() string {
	return logger.GetPath()
}

// 设置日志文件过滤串
// 默认过滤为:Y-m-d.log, eg: 2018-01-01.log
func SetFile(pattern string) {
	logger.SetFile(pattern)
}

// SetLevel sets the default logging level.
// 设置日志级别
func SetLevel(level int) {
	logger.SetLevel(level)
}

// 获取日志级别
func GetLevel() int {
	return logger.GetLevel()
}

// 设置日志输出接口
func SetWriter(writer io.Writer) {
	logger.SetWriter(writer)
}

// 获取日志输出接口
func GetWriter() io.Writer {
	return logger.GetWriter()
}

// 测试日志为高工模式
func SetDebug(debug bool) {
	logger.SetDebug(debug)
}

func GetDebug() bool {
	return logger.GetDebug()
}

// 设置是否启用异步输出日志
func SetAsync(enabled bool) {
	logger.SetAsync(enabled)
}

func GetAsync() bool {
	return logger.GetAsync()
}

// 设置是否为stdout输出，默认为true
func SetStdoutPrint(enabled bool) {
	logger.SetStdoutPrint(enabled)
}

// 设置是否打印日志头信息，默认为gtrue
func SetHeaderPrint(enabled bool) {
	logger.SetHeaderPrint(enabled)
}

// 设置日志头前缀
func SetPrefix(prefix string) {
	logger.SetPrefix(prefix)
}

func GetPrefix() string {
	return logger.GetPrefix()
}

// 设置日志扩展设置
func SetFlags(flags int) {
	logger.SetFlags(flags)
}

// 获取日志扩展设置
func GetFlags() int {
	return logger.GetFlags()
}

// 打印日志栈
func PrintStack(skip ...int) {
	logger.PrintStack(skip...)
}

// 获取日志栈
func GetStack(skip ...int) string {
	return logger.GetStack(skip...)
}

// 是否打印日志栈信息
func SetStack(enabled bool) {
	logger.SetStack(enabled)
}

func GetStackStatus() bool {
	return logger.GetStackStatus()
}

func GetStackSkip() int {
	return logger.GetStackSkip()
}
