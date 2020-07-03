package xset

import (
	"github.com/go-xe2/x/core/rwmutex"
	"github.com/go-xe2/x/type/t"
	"strings"
)

type TStrSet struct {
	mu *rwmutex.RWMutex
	m  map[string]struct{}
}

func NewStrSet(unsafe ...bool) *TStrSet {
	return &TStrSet{
		m:  make(map[string]struct{}),
		mu: rwmutex.New(unsafe...),
	}
}

func NewStrSetFrom(items []string, unsafe ...bool) *TStrSet {
	m := make(map[string]struct{})
	for _, v := range items {
		m[v] = struct{}{}
	}
	return &TStrSet{
		m:  m,
		mu: rwmutex.New(unsafe...),
	}
}

func (set *TStrSet) Iterator(f func(v string) bool) *TStrSet {
	set.mu.RLock()
	defer set.mu.RUnlock()
	for k, _ := range set.m {
		if !f(k) {
			break
		}
	}
	return set
}

func (set *TStrSet) Add(item ...string) *TStrSet {
	set.mu.Lock()
	for _, v := range item {
		set.m[v] = struct{}{}
	}
	set.mu.Unlock()
	return set
}

func (set *TStrSet) Contains(item string) bool {
	set.mu.RLock()
	_, exists := set.m[item]
	set.mu.RUnlock()
	return exists
}

func (set *TStrSet) Remove(item string) *TStrSet {
	set.mu.Lock()
	delete(set.m, item)
	set.mu.Unlock()
	return set
}

func (set *TStrSet) Size() int {
	set.mu.RLock()
	l := len(set.m)
	set.mu.RUnlock()
	return l
}

func (set *TStrSet) Clear() *TStrSet {
	set.mu.Lock()
	set.m = make(map[string]struct{})
	set.mu.Unlock()
	return set
}

func (set *TStrSet) Slice() []string {
	set.mu.RLock()
	ret := make([]string, len(set.m))
	i := 0
	for item := range set.m {
		ret[i] = item
		i++
	}

	set.mu.RUnlock()
	return ret
}

func (set *TStrSet) Join(glue string) string {
	return strings.Join(set.Slice(), ",")
}

func (set *TStrSet) String() string {
	return set.Join(",")
}

func (set *TStrSet) LockFunc(f func(m map[string]struct{})) {
	set.mu.Lock()
	defer set.mu.Unlock()
	f(set.m)
}

func (set *TStrSet) RLockFunc(f func(m map[string]struct{})) {
	set.mu.RLock()
	defer set.mu.RUnlock()
	f(set.m)
}

func (set *TStrSet) Equal(other *TStrSet) bool {
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

func (set *TStrSet) IsSubsetOf(other *TStrSet) bool {
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

func (set *TStrSet) Union(others ...*TStrSet) (newSet *TStrSet) {
	newSet = NewStrSet(true)
	set.mu.RLock()
	defer set.mu.RUnlock()
	for _, other := range others {
		if set != other {
			other.mu.RLock()
		}
		for k, v := range set.m {
			newSet.m[k] = v
		}
		if set != other {
			for k, v := range other.m {
				newSet.m[k] = v
			}
		}
		if set != other {
			other.mu.RUnlock()
		}
	}

	return
}

func (set *TStrSet) Diff(others ...*TStrSet) (newSet *TStrSet) {
	newSet = NewStrSet(true)
	set.mu.RLock()
	defer set.mu.RUnlock()
	for _, other := range others {
		if set == other {
			continue
		}
		other.mu.RLock()
		for k, v := range set.m {
			if _, ok := other.m[k]; !ok {
				newSet.m[k] = v
			}
		}
		other.mu.RUnlock()
	}
	return
}

func (set *TStrSet) Intersect(others ...*TStrSet) (newSet *TStrSet) {
	newSet = NewStrSet(true)
	set.mu.RLock()
	defer set.mu.RUnlock()
	for _, other := range others {
		if set != other {
			other.mu.RLock()
		}
		for k, v := range set.m {
			if _, ok := other.m[k]; ok {
				newSet.m[k] = v
			}
		}
		if set != other {
			other.mu.RUnlock()
		}
	}
	return
}

func (set *TStrSet) Complement(full *TStrSet) (newSet *TStrSet) {
	newSet = NewStrSet(true)
	set.mu.RLock()
	defer set.mu.RUnlock()
	if set != full {
		full.mu.RLock()
		defer full.mu.RUnlock()
	}
	for k, v := range full.m {
		if _, ok := set.m[k]; !ok {
			newSet.m[k] = v
		}
	}
	return
}

func (set *TStrSet) Merge(others ...*TStrSet) *TStrSet {
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

func (set *TStrSet) Sum() (sum int) {
	set.mu.RLock()
	defer set.mu.RUnlock()
	for k, _ := range set.m {
		sum += t.Int(k)
	}
	return
}

func (set *TStrSet) Pop(size int) string {
	set.mu.RLock()
	defer set.mu.RUnlock()
	for k, _ := range set.m {
		return k
	}
	return ""
}

func (set *TStrSet) Pops(size int) []string {
	set.mu.RLock()
	defer set.mu.RUnlock()
	if size > len(set.m) {
		size = len(set.m)
	}
	index := 0
	array := make([]string, size)
	for k, _ := range set.m {
		array[index] = k
		index++
		if index == size {
			break
		}
	}
	return array
}
