package xcache

import (
	"github.com/go-xe2/x/container/xstackQe"
	"github.com/go-xe2/x/os/xtimer"
	_type "github.com/go-xe2/x/sync/type"
	"github.com/go-xe2/x/sync/xsafeMap"
	"github.com/go-xe2/x/sync/xsafeStack"
	"time"
)

// LRU cache object.
// It uses list.List from stdlib for its underlying doubly linked list.
type tMemCacheLru struct {
	cache   *tMemCache               // Parent cache object.
	data    *xsafeMap.TAnyAnyMap     // Key mapping to the item of the list.
	list    *xsafeStack.TSafeStackQe // Key list.
	rawList *xsafeStack.TSafeStackQe // History for key adding.
	closed  *_type.TBool             // Closed or not.
}

// newMemCacheLru creates and returns a new LRU object.
func newMemCacheLru(cache *tMemCache) *tMemCacheLru {
	lru := &tMemCacheLru{
		cache:   cache,
		data:    xsafeMap.NewAnyAnyMap(),
		list:    xsafeStack.New(),
		rawList: xsafeStack.New(),
		closed:  _type.NewBool(),
	}
	xtimer.AddSingleton(time.Second, lru.SyncAndClear)
	return lru
}

// Close closes the LRU object.
func (lru *tMemCacheLru) Close() {
	lru.closed.Set(true)
}

// Remove deletes the <key> FROM <lru>.
func (lru *tMemCacheLru) Remove(key interface{}) {
	if v := lru.data.Get(key); v != nil {
		lru.data.Remove(key)
		lru.list.Remove(v.(*xstackQe.TStackElement))
	}
}

// Size returns the size of <lru>.
func (lru *tMemCacheLru) Size() int {
	return lru.data.Size()
}

// Push pushes <key> to the tail of <lru>.
func (lru *tMemCacheLru) Push(key interface{}) {
	lru.rawList.PushBack(key)
}

// Pop deletes and returns the key from tail of <lru>.
func (lru *tMemCacheLru) Pop() interface{} {
	if v := lru.list.PopBack(); v != nil {
		lru.data.Remove(v)
		return v
	}
	return nil
}

// SyncAndClear synchronizes the keys from <rawList> to <list> and <data>
// using Least Recently Used algorithm.
func (lru *tMemCacheLru) SyncAndClear() {
	if lru.closed.Val() {
		xtimer.Exit()
		return
	}
	// Data synchronization.
	for {
		if v := lru.rawList.PopFront(); v != nil {
			// Deleting the key from list.
			if v := lru.data.Get(v); v != nil {
				lru.list.Remove(v.(*xstackQe.TStackElement))
			}
			// Pushing key to the head of the list
			// and setting its list item to hash table for quick indexing.
			lru.data.Set(v, lru.list.PushFront(v))
		} else {
			break
		}
	}
	// Data cleaning up.
	for i := lru.Size() - lru.cache.cap; i > 0; i-- {
		if s := lru.Pop(); s != nil {
			lru.cache.clearByKey(s, true)
		}
	}
}
