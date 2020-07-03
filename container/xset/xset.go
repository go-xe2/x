package xset

import (
	"github.com/go-xe2/x/core/rwmutex"
	"github.com/go-xe2/x/type/t"
	"strings"
)

type TSet struct {
	mu *rwmutex.RWMutex
	m  map[interface{}]struct{}
}

func New(unsafe ...bool) *TSet {
	return NewSet(unsafe...)
}

func NewSet(unsafe ...bool) *TSet {
	return &TSet{
		m:  make(map[interface{}]struct{}),
		mu: rwmutex.New(unsafe...),
	}
}

func NewFrom(items interface{}, unsafe ...bool) *TSet {
	m := make(map[interface{}]struct{})
	for _, v := range t.Interfaces(items) {
		m[v] = struct{}{}
	}
	return &TSet{
		m:  m,
		mu: rwmutex.New(unsafe...),
	}
}

func (set *TSet) Iterator(f func(v interface{}) bool) *TSet {
	set.mu.RLock()
	defer set.mu.RUnlock()
	for k, _ := range set.m {
		if !f(k) {
			break
		}
	}
	return set
}

func (set *TSet) Add(item ...interface{}) *TSet {
	set.mu.Lock()
	for _, v := range item {
		set.m[v] = struct{}{}
	}
	set.mu.Unlock()
	return set
}

func (set *TSet) Contains(item interface{}) bool {
	set.mu.RLock()
	_, exists := set.m[item]
	set.mu.RUnlock()
	return exists
}

func (set *TSet) Remove(item interface{}) *TSet {
	set.mu.Lock()
	delete(set.m, item)
	set.mu.Unlock()
	return set
}

func (set *TSet) Size() int {
	set.mu.RLock()
	l := len(set.m)
	set.mu.RUnlock()
	return l
}

func (set *TSet) Clear() *TSet {
	set.mu.Lock()
	set.m = make(map[interface{}]struct{})
	set.mu.Unlock()
	return set
}

func (set *TSet) Slice() []interface{} {
	set.mu.RLock()
	i := 0
	ret := make([]interface{}, len(set.m))
	for item := range set.m {
		ret[i] = item
		i++
	}
	set.mu.RUnlock()
	return ret
}

func (set *TSet) Join(glue string) string {
	return strings.Join(t.Strings(set.Slice()), ",")
}

func (set *TSet) String() string {
	return set.Join(",")
}

func (set *TSet) LockFunc(f func(m map[interface{}]struct{})) {
	set.mu.Lock()
	defer set.mu.Unlock()
	f(set.m)
}

func (set *TSet) RLockFunc(f func(m map[interface{}]struct{})) {
	set.mu.RLock()
	defer set.mu.RUnlock()
	f(set.m)
}

func (set *TSet) Equal(other *TSet) bool {
	if set == other {
		return true
	}
	set.mu.RLock()
	defer set.mu.RUnlock()
	other.mu.RLock()
	defer other.mu.RUnlock()
	if len(set.m) != len(other.m) {
		return false
	}
	for key := range set.m {
		if _, ok := other.m[key]; !ok {
			return false
		}
	}
	return true
}

func (set *TSet) IsSubsetOf(other *TSet) bool {
	if set == other {
		return true
	}
	set.mu.RLock()
	defer set.mu.RUnlock()
	other.mu.RLock()
	defer other.mu.RUnlock()
	for key := range set.m {
		if _, ok := other.m[key]; !ok {
			return false
		}
	}
	return true
}

func (set *TSet) Union(others ...*TSet) (newTSet *TSet) {
	newTSet = NewSet(true)
	set.mu.RLock()
	defer set.mu.RUnlock()
	for _, other := range others {
		if set != other {
			other.mu.RLock()
		}
		for k, v := range set.m {
			newTSet.m[k] = v
		}
		if set != other {
			for k, v := range other.m {
				newTSet.m[k] = v
			}
		}
		if set != other {
			other.mu.RUnlock()
		}
	}

	return
}

func (set *TSet) Diff(others ...*TSet) (newTSet *TSet) {
	newTSet = NewSet(true)
	set.mu.RLock()
	defer set.mu.RUnlock()
	for _, other := range others {
		if set == other {
			continue
		}
		other.mu.RLock()
		for k, v := range set.m {
			if _, ok := other.m[k]; !ok {
				newTSet.m[k] = v
			}
		}
		other.mu.RUnlock()
	}
	return
}

func (set *TSet) Intersect(others ...*TSet) (newTSet *TSet) {
	newTSet = NewSet(true)
	set.mu.RLock()
	defer set.mu.RUnlock()
	for _, other := range others {
		if set != other {
			other.mu.RLock()
		}
		for k, v := range set.m {
			if _, ok := other.m[k]; ok {
				newTSet.m[k] = v
			}
		}
		if set != other {
			other.mu.RUnlock()
		}
	}
	return
}

func (set *TSet) Complement(full *TSet) (newTSet *TSet) {
	newTSet = NewSet(true)
	set.mu.RLock()
	defer set.mu.RUnlock()
	if set != full {
		full.mu.RLock()
		defer full.mu.RUnlock()
	}
	for k, v := range full.m {
		if _, ok := set.m[k]; !ok {
			newTSet.m[k] = v
		}
	}
	return
}

func (set *TSet) Merge(others ...*TSet) *TSet {
	set.mu.Lock()
	defer set.mu.Unlock()
	for _, other := range others {
		if set != other {
			other.mu.RLock()
		}
		for k, v := range other.m {
			set.m[k] = v
		}
		if set != other {
			other.mu.RUnlock()
		}
	}
	return set
}

func (set *TSet) Sum() (sum int) {
	set.mu.RLock()
	defer set.mu.RUnlock()
	for k, _ := range set.m {
		sum += t.Int(k)
	}
	return
}

func (set *TSet) Pop(size int) interface{} {
	set.mu.RLock()
	defer set.mu.RUnlock()
	for k, _ := range set.m {
		return k
	}
	return nil
}

func (set *TSet) Pops(size int) []interface{} {
	set.mu.RLock()
	defer set.mu.RUnlock()
	if size > len(set.m) {
		size = len(set.m)
	}
	index := 0
	array := make([]interface{}, size)
	for k, _ := range set.m {
		array[index] = k
		index++
		if index == size {
			break
		}
	}
	return array
}
