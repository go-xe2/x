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
	"strings"
)

type TSortedStringArray struct {
	mu         *rwmutex.RWMutex
	array      []string
	unique     *_type.TBool
	comparator func(v1, v2 string) int
}

func NewSortedStringArray(unsafe ...bool) *TSortedStringArray {
	return NewSortedStringArraySize(0, unsafe...)
}

func NewSortedStringArraySize(cap int, unsafe ...bool) *TSortedStringArray {
	return &TSortedStringArray{
		mu:     rwmutex.New(unsafe...),
		array:  make([]string, 0, cap),
		unique: _type.NewBool(),
		comparator: func(v1, v2 string) int {
			return strings.Compare(v1, v2)
		},
	}
}

func NewTSortedStringArrayFrom(array []string, unsafe ...bool) *TSortedStringArray {
	a := NewSortedStringArraySize(0, unsafe...)
	a.array = array
	sort.Strings(a.array)
	return a
}

func NewSortedStringArrayFromCopy(array []string, unsafe ...bool) *TSortedStringArray {
	newArray := make([]string, len(array))
	copy(newArray, array)
	return NewTSortedStringArrayFrom(newArray, unsafe...)
}

func (a *TSortedStringArray) SetArray(array []string) *TSortedStringArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.array = array
	sort.Strings(a.array)
	return a
}

func (a *TSortedStringArray) Sort() *TSortedStringArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	sort.Strings(a.array)
	return a
}

func (a *TSortedStringArray) Add(values ...string) *TSortedStringArray {
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
		rear := append([]string{}, a.array[index:]...)
		a.array = append(a.array[0:index], value)
		a.array = append(a.array, rear...)
	}
	return a
}

func (a *TSortedStringArray) Get(index int) string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	value := a.array[index]
	return value
}

func (a *TSortedStringArray) Remove(index int) string {
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

func (a *TSortedStringArray) PopLeft() string {
	a.mu.Lock()
	defer a.mu.Unlock()
	value := a.array[0]
	a.array = a.array[1:]
	return value
}

func (a *TSortedStringArray) PopRight() string {
	a.mu.Lock()
	defer a.mu.Unlock()
	index := len(a.array) - 1
	value := a.array[index]
	a.array = a.array[:index]
	return value
}

func (a *TSortedStringArray) PopRand() string {
	return a.Remove(xrand.Intn(len(a.array)))
}

func (a *TSortedStringArray) PopRands(size int) []string {
	a.mu.Lock()
	defer a.mu.Unlock()
	if size > len(a.array) {
		size = len(a.array)
	}
	array := make([]string, size)
	for i := 0; i < size; i++ {
		index := xrand.Intn(len(a.array))
		array[i] = a.array[index]
		a.array = append(a.array[:index], a.array[index+1:]...)
	}
	return array
}

func (a *TSortedStringArray) PopLefts(size int) []string {
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

func (a *TSortedStringArray) PopRights(size int) []string {
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

func (a *TSortedStringArray) Range(start int, end ...int) []string {
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
	array := ([]string)(nil)
	if a.mu.IsSafe() {
		array = make([]string, offsetEnd-start)
		copy(array, a.array[start:offsetEnd])
	} else {
		array = a.array[start:offsetEnd]
	}
	return array
}

func (a *TSortedStringArray) SubSlice(offset int, length ...int) []string {
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
		s := make([]string, size)
		copy(s, a.array[offset:])
		return s
	} else {
		return a.array[offset:end]
	}
}

func (a *TSortedStringArray) Sum() (sum int) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, v := range a.array {
		sum += t.Int(v)
	}
	return
}

func (a *TSortedStringArray) Len() int {
	a.mu.RLock()
	length := len(a.array)
	a.mu.RUnlock()
	return length
}

func (a *TSortedStringArray) Slice() []string {
	array := ([]string)(nil)
	if a.mu.IsSafe() {
		a.mu.RLock()
		defer a.mu.RUnlock()
		array = make([]string, len(a.array))
		copy(array, a.array)
	} else {
		array = a.array
	}
	return array
}

func (a *TSortedStringArray) Contains(value string) bool {
	return a.Search(value) != -1
}

func (a *TSortedStringArray) Search(value string) (index int) {
	if i, r := a.binSearch(value, true); r == 0 {
		return i
	}
	return -1
}

func (a *TSortedStringArray) binSearch(value string, lock bool) (index int, result int) {
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

func (a *TSortedStringArray) SetUnique(unique bool) *TSortedStringArray {
	oldUnique := a.unique.Val()
	a.unique.Set(unique)
	if unique && oldUnique != unique {
		a.Unique()
	}
	return a
}

func (a *TSortedStringArray) Unique() *TSortedStringArray {
	a.mu.Lock()
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
	a.mu.Unlock()
	return a
}

func (a *TSortedStringArray) Clone() (newArray *TSortedStringArray) {
	a.mu.RLock()
	array := make([]string, len(a.array))
	copy(array, a.array)
	a.mu.RUnlock()
	return NewTSortedStringArrayFrom(array, !a.mu.IsSafe())
}

func (a *TSortedStringArray) Clear() *TSortedStringArray {
	a.mu.Lock()
	if len(a.array) > 0 {
		a.array = make([]string, 0)
	}
	a.mu.Unlock()
	return a
}

func (a *TSortedStringArray) LockFunc(f func(array []string)) *TSortedStringArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	f(a.array)
	return a
}

func (a *TSortedStringArray) RLockFunc(f func(array []string)) *TSortedStringArray {
	a.mu.RLock()
	defer a.mu.RUnlock()
	f(a.array)
	return a
}

func (a *TSortedStringArray) Merge(array interface{}) *TSortedStringArray {
	switch v := array.(type) {
	case *TArray:
		a.Add(t.Strings(v.Slice())...)
	case *TIntArray:
		a.Add(t.Strings(v.Slice())...)
	case *TStringArray:
		a.Add(t.Strings(v.Slice())...)
	case *TSortedArray:
		a.Add(t.Strings(v.Slice())...)
	case *TSortedIntArray:
		a.Add(t.Strings(v.Slice())...)
	case *TSortedStringArray:
		a.Add(t.Strings(v.Slice())...)
	default:
		a.Add(t.Strings(array)...)
	}
	return a
}

func (a *TSortedStringArray) Chunk(size int) [][]string {
	if size < 1 {
		return nil
	}
	a.mu.RLock()
	defer a.mu.RUnlock()
	length := len(a.array)
	chunks := int(math.Ceil(float64(length) / float64(size)))
	var n [][]string
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

func (a *TSortedStringArray) Rand() string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.array[xrand.Intn(len(a.array))]
}

func (a *TSortedStringArray) Rands(size int) []string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if size > len(a.array) {
		size = len(a.array)
	}
	n := make([]string, size)
	for i, v := range xrand.Perm(len(a.array)) {
		n[i] = a.array[v]
		if i == size-1 {
			break
		}
	}
	return n
}

func (a *TSortedStringArray) Join(glue string) string {
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

func (a *TSortedStringArray) CountValues() map[string]int {
	m := make(map[string]int)
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, v := range a.array {
		m[v]++
	}
	return m
}

func (a *TSortedStringArray) String() string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	jsonContent, _ := json.Marshal(a.array)
	return string(jsonContent)
}

func (a *TSortedStringArray) MarshalJSON() ([]byte, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return json.Marshal(a.array)
}
