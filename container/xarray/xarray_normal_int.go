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

type TIntArray struct {
	mu    *rwmutex.RWMutex
	array []int
}

func NewTIntArray(unsafe ...bool) *TIntArray {
	return NewIntArraySize(0, 0, unsafe...)
}

func NewIntArraySize(size int, cap int, unsafe ...bool) *TIntArray {
	return &TIntArray{
		mu:    rwmutex.New(unsafe...),
		array: make([]int, size, cap),
	}
}

func NewIntArrayFrom(array []int, unsafe ...bool) *TIntArray {
	return &TIntArray{
		mu:    rwmutex.New(unsafe...),
		array: array,
	}
}

func NewIntArrayFromCopy(array []int, unsafe ...bool) *TIntArray {
	newArray := make([]int, len(array))
	copy(newArray, array)
	return &TIntArray{
		mu:    rwmutex.New(unsafe...),
		array: newArray,
	}
}

func (a *TIntArray) Get(index int) int {
	a.mu.RLock()
	defer a.mu.RUnlock()
	value := a.array[index]
	return value
}

func (a *TIntArray) Set(index int, value int) *TIntArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.array[index] = value
	return a
}

func (a *TIntArray) SetArray(array []int) *TIntArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.array = array
	return a
}

func (a *TIntArray) Replace(array []int) *TIntArray {
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

func (a *TIntArray) Sum() (sum int) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, v := range a.array {
		sum += v
	}
	return
}

func (a *TIntArray) Sort(reverse ...bool) *TIntArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	if len(reverse) > 0 && reverse[0] {
		sort.Slice(a.array, func(i, j int) bool {
			if a.array[i] < a.array[j] {
				return false
			}
			return true
		})
	} else {
		sort.Ints(a.array)
	}
	return a
}

func (a *TIntArray) SortFunc(less func(v1, v2 int) bool) *TIntArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	sort.Slice(a.array, func(i, j int) bool {
		return less(a.array[i], a.array[j])
	})
	return a
}

func (a *TIntArray) InsertBefore(index int, value int) *TIntArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	rear := append([]int{}, a.array[index:]...)
	a.array = append(a.array[0:index], value)
	a.array = append(a.array, rear...)
	return a
}

func (a *TIntArray) InsertAfter(index int, value int) *TIntArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	rear := append([]int{}, a.array[index+1:]...)
	a.array = append(a.array[0:index+1], value)
	a.array = append(a.array, rear...)
	return a
}

func (a *TIntArray) Remove(index int) int {
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

func (a *TIntArray) PushLeft(value ...int) *TIntArray {
	a.mu.Lock()
	a.array = append(value, a.array...)
	a.mu.Unlock()
	return a
}

func (a *TIntArray) PushRight(value ...int) *TIntArray {
	a.mu.Lock()
	a.array = append(a.array, value...)
	a.mu.Unlock()
	return a
}

func (a *TIntArray) PopLeft() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	value := a.array[0]
	a.array = a.array[1:]
	return value
}

func (a *TIntArray) PopRight() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	index := len(a.array) - 1
	value := a.array[index]
	a.array = a.array[:index]
	return value
}

func (a *TIntArray) PopRand() int {
	return a.Remove(xrand.Intn(len(a.array)))
}

func (a *TIntArray) PopRands(size int) []int {
	a.mu.Lock()
	defer a.mu.Unlock()
	if size > len(a.array) {
		size = len(a.array)
	}
	array := make([]int, size)
	for i := 0; i < size; i++ {
		index := xrand.Intn(len(a.array))
		array[i] = a.array[index]
		a.array = append(a.array[:index], a.array[index+1:]...)
	}
	return array
}

