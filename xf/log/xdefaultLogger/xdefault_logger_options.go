package xdefaultLogger

import (
	"github.com/go-xe2/x/os/xlog"
	"io"
)

type TLoggerOptions struct {
	// 设置是否为调试模式
	debug bool
	// 设置日志级别
	level int
	// 设置是否启用异步输出日志
	async bool
	// 设置日扩展标识
	flags int
	// 设置是否启用输出日志栈
	enableStack bool
	//设置输出日志栈层级数
	stackSkip int
	writer    io.Writer
}

func NewLoggerOptions() *TLoggerOptions {
	return &TLoggerOptions{
		debug:       xlog.GetDebug(),
		level:       xlog.GetLevel(),
		async:       xlog.GetAsync(),
		flags:       xlog.GetFlags(),
		enableStack: xlog.GetStackStatus(),
		stackSkip:   xlog.GetStackSkip(),
		writer:      xlog.GetWriter(),
	}
}

func (lo *TLoggerOptions) SetDebug(debug bool) *TLoggerOptions {
	lo.debug = debug
	return lo
}

func (lo *TLoggerOptions) SetLevel(level int) *TLoggerOptions {
	lo.level = level
	return lo
}

func (lo *TLoggerOptions) SetAsync(async bool) *TLoggerOptions {
	lo.async = async
	return lo
}

func (lo *TLoggerOptions) SetFlags(flags int) *TLoggerOptions {
	lo.flags = flags
	return lo
}

func (lo *TLoggerOptions) SetEnableStack(enable bool) *TLoggerOptions {
	lo.enableStack = enable
	return lo
}

func (lo *TLoggerOptions) SetStackSkip(skip int) *TLoggerOptions {
	lo.stackSkip = skip
	return lo
}

func (lo *TLoggerOptions) SetWriter(writer io.Writer) *TLoggerOptions {
	lo.writer = writer
	return lo
}

func (lo *TLoggerOptions) GetDebug() bool {
	return lo.debug
}

func (lo *TLoggerOptions) GetLevel() int {
	return lo.level
}

func (lo *TLoggerOptions) GetAsync() bool {
	return lo.async
}

func (lo *TLoggerOptions) GetFlags() int {
	return lo.flags
}

func (lo *TLoggerOptions) GetEnableStack() bool {
	return lo.enableStack
}

func (lo *TLoggerOptions) GetStackSkip() int {
	return lo.stackSkip
}

func (lo *TLoggerOptions) GetWriter() io.Writer {
	return lo.writer
}
