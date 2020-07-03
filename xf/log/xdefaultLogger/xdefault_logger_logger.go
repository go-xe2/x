package xdefaultLogger

import (
	. "github.com/go-xe2/x/core/logger"
	"github.com/go-xe2/x/os/xlog"
)

type TDefaultLogger struct {
	logger *xlog.TLogger
}

var _ ILogger = &TDefaultLogger{}

func (lg *TDefaultLogger) log(tag string, level LogLevel, v ...interface{}) {
	lg.logger.SetPrefix(tag)
	switch level {
	case LEVEL_ERRO:
		lg.logger.Error(v...)
		break
	case LEVEL_WARN:
		lg.logger.Warning(v...)
		break
	case LEVEL_INFO:
		lg.logger.Info(v...)
		break
	case LEVEL_DEBU:
		lg.logger.Debug(v...)
		break
	case LEVEL_NOTI:
		lg.logger.Notice(v...)
		break
	case LEVEL_CRIT:
		lg.logger.Critical(v...)
		break
	}
}

func (lg *TDefaultLogger) logf(tag string, level LogLevel, format string, v ...interface{}) {
	lg.logger.SetPrefix(tag)
	switch level {
	case LEVEL_ERRO:
		lg.logger.Errorf(format, v...)
		break
	case LEVEL_WARN:
		lg.logger.Warningf(format, v...)
		break
	case LEVEL_INFO:
		lg.logger.Infof(format, v...)
		break
	case LEVEL_DEBU:
		lg.logger.Debugf(format, v...)
		break
	case LEVEL_NOTI:
		lg.logger.Noticef(format, v...)
		break
	case LEVEL_CRIT:
		lg.logger.Criticalf(format, v...)
		break
	}
}

func (lg *TDefaultLogger) logfln(tag string, level LogLevel, format string, v ...interface{}) {
	lg.logger.SetPrefix(tag)
	switch level {
	case LEVEL_ERRO:
		lg.logger.Errorfln(format, v...)
		break
	case LEVEL_WARN:
		lg.logger.Warningfln(format, v...)
		break
	case LEVEL_INFO:
		lg.logger.Infofln(format, v...)
		break
	case LEVEL_DEBU:
		lg.logger.Debugfln(format, v...)
		break
	case LEVEL_NOTI:
		lg.logger.Noticefln(format, v...)
		break
	case LEVEL_CRIT:
		lg.logger.Criticalfln(format, v...)
		break
	}
}

// 输出错误信息
func (lg *TDefaultLogger) Error(tag string, v ...interface{}) {
	lg.log(tag, LEVEL_ERRO, v...)
}

// 格式化输出错误信息
func (lg *TDefaultLogger) ErrorF(tag string, format string, v ...interface{}) {
	lg.logf(tag, LEVEL_ERRO, format, v...)
}

// 格式化输出错误信息及换行符
func (lg *TDefaultLogger) ErrorFLn(tag string, format string, v ...interface{}) {
	lg.logfln(tag, LEVEL_ERRO, format, v...)
}

// 输出警告信息
func (lg *TDefaultLogger) Warning(tag string, v ...interface{}) {
	lg.log(tag, LEVEL_WARN, v...)
}

// 格式化输出警告信息
func (lg *TDefaultLogger) WarningF(tag string, format string, v ...interface{}) {
	lg.logf(tag, LEVEL_WARN, format, v...)
}

// 格式化输出警告信息及换行
func (lg *TDefaultLogger) WarningFLn(tag string, format string, v ...interface{}) {
	lg.logfln(tag, LEVEL_WARN, format, v...)
}

// 输出调试信息
func (lg *TDefaultLogger) Debug(tag string, v ...interface{}) {
	lg.log(tag, LEVEL_DEBU, v...)
}

// 格式化输出调度信息
func (lg *TDefaultLogger) DebugF(tag string, format string, v ...interface{}) {
	lg.logf(tag, LEVEL_DEBU, format, v...)
}

// 格式化输出调度信息及换行
func (lg *TDefaultLogger) DebugFLn(tag string, format string, v ...interface{}) {
	lg.logfln(tag, LEVEL_DEBU, format, v...)
}

// 输出信息
func (lg *TDefaultLogger) Info(tag string, v ...interface{}) {
	lg.log(tag, LEVEL_INFO, v...)
}

// 格式化输出信息
func (lg *TDefaultLogger) InfoF(tag string, format string, v ...interface{}) {
	lg.logf(tag, LEVEL_INFO, format, v...)
}

// 格式输出信息及换行符
func (lg *TDefaultLogger) InfoFLn(tag string, format string, v ...interface{}) {
	lg.logfln(tag, LEVEL_INFO, format, v...)
}

// 控制台打印信息
func (lg *TDefaultLogger) Print(tag string, level uint8, v ...interface{}) {
	lg.logger.SetPrefix(tag)
	lg.logger.Print(v...)
}

// 控制台打印格式化信息
func (lg *TDefaultLogger) PrintF(tag string, level uint8, format string, v ...interface{}) {
	lg.logger.SetPrefix(tag)
	lg.logger.Printf(format, v...)
}

// 控制台打印格式化信息及换行符
func (lg *TDefaultLogger) PrintFLn(tag string, level uint8, format string, v ...interface{}) {
	lg.logger.SetPrefix(tag)
	lg.logger.Printfln(format, v...)
}

// 输出通知信息
func (lg *TDefaultLogger) Notice(tag string, v ...interface{}) {
	lg.log(tag, LEVEL_NOTI, v...)
}

// 格式输出通知信息
func (lg *TDefaultLogger) NoticeF(tag string, format string, v ...interface{}) {
	lg.logf(tag, LEVEL_NOTI, format, v...)
}

// 格式输出通知信息及换行
func (lg *TDefaultLogger) NoticeFLn(tag string, format string, v ...interface{}) {
	lg.logfln(tag, LEVEL_NOTI, format, v...)
}

// 输出关键信息
func (lg *TDefaultLogger) Critical(tag string, v ...interface{}) {
	lg.log(tag, LEVEL_CRIT, v...)
}

// 格式化输出关键信息
func (lg *TDefaultLogger) CriticalF(tag string, format string, v ...interface{}) {
	lg.logf(tag, LEVEL_CRIT, format, v...)
}

// 格式化输出关键信息及换行
func (lg *TDefaultLogger) CriticalFLn(tag string, format string, v ...interface{}) {
	lg.logfln(tag, LEVEL_CRIT, format, v...)
}
