package xarray

import (
	"bytes"
	"encoding/json"
	"github.com/go-xe2/x/core/rwmutex"
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/utils/xrand"
	"math"
	"sort"
	"strings"
)

type TStringArray struct {
	mu    *rwmutex.RWMutex
	array []string
}

func NewStringArray(unsafe ...bool) *TStringArray {
	return NewStringArraySize(0, 0, unsafe...)
}

func NewStringArraySize(size int, cap int, unsafe ...bool) *TStringArray {
	return &TStringArray{
		mu:    rwmutex.New(unsafe...),
		array: make([]string, size, cap),
	}
}

func NewStringArrayFrom(array []string, unsafe ...bool) *TStringArray {
	return &TStringArray{
		mu:    rwmutex.New(unsafe...),
		array: array,
	}
}

func NewStringArrayFromCopy(array []string, unsafe ...bool) *TStringArray {
	newArray := make([]string, len(array))
	copy(newArray, array)
	return &TStringArray{
		mu:    rwmutex.New(unsafe...),
		array: newArray,
	}
}

func (a *TStringArray) Get(index int) string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	value := a.array[index]
	return value
}

func (a *TStringArray) Set(index int, value string) *TStringArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.array[index] = value
	return a
}

func (a *TStringArray) SetArray(array []string) *TStringArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.array = array
	return a
}

func (a *TStringArray) Replace(array []string) *TStringArray {
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

func (a *TStringArray) Sum() (sum int) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, v := range a.array {
		sum += t.Int(v)
	}
	return
}

func (a *TStringArray) Sort(reverse ...bool) *TStringArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	if len(reverse) > 0 && reverse[0] {
		sort.Slice(a.array, func(i, j int) bool {
			if strings.Compare(a.array[i], a.array[j]) < 0 {
				return false
			}
			return true
		})
	} else {
		sort.Strings(a.array)
	}
	return a
}

func (a *TStringArray) SortFunc(less func(v1, v2 string) bool) *TStringArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	sort.Slice(a.array, func(i, j int) bool {
		return less(a.array[i], a.array[j])
	})
	return a
}

func (a *TStringArray) InsertBefore(index int, value string) *TStringArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	rear := append([]string{}, a.array[index:]...)
	a.array = append(a.array[0:index], value)
	a.array = append(a.array, rear...)
	return a
}

func (a *TStringArray) InsertAfter(index int, value string) *TStringArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	rear := append([]string{}, a.array[index+1:]...)
	a.array = append(a.array[0:index+1], value)
	a.array = append(a.array, rear...)
	return a
}

func (a *TStringArray) Remove(index int) string {
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

func (a *TStringArray) PushLeft(value ...string) *TStringArray {
	a.mu.Lock()
	a.array = append(value, a.array...)
	a.mu.Unlock()
	return a
}

func (a *TStringArray) PushRight(value ...string) *TStringArray {
	a.mu.Lock()
	a.array = append(a.array, value...)
	a.mu.Unlock()
	return a
}

func (a *TStringArray) PopLeft() string {
	a.mu.Lock()
	defer a.mu.Unlock()
	value := a.array[0]
	a.array = a.array[1:]
	return value
}

func (a *TStringArray) PopRight() string {
	a.mu.Lock()
	defer a.mu.Unlock()
	index := len(a.array) - 1
	value := a.array[index]
	a.array = a.array[:index]
	return value
}

func (a *TStringArray) PopRand() string {
	return a.Remove(xrand.Intn(len(a.array)))
}

func (a *TStringArray) PopRands(size int) []string {
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

func (a *TStringArray) PopLefts(size int) []string {
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

func (a *TStringArray) PopRights(size int) []string {
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

func (a *TStringArray) Range(start int, end ...int) []string {
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

func (a *TStringArray) SubSlice(offset int, length ...int) []string {
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

func (a *TStringArray) Append(value ...string) *TStringArray {
	a.mu.Lock()
	a.array = append(a.array, value...)
	a.mu.Unlock()
	return a
}

func (a *TStringArray) Len() int {
	a.mu.RLock()
	length := len(a.array)
	a.mu.RUnlock()
	return length
}

func (a *TStringArray) Slice() []string {
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

func (a *TStringArray) Clone() (newArray *TStringArray) {
	a.mu.RLock()
	array := make([]string, len(a.array))
	copy(array, a.array)
	a.mu.RUnlock()
	return NewStringArrayFrom(array, !a.mu.IsSafe())
}

func (a *TStringArray) Clear() *TStringArray {
	a.mu.Lock()
	if len(a.array) > 0 {
		a.array = make([]string, 0)
	}
	a.mu.Unlock()
	return a
}

func (a *TStringArray) Contains(value string) bool {
	return a.Search(value) != -1
}

func (a *TStringArray) Search(value string) int {
	if len(a.array) == 0 {
		return -1
	}
	a.mu.RLock()
	result := -1
	for index, v := range a.array {
		if strings.Compare(v, value) == 0 {
			result = index
			break
		}
	}
	a.mu.RUnlock()
	return result
}

func (a *TStringArray) Unique() *TStringArray {
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

func (a *TStringArray) LockFunc(f func(array []string)) *TStringArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	f(a.array)
	return a
}

func (a *TStringArray) RLockFunc(f func(array []string)) *TStringArray {
	a.mu.RLock()
	defer a.mu.RUnlock()
	f(a.array)
	return a
}

func (a *TStringArray) Merge(array interface{}) *TStringArray {
	switch v := array.(type) {
	case *TArray:
		a.Append(t.Strings(v.Slice())...)
	case *TIntArray:
		a.Append(t.Strings(v.Slice())...)
	case *TStringArray:
		a.Append(t.Strings(v.Slice())...)
	case *TSortedArray:
		a.Append(t.Strings(v.Slice())...)
	case *TSortedIntArray:
		a.Append(t.Strings(v.Slice())...)
	case *TSortedStringArray:
		a.Append(t.Strings(v.Slice())...)
	default:
		a.Append(t.Strings(array)...)
	}
	return a
}

func (a *TStringArray) Fill(startIndex int, num int, value string) *TStringArray {
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

func (a *TStringArray) Chunk(size int) [][]string {
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

func (a *TStringArray) Pad(size int, value string) *TStringArray {
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
	tmp := make([]string, n)
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

func (a *TStringArray) Rand() string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.array[xrand.Intn(len(a.array))]
}

func (a *TStringArray) Rands(size int) []string {
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

func (a *TStringArray) Shuffle() *TStringArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	for i, v := range xrand.Perm(len(a.array)) {
		a.array[i], a.array[v] = a.array[v], a.array[i]
	}
	return a
}

func (a *TStringArray) Reverse() *TStringArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	for i, j := 0, len(a.array)-1; i < j; i, j = i+1, j-1 {
		a.array[i], a.array[j] = a.array[j], a.array[i]
	}
	return a
}

func (a *TStringArray) Join(glue string) string {
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

func (a *TStringArray) CountValues() map[string]int {
	m := make(map[string]int)
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, v := range a.array {
		m[v]++
	}
	return m
}

func (a *TStringArray) String() string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	jsonContent, _ := json.Marshal(a.array)
	return string(jsonContent)
}

func (a *TStringArray) MarshalJSON() ([]byte, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return json.Marshal(a.array)
}
