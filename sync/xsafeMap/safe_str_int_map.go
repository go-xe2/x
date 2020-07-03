package xsafeMap

import (
	"encoding/json"
	"github.com/go-xe2/x/core/rwmutex"
	. "github.com/go-xe2/x/sync/type"
	"github.com/go-xe2/x/type/t"
)

type (
	StrIntMapForeachFunc func(k string, v int) bool
	StrIntMapLockFunc    func(m map[string]int)
)

type TStrIntMap struct {
	mu   *rwmutex.RWMutex
	data map[string]int
}

func NewStrIntMap(unsafe ...bool) *TStrIntMap {
	return &TStrIntMap{
		mu:   rwmutex.New(unsafe...),
		data: make(map[string]int),
	}
}

func NewStrIntMapFrom(data map[string]int, unsafe ...bool) *TStrIntMap {
	return &TStrIntMap{
		mu:   rwmutex.New(unsafe...),
		data: data,
	}
}

func (m *TStrIntMap) Foreach(fn StrIntMapForeachFunc) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for k, v := range m.data {
		if !fn(k, v) {
			break
		}
	}
}

func (m *TStrIntMap) Clone(unsafe ...bool) *TStrIntMap {
	return NewStrIntMapFrom(m.data, unsafe...)
}

func (m *TStrIntMap) Map() map[string]int {
	return m.data
}

func (m *TStrIntMap) Set(key string, val int) {
	m.mu.Lock()
	m.data[key] = val
	m.mu.Unlock()
}

func (m *TStrIntMap) Sets(data map[string]int) {
	m.mu.Lock()
	for k, v := range data {
		m.data[k] = v
	}
	m.mu.Unlock()
}

func (m *TStrIntMap) Search(key string) (value int, found bool) {
	m.mu.RLock()
	value, found = m.data[key]
	m.mu.RUnlock()
	return
}

func (m *TStrIntMap) Get(key string) int {
	m.mu.RLock()
	val, _ := m.data[key]
	m.mu.RUnlock()
	return val
}

func (m *TStrIntMap) doSetWithLockCheck(key string, value interface{}) int {
	m.mu.Lock()
	defer m.mu.Unlock()
	if v, ok := m.data[key]; ok {
		return v
	}
	if f, ok := value.(GetIntFunc); ok {
		value = f()
	}
	s := t.Int(value)
	m.data[key] = s
	return s
}

func (m *TStrIntMap) GetOrSet(key string, value interface{}) int {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, value)
	} else {
		return v
	}
}

func (m *TStrIntMap) GetOrSetFunc(key string, fn GetIntFunc) int {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, fn())
	} else {
		return v
	}
}

func (m *TStrIntMap) GetOrSetFuncLock(key string, fn GetIntFunc) int {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, fn)
	} else {
		return v
	}
}

func (m *TStrIntMap) SetIfNotExist(key string, value interface{}) bool {
	if !m.Contains(key) {
		m.doSetWithLockCheck(key, value)
		return true
	}
	return false
}

func (m *TStrIntMap) SetIfNotExistFunc(key string, fn GetIntFunc) bool {
	if !m.Contains(key) {
		m.doSetWithLockCheck(key, fn())
		return true
	}
	return false
}

func (m *TStrIntMap) SetIfNotExistFuncLock(key string, fn GetIntFunc) bool {
	if !m.Contains(key) {
		m.doSetWithLockCheck(key, fn)
		return true
	}
	return false
}

func (m *TStrIntMap) Remove(key string) int {
	m.mu.Lock()
	val, exists := m.data[key]
	if exists {
		delete(m.data, key)
	}
	m.mu.Unlock()
	return val
}

func (m *TStrIntMap) Removes(keys []string) {
	m.mu.Lock()
	for _, key := range keys {
		delete(m.data, key)
	}
	m.mu.Unlock()
}

func (m *TStrIntMap) Keys() []string {
	m.mu.RLock()
	keys := make([]string, len(m.data))
	index := 0
	for key := range m.data {
		keys[index] = key
		index++
	}
	m.mu.RUnlock()
	return keys
}

func (m *TStrIntMap) Values() []int {
	m.mu.RLock()
	values := make([]int, len(m.data))
	index := 0
	for _, value := range m.data {
		values[index] = value
		index++
	}
	m.mu.RUnlock()
	return values
}

func (m *TStrIntMap) Contains(key string) bool {
	m.mu.RLock()
	_, exists := m.data[key]
	m.mu.RUnlock()
	return exists
}

func (m *TStrIntMap) Size() int {
	m.mu.RLock()
	length := len(m.data)
	m.mu.RUnlock()
	return length
}

func (m *TStrIntMap) IsEmpty() bool {
	m.mu.RLock()
	empty := len(m.data) == 0
	m.mu.RUnlock()
	return empty
}

func (m *TStrIntMap) Clear() {
	m.mu.Lock()
	m.data = make(map[string]int)
	m.mu.Unlock()
}

func (m *TStrIntMap) LockFunc(fn StrIntMapLockFunc) {
	m.mu.Lock()
	defer m.mu.Unlock()
	fn(m.data)
}

func (m *TStrIntMap) RLockFunc(fn StrIntMapLockFunc) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	fn(m.data)
}

func (m *TStrIntMap) Merge(other *TStrIntMap) {
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

func (m *TStrIntMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Map())
}

func (m *TStrIntMap) String() string {
	b, err := m.MarshalJSON()
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func (m *TStrIntMap) GetVar(key string) *TVar {
	return NewVar(m.Get(key), true)
}

func (m *TStrIntMap) GetVarOrSet(key string, value interface{}) *TVar {
	return NewVar(m.GetOrSet(key, value), true)
}

func (m *TStrIntMap) GetVarOrSetFunc(key string, fn GetIntFunc) *TVar {
	return NewVar(m.GetOrSetFunc(key, fn), true)
}

func (m *TStrIntMap) GetVarOrSetFuncLock(key string, fn GetIntFunc) *TVar {
	return NewVar(m.GetOrSetFuncLock(key, fn), true)
}
