package xarray

import (
	"bytes"
	"encoding/json"
	"github.com/go-xe2/x/core/rwmutex"
	_type "github.com/go-xe2/x/sync/type"
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/utils/xrand"
	"math"
	"sort"
)

type TSortedArray struct {
	mu         *rwmutex.RWMutex
	array      []interface{}
	unique     *_type.TBool
	comparator func(v1, v2 interface{}) int
}

func NewSortedArray(comparator func(v1, v2 interface{}) int, unsafe ...bool) *TSortedArray {
	return NewSortedArraySize(0, comparator, unsafe...)
}

func NewSortedArraySize(cap int, comparator func(v1, v2 interface{}) int, unsafe ...bool) *TSortedArray {
	return &TSortedArray{
		mu:         rwmutex.New(unsafe...),
		unique:     _type.NewBool(),
		array:      make([]interface{}, 0, cap),
		comparator: comparator,
	}
}

func NewSortedArrayFrom(array []interface{}, comparator func(v1, v2 interface{}) int, unsafe ...bool) *TSortedArray {
	a := NewSortedArraySize(0, comparator, unsafe...)
	a.array = array
	sort.Slice(a.array, func(i, j int) bool {
		return a.comparator(a.array[i], a.array[j]) < 0
	})
	return a
}

func NewSortedArrayFromCopy(array []interface{}, comparator func(v1, v2 interface{}) int, unsafe ...bool) *TSortedArray {
	newArray := make([]interface{}, len(array))
	copy(newArray, array)
	return NewSortedArrayFrom(newArray, comparator, unsafe...)
}

func (a *TSortedArray) SetArray(array []interface{}) *TSortedArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.array = array
	sort.Slice(a.array, func(i, j int) bool {
		return a.comparator(a.array[i], a.array[j]) < 0
	})
	return a
}

func (a *TSortedArray) Sort() *TSortedArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	sort.Slice(a.array, func(i, j int) bool {
		return a.comparator(a.array[i], a.array[j]) < 0
	})
	return a
}

func (a *TSortedArray) Add(values ...interface{}) *TSortedArray {
	if len(values) == 0 {
		return a
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	for _, value := range values {
		index, cmp := a.binSearch(value, false)
		if a.unique.Val() && cmp == 0 {
			continue
		}
		if index < 0 {
			a.array = append(a.array, value)
			continue
		}
		if cmp > 0 {
			index++
		}
		rear := append([]interface{}{}, a.array[index:]...)
		a.array = append(a.array[0:index], value)
		a.array = append(a.array, rear...)
	}
	return a
}

func (a *TSortedArray) Get(index int) interface{} {
	a.mu.RLock()
	defer a.mu.RUnlock()
	value := a.array[index]
	return value
}

func (a *TSortedArray) Remove(index int) interface{} {
	a.mu.Lock()
	defer a.mu.Unlock()
	if index == 0 {
		value := a.array[0]
		a.array = a.array[1:]
		return value
	} else if index == len(a.array)-1 {
		value := a.array[index]
		a.array = a.array[:index]
		return value
	}
	value := a.array[index]
	a.array = append(a.array[:index], a.array[index+1:]...)
	return value
}

func (a *TSortedArray) PopLeft() interface{} {
	a.mu.Lock()
	defer a.mu.Unlock()
	value := a.array[0]
	a.array = a.array[1:]
	return value
}

func (a *TSortedArray) PopRight() interface{} {
	a.mu.Lock()
	defer a.mu.Unlock()
	index := len(a.array) - 1
	value := a.array[index]
	a.array = a.array[:index]
	return value
}

func (a *TSortedArray) PopRand() interface{} {
	return a.Remove(xrand.Intn(len(a.array)))
}

func (a *TSortedArray) PopRands(size int) []interface{} {
	a.mu.Lock()
	defer a.mu.Unlock()
	if size > len(a.array) {
		size = len(a.array)
	}
	array := make([]interface{}, size)
	for i := 0; i < size; i++ {
		index := xrand.Intn(len(a.array))
		array[i] = a.array[index]
		a.array = append(a.array[:index], a.array[index+1:]...)
	}
	return array
}

func (a *TSortedArray) PopLefts(size int) []interface{} {
	a.mu.Lock()
	defer a.mu.Unlock()
	length := len(a.array)
	if size > length {
		size = length
	}
	value := a.array[0:size]
	a.array = a.array[size:]
	return value
}

func (a *TSortedArray) PopRights(size int) []interface{} {
	a.mu.Lock()
	defer a.mu.Unlock()
	index := len(a.array) - size
	if index < 0 {
		index = 0
	}
	value := a.array[index:]
	a.array = a.array[:index]
	return value
}

func (a *TSortedArray) Range(start int, end ...int) []interface{} {
	a.mu.RLock()
	defer a.mu.RUnlock()
	offsetEnd := len(a.array)
	if len(end) > 0 && end[0] < offsetEnd {
		offsetEnd = end[0]
	}
	if start > offsetEnd {
		return nil
	}
	if start < 0 {
		start = 0
	}
	array := ([]interface{})(nil)
	if a.mu.IsSafe() {
		array = make([]interface{}, offsetEnd-start)
		copy(array, a.array[start:offsetEnd])
	} else {
		array = a.array[start:offsetEnd]
	}
	return array
}

func (a *TSortedArray) SubSlice(offset int, length ...int) []interface{} {
	a.mu.RLock()
	defer a.mu.RUnlock()
	size := len(a.array)
	if len(length) > 0 {
		size = length[0]
	}
	if offset > len(a.array) {
		return nil
	}
	if offset < 0 {
		offset = len(a.array) + offset
		if offset < 0 {
			return nil
		}
	}
	if size < 0 {
		offset += size
		size = -size
		if offset < 0 {
			return nil
		}
	}
	end := offset + size
	if end > len(a.array) {
		end = len(a.array)
		size = len(a.array) - offset
	}
	if a.mu.IsSafe() {
		s := make([]interface{}, size)
		copy(s, a.array[offset:])
		return s
	} else {
		return a.array[offset:end]
	}
}

func (a *TSortedArray) Sum() (sum int) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, v := range a.array {
		sum += t.Int(v)
	}
	return
}

