package xarray

import (
	"bytes"
	"encoding/json"
	"github.com/go-xe2/x/core/rwmutex"
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/utils/xrand"
	"math"
	"sort"
)

type TArray struct {
	mu    *rwmutex.RWMutex
	array []interface{}
}

func New(unsafe ...bool) *TArray {
	return NewArraySize(0, 0, unsafe...)
}

func NewArray(unsafe ...bool) *TArray {
	return NewArraySize(0, 0, unsafe...)
}

func NewArraySize(size int, cap int, unsafe ...bool) *TArray {
	return &TArray{
		mu:    rwmutex.New(unsafe...),
		array: make([]interface{}, size, cap),
	}
}

func NewFrom(array []interface{}, unsafe ...bool) *TArray {
	return NewArrayFrom(array, unsafe...)
}

func NewFromCopy(array []interface{}, unsafe ...bool) *TArray {
	return NewArrayFromCopy(array, unsafe...)
}

func NewArrayFrom(array []interface{}, unsafe ...bool) *TArray {
	return &TArray{
		mu:    rwmutex.New(unsafe...),
		array: array,
	}
}

func NewArrayFromCopy(array []interface{}, unsafe ...bool) *TArray {
	newTArray := make([]interface{}, len(array))
	copy(newTArray, array)
	return &TArray{
		mu:    rwmutex.New(unsafe...),
		array: newTArray,
	}
}

func (a *TArray) Get(index int) interface{} {
	a.mu.RLock()
	defer a.mu.RUnlock()
	value := a.array[index]
	return value
}

func (a *TArray) Set(index int, value interface{}) *TArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.array[index] = value
	return a
}

func (a *TArray) SetTArray(array []interface{}) *TArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.array = array
	return a
}

func (a *TArray) Replace(array []interface{}) *TArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	max := len(array)
	if max > len(a.array) {
		max = len(a.array)
	}
	for i := 0; i < max; i++ {
		a.array[i] = array[i]
	}
	return a
}

func (a *TArray) Sum() (sum int) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, v := range a.array {
		sum += t.Int(v)
	}
	return
}

func (a *TArray) SortFunc(less func(v1, v2 interface{}) bool) *TArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	sort.Slice(a.array, func(i, j int) bool {
		return less(a.array[i], a.array[j])
	})
	return a
}

func (a *TArray) InsertBefore(index int, value interface{}) *TArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	rear := append([]interface{}{}, a.array[index:]...)
	a.array = append(a.array[0:index], value)
	a.array = append(a.array, rear...)
	return a
}

func (a *TArray) InsertAfter(index int, value interface{}) *TArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	rear := append([]interface{}{}, a.array[index+1:]...)
	a.array = append(a.array[0:index+1], value)
	a.array = append(a.array, rear...)
	return a
}

