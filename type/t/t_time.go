package t

import (
	"github.com/go-xe2/x/type/xconv"
	"github.com/go-xe2/x/type/xtime"
	"time"
)

// 转换成time.Time
func Time(i interface{}, format ...string) time.Time {
	return xconv.Time(i, format...)
}

//字符串或时间戳转换成time.Duration
func Duration(i interface{}) time.Duration {
	return xconv.Duration(i)
}

// 转换成xtime
func XTime(i interface{}, format ...string) *xtime.Time {
	return xconv.XTime(i, format...)
}
