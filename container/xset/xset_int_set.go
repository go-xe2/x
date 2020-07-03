package xset

import (
	"github.com/go-xe2/x/core/rwmutex"
	"github.com/go-xe2/x/type/t"
	"strings"
)

type TIntSet struct {
	mu *rwmutex.RWMutex
	m  map[int]struct{}
}

func NewIntSet(unsafe ...bool) *TIntSet {
	return &TIntSet{
		m:  make(map[int]struct{}),
		mu: rwmutex.New(unsafe...),
	}
}

func NewIntSetFrom(items []int, unsafe ...bool) *TIntSet {
	m := make(map[int]struct{})
	for _, v := range items {
		m[v] = struct{}{}
	}
	return &TIntSet{
		m:  m,
		mu: rwmutex.New(unsafe...),
	}
}

func (set *TIntSet) Iterator(f func(v int) bool) *TIntSet {
	set.mu.RLock()
	defer set.mu.RUnlock()
	for k, _ := range set.m {
		if !f(k) {
			break
		}
	}
	return set
}

func (set *TIntSet) Add(item ...int) *TIntSet {
	set.mu.Lock()
	for _, v := range item {
		set.m[v] = struct{}{}
	}
	set.mu.Unlock()
	return set
}

func (set *TIntSet) Contains(item int) bool {
	set.mu.RLock()
	_, exists := set.m[item]
	set.mu.RUnlock()
	return exists
}

func (set *TIntSet) Remove(item int) *TIntSet {
	set.mu.Lock()
	delete(set.m, item)
	set.mu.Unlock()
	return set
}

func (set *TIntSet) Size() int {
	set.mu.RLock()
	l := len(set.m)
	set.mu.RUnlock()
	return l
}

func (set *TIntSet) Clear() *TIntSet {
	set.mu.Lock()
	set.m = make(map[int]struct{})
	set.mu.Unlock()
	return set
}

func (set *TIntSet) Slice() []int {
	set.mu.RLock()
	ret := make([]int, len(set.m))
	i := 0
	for k, _ := range set.m {
		ret[i] = k
		i++
	}
	set.mu.RUnlock()
	return ret
}

func (set *TIntSet) Join(glue string) string {
	return strings.Join(t.Strings(set.Slice()), ",")
}

func (set *TIntSet) String() string {
	return set.Join(",")
}

func (set *TIntSet) LockFunc(f func(m map[int]struct{})) {
	set.mu.Lock()
	defer set.mu.Unlock()
	f(set.m)
}

func (set *TIntSet) RLockFunc(f func(m map[int]struct{})) {
	set.mu.RLock()
	defer set.mu.RUnlock()
	f(set.m)
}

func (set *TIntSet) Equal(other *TIntSet) bool {
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

func (set *TIntSet) IsSubsetOf(other *TIntSet) bool {
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

func (set *TIntSet) Union(others ...*TIntSet) (newSet *TIntSet) {
	newSet = NewIntSet(true)
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

func (set *TIntSet) Diff(others ...*TIntSet) (newSet *TIntSet) {
	newSet = NewIntSet(true)
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

func (set *TIntSet) Intersect(others ...*TIntSet) (newSet *TIntSet) {
	newSet = NewIntSet(true)
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

func (set *TIntSet) Complement(full *TIntSet) (newSet *TIntSet) {
	newSet = NewIntSet(true)
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

func (set *TIntSet) Merge(others ...*TIntSet) *TIntSet {
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

func (set *TIntSet) Sum() (sum int) {
	set.mu.RLock()
	defer set.mu.RUnlock()
	for k, _ := range set.m {
		sum += k
	}
	return
}

func (set *TIntSet) Pop(size int) int {
	set.mu.RLock()
	defer set.mu.RUnlock()
	for k, _ := range set.m {
		return k
	}
	return 0
}

func (set *TIntSet) Pops(size int) []int {
	set.mu.RLock()
	defer set.mu.RUnlock()
	if size > len(set.m) {
		size = len(set.m)
	}
	index := 0
	array := make([]int, size)
	for k, _ := range set.m {
		array[index] = k
		index++
		if index == size {
			break
		}
	}
	return array
}