func (a *TArray) Remove(index int) interface{} {
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

func (a *TArray) PushLeft(value ...interface{}) *TArray {
	a.mu.Lock()
	a.array = append(value, a.array...)
	a.mu.Unlock()
	return a
}

func (a *TArray) PushRight(value ...interface{}) *TArray {
	a.mu.Lock()
	a.array = append(a.array, value...)
	a.mu.Unlock()
	return a
}

func (a *TArray) PopRand() interface{} {
	return a.Remove(xrand.Intn(len(a.array)))
}

func (a *TArray) PopRands(size int) []interface{} {
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

func (a *TArray) PopLeft() interface{} {
	a.mu.Lock()
	defer a.mu.Unlock()
	value := a.array[0]
	a.array = a.array[1:]
	return value
}

func (a *TArray) PopRight() interface{} {
	a.mu.Lock()
	defer a.mu.Unlock()
	index := len(a.array) - 1
	value := a.array[index]
	a.array = a.array[:index]
	return value
}

func (a *TArray) PopLefts(size int) []interface{} {
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

func (a *TArray) PopRights(size int) []interface{} {
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

func (a *TArray) Range(start int, end ...int) []interface{} {
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

func (a *TArray) SubSlice(offset int, length ...int) []interface{} {
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

func (a *TArray) Append(value ...interface{}) *TArray {
	a.PushRight(value...)
	return a
}

func (a *TArray) Len() int {
	a.mu.RLock()
	length := len(a.array)
	a.mu.RUnlock()
	return length
}

func (a *TArray) Slice() []interface{} {
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

func (a *TArray) Clone() (newTArray *TArray) {
	a.mu.RLock()
	array := make([]interface{}, len(a.array))
	copy(array, a.array)
	a.mu.RUnlock()
	return NewArrayFrom(array, !a.mu.IsSafe())
}

func (a *TArray) Clear() *TArray {
	a.mu.Lock()
	if len(a.array) > 0 {
		a.array = make([]interface{}, 0)
	}
	a.mu.Unlock()
	return a
}

func (a *TArray) Contains(value interface{}) bool {
	return a.Search(value) != -1
}

func (a *TArray) Search(value interface{}) int {
	if len(a.array) == 0 {
		return -1
	}
	a.mu.RLock()
	result := -1
	for index, v := range a.array {
		if v == value {
			result = index
			break
		}
	}
	a.mu.RUnlock()

	return result
}

func (a *TArray) Unique() *TArray {
	a.mu.Lock()
	for i := 0; i < len(a.array)-1; i++ {
		for j := i + 1; j < len(a.array); j++ {
			if a.array[i] == a.array[j] {
				a.array = append(a.array[:j], a.array[j+1:]...)
			}
		}
	}
	a.mu.Unlock()
	return a
}

func (a *TArray) LockFunc(f func(array []interface{})) *TArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	f(a.array)
	return a
}

func (a *TArray) RLockFunc(f func(array []interface{})) *TArray {
	a.mu.RLock()
	defer a.mu.RUnlock()
	f(a.array)
	return a
}

func (a *TArray) Merge(array interface{}) *TArray {
	switch v := array.(type) {
	case *TArray:
		a.Append(t.Interfaces(v.Slice())...)
	case *TIntArray:
		a.Append(t.Interfaces(v.Slice())...)
	case *TStringArray:
		a.Append(t.Interfaces(v.Slice())...)
	case *TSortedArray:
		a.Append(t.Interfaces(v.Slice())...)
	case *TSortedIntArray:
		a.Append(t.Interfaces(v.Slice())...)
	case *TSortedStringArray:
		a.Append(t.Interfaces(v.Slice())...)
	default:
		a.Append(t.Interfaces(array)...)
	}
	return a
}

func (a *TArray) Fill(startIndex int, num int, value interface{}) *TArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	if startIndex < 0 {
		startIndex = 0
	}
	for i := startIndex; i < startIndex+num; i++ {
		if i > len(a.array)-1 {
			a.array = append(a.array, value)
		} else {
			a.array[i] = value
		}
	}
	return a
}

func (a *TArray) Chunk(size int) [][]interface{} {
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

func (a *TArray) Pad(size int, val interface{}) *TArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	if size == 0 || (size > 0 && size < len(a.array)) || (size < 0 && size > -len(a.array)) {
		return a
	}
	n := size
	if size < 0 {
		n = -size
	}
	n -= len(a.array)
	tmp := make([]interface{}, n)
	for i := 0; i < n; i++ {
		tmp[i] = val
	}
	if size > 0 {
		a.array = append(a.array, tmp...)
	} else {
		a.array = append(tmp, a.array...)
	}
	return a
}

func (a *TArray) Rand() interface{} {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.array[xrand.Intn(len(a.array))]
}

func (a *TArray) Rands(size int) []interface{} {
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

func (a *TArray) Shuffle() *TArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	for i, v := range xrand.Perm(len(a.array)) {
		a.array[i], a.array[v] = a.array[v], a.array[i]
	}
	return a
}

func (a *TArray) Reverse() *TArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	for i, j := 0, len(a.array)-1; i < j; i, j = i+1, j-1 {
		a.array[i], a.array[j] = a.array[j], a.array[i]
	}
	return a
}

func (a *TArray) Join(glue string) string {
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

func (a *TArray) CountValues() map[interface{}]int {
	m := make(map[interface{}]int)
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, v := range a.array {
		m[v]++
	}
	return m
}

func (a *TArray) String() string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	jsonContent, _ := json.Marshal(a.array)
	return string(jsonContent)
}

func (a *TArray) MarshalJSON() ([]byte, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return json.Marshal(a.array)
}
