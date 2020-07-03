package xdefaultLogger

import (
	"github.com/go-xe2/x/core/logger"
	"github.com/go-xe2/x/os/xlog"
)

func New(options ...*TLoggerOptions) logger.ILogger {
	mLogger := xlog.Clone()
	if len(options) > 0 {
		opt := options[0]
		mLogger.SetFlags(opt.GetFlags())
		mLogger.SetAsync(opt.GetAsync())
		mLogger.SetDebug(opt.GetDebug())
		mLogger.SetLevel(opt.GetLevel())
		mLogger.SetStack(opt.GetEnableStack())
		mLogger.SetStackSkip(opt.GetStackSkip())
		mLogger.SetWriter(opt.GetWriter())

	}
	return &TDefaultLogger{
		logger: mLogger,
	}
}
