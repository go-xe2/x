package xfileCache

import (
	"github.com/go-xe2/x/core/cmdenv"
	"github.com/go-xe2/x/os/xcache"
	"github.com/go-xe2/x/os/xfile"
	"github.com/go-xe2/x/os/xfileNotify"
	"time"
)

const (
	mDEFAULT_CACHE_EXPIRE = 60
)

var (
	cacheExpire = cmdenv.Get("x.xfileCache.expire", mDEFAULT_CACHE_EXPIRE).Int() * 1000
)

func GetContents(path string, duration ...interface{}) string {
	return string(GetBinContents(path, duration...))
}

func GetBinContents(path string, duration ...interface{}) []byte {
	k := cacheKey(path)
	e := cacheExpire
	if len(duration) > 0 {
		e = getSecondExpire(duration[0])
	}
	r := xcache.GetOrSetFuncLock(k, func() interface{} {
		b := xfile.GetBinContents(path)
		if b != nil {
			_, _ = xfileNotify.Add(path, func(event *xfileNotify.TEvent) {
				xcache.Remove(k)
				xfileNotify.Exit()
			})
		}
		return b
	}, e*1000)
	if r != nil {
		return r.([]byte)
	}
	return nil
}

func getSecondExpire(duration interface{}) int {
	if d, ok := duration.(time.Duration); ok {
		return int(d.Nanoseconds() / 1000000000)
	} else {
		return duration.(int)
	}
}

func cacheKey(path string) string {
	return "x.xfileCache:" + path
}
