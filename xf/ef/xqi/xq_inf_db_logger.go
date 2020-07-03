package xqi

import (
	"github.com/go-xe2/x/core/logger"
	"time"
)

type DbLogger interface {
	Sql(sqlStr string, runtime time.Duration)
	Slow(sqlStr string, runtime time.Duration)
	Error(msg string)
	EnableSqlLog() bool
	EnableErrorLog() bool
	EnableSlowLog() float64
	GetLogger() logger.ILogger
	SetLogger(logger logger.ILogger)
	This() interface{}
}
