package xconv

import (
	"github.com/go-xe2/x/type/xstring"
	"github.com/go-xe2/x/type/xtime"
	"time"
)

// 转换成time.Time
func Time(i interface{}, format ...string) time.Time {
	if t := XTime(i, format...); t != nil {
		return t.Time
	}
	return time.Time{}
}

//字符串或时间戳转换成time.Duration
func Duration(i interface{}) time.Duration {
	s := String(i)
	if !xstring.IsNumeric(s) {
		d, _ := time.ParseDuration(s)
		return d
	}
	return time.Duration(Int64(i))
}

// 转换成xtime
func XTime(i interface{}, format ...string) *xtime.Time {
	s := String(i)
	if len(s) == 0 {
		return xtime.New()
	}
	if len(format) > 0 {
		t, _ := xtime.StrToTimeFormat(s, format[0])
		return t
	}
	if xstring.IsNumeric(s) {
		return xtime.NewFromTimeStamp(Int64(s))
	} else {
		t, _ := xtime.StrToTime(s)
		return t
	}
}
