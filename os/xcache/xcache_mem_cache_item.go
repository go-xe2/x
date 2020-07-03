package xcache

import "github.com/go-xe2/x/type/xtime"

type tMemCacheItem struct {
	v interface{} // Value.
	e int64       // Expire time in milliseconds.
}

func (item *tMemCacheItem) IsExpired() bool {
	if item.e >= xtime.Millisecond() {
		return false
	}
	return true
}
