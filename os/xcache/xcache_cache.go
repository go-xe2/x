package xcache

import (
	"github.com/go-xe2/x/os/xtimer"
	"sync/atomic"
	"time"
	"unsafe"
)

type TCache struct {
	*tMemCache
}

func New(lruCap ...int) *TCache {
	c := &TCache{
		tMemCache: newMemCache(lruCap...),
	}
	xtimer.AddSingleton(time.Second, c.syncEventAndClearExpired)
	return c
}

func (c *TCache) Clear() {
	old := atomic.SwapPointer((*unsafe.Pointer)(unsafe.Pointer(&c.tMemCache)), unsafe.Pointer(newMemCache()))
	(*tMemCache)(old).Close()
}