func (a *TIntArray) PopLefts(size int) []int {
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

func (a *TIntArray) PopRights(size int) []int {
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

func (a *TIntArray) Range(start int, end ...int) []int {
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
	array := ([]int)(nil)
	if a.mu.IsSafe() {
		array = make([]int, offsetEnd-start)
		copy(array, a.array[start:offsetEnd])
	} else {
		array = a.array[start:offsetEnd]
	}
	return array
}

func (a *TIntArray) SubSlice(offset int, length ...int) []int {
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
		s := make([]int, size)
		copy(s, a.array[offset:])
		return s
	} else {
		return a.array[offset:end]
	}
}

func (a *TIntArray) Append(value ...int) *TIntArray {
	a.mu.Lock()
	a.array = append(a.array, value...)
	a.mu.Unlock()
	return a
}

func (a *TIntArray) Len() int {
	a.mu.RLock()
	length := len(a.array)
	a.mu.RUnlock()
	return length
}

func (a *TIntArray) Slice() []int {
	array := ([]int)(nil)
	if a.mu.IsSafe() {
		a.mu.RLock()
		defer a.mu.RUnlock()
		array = make([]int, len(a.array))
		copy(array, a.array)
	} else {
		array = a.array
	}
	return array
}

func (a *TIntArray) Clone() (newArray *TIntArray) {
	a.mu.RLock()
	array := make([]int, len(a.array))
	copy(array, a.array)
	a.mu.RUnlock()
	return NewIntArrayFrom(array, !a.mu.IsSafe())
}

func (a *TIntArray) Clear() *TIntArray {
	a.mu.Lock()
	if len(a.array) > 0 {
		a.array = make([]int, 0)
	}
	a.mu.Unlock()
	return a
}

func (a *TIntArray) Contains(value int) bool {
	return a.Search(value) != -1
}

func (a *TIntArray) Search(value int) int {
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

func (a *TIntArray) Unique() *TIntArray {
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

func (a *TIntArray) LockFunc(f func(array []int)) *TIntArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	f(a.array)
	return a
}

func (a *TIntArray) RLockFunc(f func(array []int)) *TIntArray {
	a.mu.RLock()
	defer a.mu.RUnlock()
	f(a.array)
	return a
}

func (a *TIntArray) Merge(array interface{}) *TIntArray {
	switch v := array.(type) {
	case *TArray:
		a.Append(t.Ints(v.Slice())...)
	case *TIntArray:
		a.Append(t.Ints(v.Slice())...)
	case *TStringArray:
		a.Append(t.Ints(v.Slice())...)
	case *TSortedArray:
		a.Append(t.Ints(v.Slice())...)
	case *TSortedIntArray:
		a.Append(t.Ints(v.Slice())...)
	case *TSortedStringArray:
		a.Append(t.Ints(v.Slice())...)
	default:
		a.Append(t.Ints(array)...)
	}
	return a
}

func (a *TIntArray) Fill(startIndex int, num int, value int) *TIntArray {
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

func (a *TIntArray) Chunk(size int) [][]int {
	if size < 1 {
		return nil
	}
	a.mu.RLock()
	defer a.mu.RUnlock()
	length := len(a.array)
	chunks := int(math.Ceil(float64(length) / float64(size)))
	var n [][]int
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

func (a *TIntArray) Pad(size int, value int) *TIntArray {
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
	tmp := make([]int, n)
	for i := 0; i < n; i++ {
		tmp[i] = value
	}
	if size > 0 {
		a.array = append(a.array, tmp...)
	} else {
		a.array = append(tmp, a.array...)
	}
	return a
}

func (a *TIntArray) Rand() int {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.array[xrand.Intn(len(a.array))]
}

func (a *TIntArray) Rands(size int) []int {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if size > len(a.array) {
		size = len(a.array)
	}
	n := make([]int, size)
	for i, v := range xrand.Perm(len(a.array)) {
		n[i] = a.array[v]
		if i == size-1 {
			break
		}
	}
	return n
}

func (a *TIntArray) Shuffle() *TIntArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	for i, v := range xrand.Perm(len(a.array)) {
		a.array[i], a.array[v] = a.array[v], a.array[i]
	}
	return a
}

func (a *TIntArray) Reverse() *TIntArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	for i, j := 0, len(a.array)-1; i < j; i, j = i+1, j-1 {
		a.array[i], a.array[j] = a.array[j], a.array[i]
	}
	return a
}

func (a *TIntArray) Join(glue string) string {
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

func (a *TIntArray) CountValues() map[int]int {
	m := make(map[int]int)
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, v := range a.array {
		m[v]++
	}
	return m
}

func (a *TIntArray) String() string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	jsonContent, _ := json.Marshal(a.array)
	return string(jsonContent)
}

func (a *TIntArray) MarshalJSON() ([]byte, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return json.Marshal(a.array)
}
