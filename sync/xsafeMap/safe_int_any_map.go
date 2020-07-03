package xsafeMap

import (
	"encoding/json"
	"github.com/go-xe2/x/core/rwmutex"
	_type "github.com/go-xe2/x/sync/type"
	"github.com/go-xe2/x/type/t"
)

type (
	IntAnyMapForeachFunc func(k int, v interface{}) bool
	IntAnyMapLockFunc    func(m map[int]interface{})
)

type TIntAnyMap struct {
	mu   *rwmutex.RWMutex
	data map[int]interface{}
}

func NewIntAnyMap(unsafe ...bool) *TIntAnyMap {
	return &TIntAnyMap{
		mu:   rwmutex.New(unsafe...),
		data: make(map[int]interface{}),
	}
}

func NewIntAnyMapFrom(data map[int]interface{}, unsafe ...bool) *TIntAnyMap {
	return &TIntAnyMap{
		mu:   rwmutex.New(unsafe...),
		data: data,
	}
}

func (m *TIntAnyMap) Foreach(fn IntAnyMapForeachFunc) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for k, v := range m.data {
		if !fn(k, v) {
			break
		}
	}
}

func (m *TIntAnyMap) Clone(unsafe ...bool) *TIntAnyMap {
	return NewIntAnyMapFrom(m.data, unsafe...)
}

func (m *TIntAnyMap) Map() map[int]interface{} {
	return m.data
}

func (m *TIntAnyMap) Set(key int, val interface{}) {
	m.mu.Lock()
	m.data[key] = val
	m.mu.Unlock()
}

func (m *TIntAnyMap) Sets(data map[int]interface{}) {
	m.mu.Lock()
	for k, v := range data {
		m.data[k] = v
	}
	m.mu.Unlock()
}

func (m *TIntAnyMap) Search(key int) (value interface{}, found bool) {
	m.mu.RLock()
	value, found = m.data[key]
	m.mu.RUnlock()
	return
}

func (m *TIntAnyMap) Get(key int) interface{} {
	m.mu.RLock()
	val, _ := m.data[key]
	m.mu.RUnlock()
	return val
}

func (m *TIntAnyMap) doSetWithLockCheck(key int, value interface{}) interface{} {
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

func (m *TIntAnyMap) GetOrSet(key int, value interface{}) interface{} {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, value)
	} else {
		return v
	}
}

func (m *TIntAnyMap) GetOrSetFunc(key int, fn GetAnyFunc) interface{} {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, fn())
	} else {
		return v
	}
}

func (m *TIntAnyMap) GetOrSetFuncLock(key int, fn GetAnyFunc) interface{} {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, fn)
	} else {
		return v
	}
}

func (m *TIntAnyMap) SetIfNotExist(key int, value interface{}) bool {
	if !m.Contains(key) {
		m.doSetWithLockCheck(key, value)
		return true
	}
	return false
}

func (m *TIntAnyMap) SetIfNotExistFunc(key int, fn GetAnyFunc) bool {
	if !m.Contains(key) {
		m.doSetWithLockCheck(key, fn())
		return true
	}
	return false
}

func (m *TIntAnyMap) SetIfNotExistFuncLock(key int, fn GetAnyFunc) bool {
	if !m.Contains(key) {
		m.doSetWithLockCheck(key, fn)
		return true
	}
	return false
}

func (m *TIntAnyMap) Remove(key int) interface{} {
	m.mu.Lock()
	val, exists := m.data[key]
	if exists {
		delete(m.data, key)
	}
	m.mu.Unlock()
	return val
}

func (m *TIntAnyMap) Removes(keys []int) {
	m.mu.Lock()
	for _, key := range keys {
		delete(m.data, key)
	}
	m.mu.Unlock()
}

func (m *TIntAnyMap) Keys() []int {
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

func (m *TIntAnyMap) Values() []interface{} {
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

func (m *TIntAnyMap) Contains(key int) bool {
	m.mu.RLock()
	_, exists := m.data[key]
	m.mu.RUnlock()
	return exists
}

func (m *TIntAnyMap) Size() int {
	m.mu.RLock()
	length := len(m.data)
	m.mu.RUnlock()
	return length
}

func (m *TIntAnyMap) IsEmpty() bool {
	m.mu.RLock()
	empty := len(m.data) == 0
	m.mu.RUnlock()
	return empty
}

func (m *TIntAnyMap) Clear() {
	m.mu.Lock()
	m.data = make(map[int]interface{})
	m.mu.Unlock()
}

func (m *TIntAnyMap) LockFunc(fn IntAnyMapLockFunc) {
	m.mu.Lock()
	defer m.mu.Unlock()
	fn(m.data)
}

func (m *TIntAnyMap) RLockFunc(fn IntAnyMapLockFunc) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	fn(m.data)
}

func (m *TIntAnyMap) Flip() {
	m.mu.Lock()
	defer m.mu.Unlock()
	n := make(map[int]interface{}, len(m.data))
	for k, v := range m.data {
		n[t.Int(v)] = k
	}
	m.data = n
}

func (m *TIntAnyMap) Merge(other *TIntAnyMap) {
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

func (m *TIntAnyMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Map())
}

func (m *TIntAnyMap) String() string {
	b, err := m.MarshalJSON()
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func (m *TIntAnyMap) GetVar(key int) *_type.TVar {
	return _type.NewVar(m.Get(key), true)
}

func (m *TIntAnyMap) GetVarOrSet(key int, value interface{}) *_type.TVar {
	return _type.NewVar(m.GetOrSet(key, value), true)
}

func (m *TIntAnyMap) GetVarOrSetFunc(key int, fn GetAnyFunc) *_type.TVar {
	return _type.NewVar(m.GetOrSetFunc(key, fn), true)
}

func (m *TIntAnyMap) GetVarOrSetFuncLock(key int, fn GetAnyFunc) *_type.TVar {
	return _type.NewVar(m.GetOrSetFuncLock(key, fn), true)
}
