package xdatabase

import (
	"fmt"
	xlogger "github.com/go-xe2/x/core/logger"
	. "github.com/go-xe2/x/xf/ef/xqi"
	"github.com/go-xe2/x/xf/log/xdefaultLogger"
	"sync"
	"time"
)

type LogLevel uint

const (
	LOG_SQL LogLevel = iota
	LOG_SLOW
	LOG_ERROR
)

func (l LogLevel) String() string {
	switch l {
	case LOG_SQL:
		return "SQL"
	case LOG_SLOW:
		return "SLOW"
	case LOG_ERROR:
		return "ERROR"
	}
	return ""
}

type LogOption struct {
	EnableSqlLog bool
	// 是否记录慢查询, 默认0s, 不记录, 设置记录的时间阀值, 比如 1, 则表示超过1s的都记录
	EnableSlowLog  float64
	EnableErrorLog bool
}

type TDbLogger struct {
	logger  xlogger.ILogger
	sqlLog  bool
	slowLog float64
	//infoLog  bool
	errLog bool
}

func NewDbLogOptions() *LogOption {
	return &LogOption{
		EnableSqlLog:   true,
		EnableSlowLog:  1,
		EnableErrorLog: true,
	}
}

var _ DbLogger = (*TDbLogger)(nil)

var onceLogger sync.Once
var logger *TDbLogger

func NewDbLogger(o *LogOption, log xlogger.ILogger) *TDbLogger {
	onceLogger.Do(func() {
		logger = &TDbLogger{logger: log}
		logger.sqlLog = o.EnableSqlLog
		logger.slowLog = o.EnableSlowLog
		logger.errLog = o.EnableErrorLog
	})
	return logger
}

// 日志插件
func NewDbLoggerPlugin(o *LogOption, log xlogger.ILogger) func(e DbConn) {
	return func(e DbConn) {
		e.SetLogger(NewDbLogger(o, log))
	}
}

func DefaultLogger() func(e DbConn) {
	return func(e DbConn) {
		dbLogOption := NewDbLogOptions()
		dbLogOption.EnableSlowLog = 3
		e.SetLogger(NewDbLogger(dbLogOption, xdefaultLogger.New()))
	}
}

func (l *TDbLogger) EnableSqlLog() bool {
	return l.sqlLog
}

func (l *TDbLogger) EnableErrorLog() bool {
	return l.errLog
}

func (l *TDbLogger) EnableSlowLog() float64 {
	return l.slowLog
}

func (l *TDbLogger) Slow(sqlStr string, runtime time.Duration) {
	if runtime.Seconds() > l.EnableSlowLog() {
		logger.write(LOG_SLOW, sqlStr, runtime.String())
	}
}

func (l *TDbLogger) Sql(sqlStr string, runtime time.Duration) {
	if l.EnableSqlLog() {
		logger.write(LOG_SQL, sqlStr, runtime.String())
	}
}

func (l *TDbLogger) Error(msg string) {
	if l.EnableErrorLog() {
		logger.write(LOG_ERROR, msg, "0")
	}
}

func (l *TDbLogger) This() interface{} {
	return l
}

func (l *TDbLogger) SetLogger(logger xlogger.ILogger) {
	l.logger = logger
}

func (l *TDbLogger) GetLogger() xlogger.ILogger {
	return l.logger
}

func (l *TDbLogger) write(ll LogLevel, msg string, runtime string) {
	now := time.Now()
	datetime := now.Format("2006-01-02 15:04:05")
	content := fmt.Sprintf("[%v] [%v] %v --- %v\n", ll.String(), datetime, runtime, msg)
	TagName := "DBLogger"
	if l.logger != nil {
		switch ll {
		case LOG_SQL:
			l.logger.Debug(TagName, content)
			break
		case LOG_SLOW:
			l.logger.Warning(TagName, content)
			break
		case LOG_ERROR:
			l.logger.Error(TagName, content)
			break
		default:
			l.logger.Debug(TagName, content)
		}
	}
}