func (a *TSortedArray) Len() int {
	a.mu.RLock()
	length := len(a.array)
	a.mu.RUnlock()
	return length
}

func (a *TSortedArray) Slice() []interface{} {
	array := ([]interface{})(nil)
	if a.mu.IsSafe() {
		a.mu.RLock()
		defer a.mu.RUnlock()
		array = make([]interface{}, len(a.array))
		copy(array, a.array)
	} else {
		array = a.array
	}
	return array
}

func (a *TSortedArray) Contains(value interface{}) bool {
	return a.Search(value) != -1
}

func (a *TSortedArray) Search(value interface{}) (index int) {
	if i, r := a.binSearch(value, true); r == 0 {
		return i
	}
	return -1
}

func (a *TSortedArray) binSearch(value interface{}, lock bool) (index int, result int) {
	if len(a.array) == 0 {
		return -1, -2
	}
	if lock {
		a.mu.RLock()
		defer a.mu.RUnlock()
	}
	min := 0
	max := len(a.array) - 1
	mid := 0
	cmp := -2
	for min <= max {
		mid = int((min + max) / 2)
		cmp = a.comparator(value, a.array[mid])
		switch {
		case cmp < 0:
			max = mid - 1
		case cmp > 0:
			min = mid + 1
		default:
			return mid, cmp
		}
	}
	return mid, cmp
}

func (a *TSortedArray) SetUnique(unique bool) *TSortedArray {
	oldUnique := a.unique.Val()
	a.unique.Set(unique)
	if unique && oldUnique != unique {
		a.Unique()
	}
	return a
}

func (a *TSortedArray) Unique() *TSortedArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	i := 0
	for {
		if i == len(a.array)-1 {
			break
		}
		if a.comparator(a.array[i], a.array[i+1]) == 0 {
			a.array = append(a.array[:i+1], a.array[i+1+1:]...)
		} else {
			i++
		}
	}
	return a
}

func (a *TSortedArray) Clone() (newArray *TSortedArray) {
	a.mu.RLock()
	array := make([]interface{}, len(a.array))
	copy(array, a.array)
	a.mu.RUnlock()
	return NewSortedArrayFrom(array, a.comparator, !a.mu.IsSafe())
}

func (a *TSortedArray) Clear() *TSortedArray {
	a.mu.Lock()
	if len(a.array) > 0 {
		a.array = make([]interface{}, 0)
	}
	a.mu.Unlock()
	return a
}

func (a *TSortedArray) LockFunc(f func(array []interface{})) *TSortedArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	f(a.array)
	return a
}

func (a *TSortedArray) RLockFunc(f func(array []interface{})) *TSortedArray {
	a.mu.RLock()
	defer a.mu.RUnlock()
	f(a.array)
	return a
}

func (a *TSortedArray) Merge(array interface{}) *TSortedArray {
	switch v := array.(type) {
	case *TArray:
		a.Add(t.Interfaces(v.Slice())...)
	case *TIntArray:
		a.Add(t.Interfaces(v.Slice())...)
	case *TStringArray:
		a.Add(t.Interfaces(v.Slice())...)
	case *TSortedArray:
		a.Add(t.Interfaces(v.Slice())...)
	case *TSortedIntArray:
		a.Add(t.Interfaces(v.Slice())...)
	case *TSortedStringArray:
		a.Add(t.Interfaces(v.Slice())...)
	default:
		a.Add(t.Interfaces(array)...)
	}
	return a
}

func (a *TSortedArray) Chunk(size int) [][]interface{} {
	if size < 1 {
		return nil
	}
	a.mu.RLock()
	defer a.mu.RUnlock()
	length := len(a.array)
	chunks := int(math.Ceil(float64(length) / float64(size)))
	var n [][]interface{}
	for i, end := 0, 0; chunks > 0; chunks-- {
		end = (i + 1) * size
		if end > length {
			end = length
		}
		n = append(n, a.array[i*size:end])
		i++
	}
	return n
}

func (a *TSortedArray) Rand() interface{} {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.array[xrand.Intn(len(a.array))]
}

func (a *TSortedArray) Rands(size int) []interface{} {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if size > len(a.array) {
		size = len(a.array)
	}
	n := make([]interface{}, size)
	for i, v := range xrand.Perm(len(a.array)) {
		n[i] = a.array[v]
		if i == size-1 {
			break
		}
	}
	return n
}

func (a *TSortedArray) Join(glue string) string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	buffer := bytes.NewBuffer(nil)
	for k, v := range a.array {
		buffer.WriteString(t.String(v))
		if k != len(a.array)-1 {
			buffer.WriteString(glue)
		}
	}
	return buffer.String()
}

func (a *TSortedArray) CountValues() map[interface{}]int {
	m := make(map[interface{}]int)
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, v := range a.array {
		m[v]++
	}
	return m
}

func (a *TSortedArray) String() string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	jsonContent, _ := json.Marshal(a.array)
	return string(jsonContent)
}

func (a *TSortedArray) MarshalJSON() ([]byte, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return json.Marshal(a.array)
}
