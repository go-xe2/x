package xsafeMap

import (
	"encoding/json"
	"github.com/go-xe2/x/core/rwmutex"
	. "github.com/go-xe2/x/sync/type"
	"github.com/go-xe2/x/type/t"
)

type (
	AnyAnyMapForeachFunc func(k interface{}, v interface{}) bool
	AnyAnyMapLockFunc    func(m map[interface{}]interface{})
)

type TAnyAnyMap struct {
	mu   *rwmutex.RWMutex
	data map[interface{}]interface{}
}

func NewAnyAnyMap(unsafe ...bool) *TAnyAnyMap {
	return &TAnyAnyMap{
		mu:   rwmutex.New(unsafe...),
		data: make(map[interface{}]interface{}),
	}
}

func NewAnyAnyMapFrom(data map[interface{}]interface{}, unsafe ...bool) *TAnyAnyMap {
	return &TAnyAnyMap{
		mu:   rwmutex.New(unsafe...),
		data: data,
	}
}

func (m *TAnyAnyMap) Foreach(fn AnyAnyMapForeachFunc) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for k, v := range m.data {
		if !fn(k, v) {
			break
		}
	}
}

func (m *TAnyAnyMap) Clone(unsafe ...bool) *TAnyAnyMap {
	return NewAnyAnyMapFrom(m.data, unsafe...)
}

func (m *TAnyAnyMap) Map() map[interface{}]interface{} {
	return m.data
}

func (m *TAnyAnyMap) Set(key interface{}, val interface{}) {
	m.mu.Lock()
	m.data[key] = val
	m.mu.Unlock()
}

func (m *TAnyAnyMap) Sets(data map[interface{}]interface{}) {
	m.mu.Lock()
	for k, v := range data {
		m.data[k] = v
	}
	m.mu.Unlock()
}

func (m *TAnyAnyMap) Search(key interface{}) (value interface{}, found bool) {
	m.mu.RLock()
	value, found = m.data[key]
	m.mu.RUnlock()
	return
}

func (m *TAnyAnyMap) Get(key interface{}) interface{} {
	m.mu.RLock()
	val, _ := m.data[key]
	m.mu.RUnlock()
	return val
}

func (m *TAnyAnyMap) doSetWithLockCheck(key interface{}, value interface{}) interface{} {
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

func (m *TAnyAnyMap) GetOrSet(key interface{}, value interface{}) interface{} {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, value)
	} else {
		return v
	}
}

func (m *TAnyAnyMap) GetOrSetFunc(key interface{}, fn GetAnyFunc) interface{} {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, fn())
	} else {
		return v
	}
}

func (m *TAnyAnyMap) GetOrSetFuncLock(key interface{}, fn GetAnyFunc) interface{} {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, fn)
	} else {
		return v
	}
}

func (m *TAnyAnyMap) SetIfNotExist(key interface{}, value interface{}) bool {
	if !m.Contains(key) {
		m.doSetWithLockCheck(key, value)
		return true
	}
	return false
}

func (m *TAnyAnyMap) SetIfNotExistFunc(key interface{}, fn GetAnyFunc) bool {
	if !m.Contains(key) {
		m.doSetWithLockCheck(key, fn())
		return true
	}
	return false
}

func (m *TAnyAnyMap) SetIfNotExistFuncLock(key interface{}, fn GetAnyFunc) bool {
	if !m.Contains(key) {
		m.doSetWithLockCheck(key, fn)
		return true
	}
	return false
}

func (m *TAnyAnyMap) Remove(key interface{}) interface{} {
	m.mu.Lock()
	val, exists := m.data[key]
	if exists {
		delete(m.data, key)
	}
	m.mu.Unlock()
	return val
}

func (m *TAnyAnyMap) Removes(keys []interface{}) {
	m.mu.Lock()
	for _, key := range keys {
		delete(m.data, key)
	}
	m.mu.Unlock()
}

func (m *TAnyAnyMap) Keys() []interface{} {
	m.mu.RLock()
	keys := make([]interface{}, len(m.data))
	index := 0
	for key := range m.data {
		keys[index] = key
		index++
	}
	m.mu.RUnlock()
	return keys
}

func (m *TAnyAnyMap) Values() []interface{} {
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

func (m *TAnyAnyMap) Contains(key interface{}) bool {
	m.mu.RLock()
	_, exists := m.data[key]
	m.mu.RUnlock()
	return exists
}

func (m *TAnyAnyMap) Size() int {
	m.mu.RLock()
	length := len(m.data)
	m.mu.RUnlock()
	return length
}

func (m *TAnyAnyMap) IsEmpty() bool {
	m.mu.RLock()
	empty := len(m.data) == 0
	m.mu.RUnlock()
	return empty
}

func (m *TAnyAnyMap) Clear() {
	m.mu.Lock()
	m.data = make(map[interface{}]interface{})
	m.mu.Unlock()
}

func (m *TAnyAnyMap) LockFunc(f AnyAnyMapLockFunc) {
	m.mu.Lock()
	defer m.mu.Unlock()
	f(m.data)
}

func (m *TAnyAnyMap) RLockFunc(f AnyAnyMapLockFunc) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	f(m.data)
}

func (m *TAnyAnyMap) Flip() {
	m.mu.Lock()
	defer m.mu.Unlock()
	n := make(map[interface{}]interface{}, len(m.data))
	for k, v := range m.data {
		n[v] = k
	}
	m.data = n
}

func (m *TAnyAnyMap) Merge(other *TAnyAnyMap) {
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

func (m *TAnyAnyMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.New(m.Map()).MapStrAny())
}

func (m *TAnyAnyMap) String() string {
	b, err := m.MarshalJSON()
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func (m *TAnyAnyMap) GetVar(key interface{}) *TVar {
	return NewVar(m.Get(key), true)
}

func (m *TAnyAnyMap) GetVarOrSet(key interface{}, value interface{}) *TVar {
	return NewVar(m.GetOrSet(key, value), true)
}

func (m *TAnyAnyMap) GetVarOrSetFunc(key interface{}, fn GetAnyFunc) *TVar {
	return NewVar(m.GetOrSetFunc(key, fn), true)
}

func (m *TAnyAnyMap) GetVarOrSetFuncLock(key interface{}, fn GetAnyFunc) *TVar {
	return NewVar(m.GetOrSetFuncLock(key, fn), true)
}
