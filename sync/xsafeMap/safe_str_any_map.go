package xsafeMap

import (
	"encoding/json"
	"github.com/go-xe2/x/core/rwmutex"
	. "github.com/go-xe2/x/sync/type"
	"github.com/go-xe2/x/type/t"
)

type (
	StrAnyMapForeachFunc func(k string, v interface{}) bool
	StrAnyMapLockFunc    func(m map[string]interface{})
)

type TStrAnyMap struct {
	mu   *rwmutex.RWMutex
	data map[string]interface{}
}

func NewStrAnyMap(unsafe ...bool) *TStrAnyMap {
	return &TStrAnyMap{
		mu:   rwmutex.New(unsafe...),
		data: make(map[string]interface{}),
	}
}

func NewStrAnyMapFrom(data map[string]interface{}, unsafe ...bool) *TStrAnyMap {
	return &TStrAnyMap{
		mu:   rwmutex.New(unsafe...),
		data: data,
	}
}

func (m *TStrAnyMap) Foreach(fn StrAnyMapForeachFunc) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for k, v := range m.data {
		if !fn(k, v) {
			break
		}
	}
}

func (m *TStrAnyMap) Clone(unsafe ...bool) *TStrAnyMap {
	return NewStrAnyMapFrom(m.data, unsafe...)
}

func (m *TStrAnyMap) Map() map[string]interface{} {
	return m.data
}

func (m *TStrAnyMap) Set(key string, val interface{}) {
	m.mu.Lock()
	m.data[key] = val
	m.mu.Unlock()
}

func (m *TStrAnyMap) Sets(data map[string]interface{}) {
	m.mu.Lock()
	for k, v := range data {
		m.data[k] = v
	}
	m.mu.Unlock()
}

func (m *TStrAnyMap) Search(key string) (value interface{}, found bool) {
	m.mu.RLock()
	value, found = m.data[key]
	m.mu.RUnlock()
	return
}

func (m *TStrAnyMap) Get(key string) interface{} {
	m.mu.RLock()
	val, _ := m.data[key]
	m.mu.RUnlock()
	return val
}

func (m *TStrAnyMap) doSetWithLockCheck(key string, value interface{}) interface{} {
	m.mu.Lock()
	defer m.mu.Unlock()
	if v, ok := m.data[key]; ok {
		return v
	}
	if f, ok := value.(GetAnyFunc); ok {
		value = f()
	}
	m.data[key] = value
	return value
}

func (m *TStrAnyMap) GetOrSet(key string, value interface{}) interface{} {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, value)
	} else {
		return v
	}
}

func (m *TStrAnyMap) GetOrSetFunc(key string, fn GetAnyFunc) interface{} {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, fn())
	} else {
		return v
	}
}

func (m *TStrAnyMap) GetOrSetFuncLock(key string, fn GetAnyFunc) interface{} {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, fn)
	} else {
		return v
	}
}

func (m *TStrAnyMap) SetIfNotExist(key string, value interface{}) bool {
	if !m.Contains(key) {
		m.doSetWithLockCheck(key, value)
		return true
	}
	return false
}

func (m *TStrAnyMap) SetIfNotExistFunc(key string, fn GetAnyFunc) bool {
	if !m.Contains(key) {
		m.doSetWithLockCheck(key, fn())
		return true
	}
	return false
}

func (m *TStrAnyMap) SetIfNotExistFuncLock(key string, fn GetAnyFunc) bool {
	if !m.Contains(key) {
		m.doSetWithLockCheck(key, fn)
		return true
	}
	return false
}

func (m *TStrAnyMap) Remove(key string) interface{} {
	m.mu.Lock()
	val, exists := m.data[key]
	if exists {
		delete(m.data, key)
	}
	m.mu.Unlock()
	return val
}

func (m *TStrAnyMap) Removes(keys []string) {
	m.mu.Lock()
	for _, key := range keys {
		delete(m.data, key)
	}
	m.mu.Unlock()
}

func (m *TStrAnyMap) Keys() []string {
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

func (m *TStrAnyMap) Values() []interface{} {
	m.mu.RLock()
	values := make([]interface{}, len(m.data))
	index := 0
	for _, value := range m.data {
		values[index] = value
		index++
	}
	m.mu.RUnlock()
	return values
}

func (m *TStrAnyMap) Contains(key string) bool {
	m.mu.RLock()
	_, exists := m.data[key]
	m.mu.RUnlock()
	return exists
}

func (m *TStrAnyMap) Size() int {
	m.mu.RLock()
	length := len(m.data)
	m.mu.RUnlock()
	return length
}

func (m *TStrAnyMap) IsEmpty() bool {
	m.mu.RLock()
	empty := len(m.data) == 0
	m.mu.RUnlock()
	return empty
}

func (m *TStrAnyMap) Clear() {
	m.mu.Lock()
	m.data = make(map[string]interface{})
	m.mu.Unlock()
}

func (m *TStrAnyMap) LockFunc(fn StrAnyMapLockFunc) {
	m.mu.Lock()
	defer m.mu.Unlock()
	fn(m.data)
}

func (m *TStrAnyMap) RLockFunc(fn StrAnyMapLockFunc) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	fn(m.data)
}

func (m *TStrAnyMap) Flip() {
	m.mu.Lock()
	defer m.mu.Unlock()
	n := make(map[string]interface{}, len(m.data))
	for k, v := range m.data {
		n[t.String(v)] = k
	}
	m.data = n
}

func (m *TStrAnyMap) Merge(other *TStrAnyMap) {
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

func (m *TStrAnyMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Map())
}

func (m *TStrAnyMap) String() string {
	b, err := m.MarshalJSON()
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func (m *TStrAnyMap) GetVar(key string) *TVar {
	return NewVar(m.Get(key), true)
}

func (m *TStrAnyMap) GetVarOrSet(key string, value interface{}) *TVar {
	return NewVar(m.GetOrSet(key, value), true)
}

func (m *TStrAnyMap) GetVarOrSetFunc(key string, fn GetAnyFunc) *TVar {
	return NewVar(m.GetOrSetFunc(key, fn), true)
}

func (m *TStrAnyMap) GetVarOrSetFuncLock(key string, fn GetAnyFunc) *TVar {
	return NewVar(m.GetOrSetFuncLock(key, fn), true)
}
