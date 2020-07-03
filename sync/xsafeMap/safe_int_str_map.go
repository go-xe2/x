package xsafeMap

import (
	"encoding/json"
	"github.com/go-xe2/x/core/rwmutex"
	. "github.com/go-xe2/x/sync/type"
	"github.com/go-xe2/x/type/t"
)

type (
	IntStrMapForeachFunc func(k int, v string) bool
	IntStrMapLockFunc    func(m map[int]string)
)

type TIntStrMap struct {
	mu   *rwmutex.RWMutex
	data map[int]string
}

func NewIntStrMap(unsafe ...bool) *TIntStrMap {
	return &TIntStrMap{
		mu:   rwmutex.New(unsafe...),
		data: make(map[int]string),
	}
}

func NewIntStrMapFrom(data map[int]string, unsafe ...bool) *TIntStrMap {
	return &TIntStrMap{
		mu:   rwmutex.New(unsafe...),
		data: data,
	}
}

func (m *TIntStrMap) Foreach(fn IntStrMapForeachFunc) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for k, v := range m.data {
		if !fn(k, v) {
			break
		}
	}
}

func (m *TIntStrMap) Clone(unsafe ...bool) *TIntStrMap {
	return NewIntStrMapFrom(m.data, unsafe...)
}

func (m *TIntStrMap) Map() map[int]string {
	return m.data
}

func (m *TIntStrMap) Set(key int, val string) {
	m.mu.Lock()
	m.data[key] = val
	m.mu.Unlock()
}

func (m *TIntStrMap) Sets(data map[int]string) {
	m.mu.Lock()
	for k, v := range data {
		m.data[k] = v
	}
	m.mu.Unlock()
}

func (m *TIntStrMap) Search(key int) (value string, found bool) {
	m.mu.RLock()
	value, found = m.data[key]
	m.mu.RUnlock()
	return
}

func (m *TIntStrMap) Get(key int) string {
	m.mu.RLock()
	val, _ := m.data[key]
	m.mu.RUnlock()
	return val
}

func (m *TIntStrMap) doSetWithLockCheck(key int, value interface{}) string {
	m.mu.Lock()
	defer m.mu.Unlock()
	if v, ok := m.data[key]; ok {
		return v
	}
	if f, ok := value.(GetStrFunc); ok {
		value = f()
	}
	s := t.String(value)
	m.data[key] = s
	return s
}

func (m *TIntStrMap) GetOrSet(key int, value interface{}) string {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, value)
	} else {
		return v
	}
}

func (m *TIntStrMap) GetOrSetFunc(key int, fn GetStrFunc) string {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, fn())
	} else {
		return v
	}
}

func (m *TIntStrMap) GetOrSetFuncLock(key int, fn GetStrFunc) string {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, fn)
	} else {
		return v
	}
}

func (m *TIntStrMap) SetIfNotExist(key int, value interface{}) bool {
	if !m.Contains(key) {
		m.doSetWithLockCheck(key, value)
		return true
	}
	return false
}

func (m *TIntStrMap) SetIfNotExistFunc(key int, fn GetStrFunc) bool {
	if !m.Contains(key) {
		m.doSetWithLockCheck(key, fn())
		return true
	}
	return false
}

func (m *TIntStrMap) SetIfNotExistFuncLock(key int, fn GetStrFunc) bool {
	if !m.Contains(key) {
		m.doSetWithLockCheck(key, fn)
		return true
	}
	return false
}

func (m *TIntStrMap) Remove(key int) string {
	m.mu.Lock()
	val, exists := m.data[key]
	if exists {
		delete(m.data, key)
	}
	m.mu.Unlock()
	return val
}

func (m *TIntStrMap) Removes(keys []int) {
	m.mu.Lock()
	for _, key := range keys {
		delete(m.data, key)
	}
	m.mu.Unlock()
}

func (m *TIntStrMap) Keys() []int {
	m.mu.RLock()
	keys := make([]int, len(m.data))
	index := 0
	for key := range m.data {
		keys[index] = key
		index++
	}
	m.mu.RUnlock()
	return keys
}

func (m *TIntStrMap) Values() []string {
	m.mu.RLock()
	values := make([]string, len(m.data))
	index := 0
	for _, value := range m.data {
		values[index] = value
		index++
	}
	m.mu.RUnlock()
	return values
}

func (m *TIntStrMap) Contains(key int) bool {
	m.mu.RLock()
	_, exists := m.data[key]
	m.mu.RUnlock()
	return exists
}

func (m *TIntStrMap) Size() int {
	m.mu.RLock()
	length := len(m.data)
	m.mu.RUnlock()
	return length
}

func (m *TIntStrMap) IsEmpty() bool {
	m.mu.RLock()
	empty := len(m.data) == 0
	m.mu.RUnlock()
	return empty
}

func (m *TIntStrMap) Clear() {
	m.mu.Lock()
	m.data = make(map[int]string)
	m.mu.Unlock()
}

func (m *TIntStrMap) LockFunc(fn IntStrMapLockFunc) {
	m.mu.Lock()
	defer m.mu.Unlock()
	fn(m.data)
}

func (m *TIntStrMap) RLockFunc(fn IntStrMapLockFunc) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	fn(m.data)
}

func (m *TIntStrMap) Merge(other *TIntStrMap) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if other != m {
		other.mu.Lock()
		defer other.mu.Unlock()
	}
	data := other.data
	for k, v := range data {
		m.data[k] = v
	}
}

func (m *TIntStrMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Map())
}

func (m *TIntStrMap) String() string {
	b, err := m.MarshalJSON()
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func (m *TIntStrMap) GetVar(key int) *TVar {
	return NewVar(m.Get(key), true)
}

func (m *TIntStrMap) GetVarOrSet(key int, value interface{}) *TVar {
	return NewVar(m.GetOrSet(key, value), true)
}

func (m *TIntStrMap) GetVarOrSetFunc(key int, fn GetStrFunc) *TVar {
	return NewVar(m.GetOrSetFunc(key, fn), true)
}

func (m *TIntStrMap) GetVarOrSetFuncLock(key int, fn GetStrFunc) *TVar {
	return NewVar(m.GetOrSetFuncLock(key, fn), true)
}
