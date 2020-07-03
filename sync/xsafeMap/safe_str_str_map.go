package xsafeMap

import (
	"encoding/json"
	"github.com/go-xe2/x/core/rwmutex"
	. "github.com/go-xe2/x/sync/type"
	"github.com/go-xe2/x/type/t"
)

type (
	StrStrMapForeachFunc func(k string, v string) bool
	StrStrMapLockFunc    func(m map[string]string)
)

type TStrStrMap struct {
	mu   *rwmutex.RWMutex
	data map[string]string
}

func NewStrStrMap(unsafe ...bool) *TStrStrMap {
	return &TStrStrMap{
		mu:   rwmutex.New(unsafe...),
		data: make(map[string]string),
	}
}

func NewStrStrMapFrom(data map[string]string, unsafe ...bool) *TStrStrMap {
	return &TStrStrMap{
		mu:   rwmutex.New(unsafe...),
		data: data,
	}
}

func (m *TStrStrMap) Foreach(fn StrStrMapForeachFunc) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for k, v := range m.data {
		if !fn(k, v) {
			break
		}
	}
}

func (m *TStrStrMap) Clone(unsafe ...bool) *TStrStrMap {
	return NewStrStrMapFrom(m.data, unsafe...)
}

func (m *TStrStrMap) Map() map[string]string {
	return m.data
}

func (m *TStrStrMap) Set(key string, val string) {
	m.mu.Lock()
	m.data[key] = val
	m.mu.Unlock()
}

func (m *TStrStrMap) Sets(data map[string]string) {
	m.mu.Lock()
	for k, v := range data {
		m.data[k] = v
	}
	m.mu.Unlock()
}

func (m *TStrStrMap) Search(key string) (value string, found bool) {
	m.mu.RLock()
	value, found = m.data[key]
	m.mu.RUnlock()
	return
}

func (m *TStrStrMap) Get(key string) string {
	m.mu.RLock()
	val, _ := m.data[key]
	m.mu.RUnlock()
	return val
}

func (m *TStrStrMap) doSetWithLockCheck(key string, value interface{}) string {
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

func (m *TStrStrMap) GetOrSet(key string, value interface{}) string {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, value)
	} else {
		return v
	}
}

func (m *TStrStrMap) GetOrSetFunc(key string, fn GetStrFunc) string {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, fn())
	} else {
		return v
	}
}

func (m *TStrStrMap) GetOrSetFuncLock(key string, fn GetStrFunc) string {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, fn)
	} else {
		return v
	}
}

func (m *TStrStrMap) SetIfNotExist(key string, value interface{}) bool {
	if !m.Contains(key) {
		m.doSetWithLockCheck(key, value)
		return true
	}
	return false
}

func (m *TStrStrMap) SetIfNotExistFunc(key string, fn GetStrFunc) bool {
	if !m.Contains(key) {
		m.doSetWithLockCheck(key, fn())
		return true
	}
	return false
}

func (m *TStrStrMap) SetIfNotExistFuncLock(key string, fn GetStrFunc) bool {
	if !m.Contains(key) {
		m.doSetWithLockCheck(key, fn)
		return true
	}
	return false
}

func (m *TStrStrMap) Remove(key string) string {
	m.mu.Lock()
	val, exists := m.data[key]
	if exists {
		delete(m.data, key)
	}
	m.mu.Unlock()
	return val
}

func (m *TStrStrMap) Removes(keys []string) {
	m.mu.Lock()
	for _, key := range keys {
		delete(m.data, key)
	}
	m.mu.Unlock()
}

func (m *TStrStrMap) Keys() []string {
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

func (m *TStrStrMap) Values() []string {
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

func (m *TStrStrMap) Contains(key string) bool {
	m.mu.RLock()
	_, exists := m.data[key]
	m.mu.RUnlock()
	return exists
}

func (m *TStrStrMap) Size() int {
	m.mu.RLock()
	length := len(m.data)
	m.mu.RUnlock()
	return length
}

func (m *TStrStrMap) IsEmpty() bool {
	m.mu.RLock()
	empty := len(m.data) == 0
	m.mu.RUnlock()
	return empty
}

func (m *TStrStrMap) Clear() {
	m.mu.Lock()
	m.data = make(map[string]string)
	m.mu.Unlock()
}

func (m *TStrStrMap) LockFunc(fn StrStrMapLockFunc) {
	m.mu.Lock()
	defer m.mu.Unlock()
	fn(m.data)
}

func (m *TStrStrMap) RLockFunc(fn StrStrMapLockFunc) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	fn(m.data)
}

func (m *TStrStrMap) Flip() {
	m.mu.Lock()
	defer m.mu.Unlock()
	n := make(map[string]string, len(m.data))
	for k, v := range m.data {
		n[t.String(v)] = k
	}
	m.data = n
}

func (m *TStrStrMap) Merge(other *TStrStrMap) {
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

func (m *TStrStrMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Map())
}

func (m *TStrStrMap) String() string {
	b, err := m.MarshalJSON()
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func (m *TStrStrMap) GetVar(key string) *TVar {
	return NewVar(m.Get(key), true)
}

func (m *TStrStrMap) GetVarOrSet(key string, value interface{}) *TVar {
	return NewVar(m.GetOrSet(key, value), true)
}

func (m *TStrStrMap) GetVarOrSetFunc(key string, fn GetStrFunc) *TVar {
	return NewVar(m.GetOrSetFunc(key, fn), true)
}

func (m *TStrStrMap) GetVarOrSetFuncLock(key string, fn GetStrFunc) *TVar {
	return NewVar(m.GetOrSetFuncLock(key, fn), true)
}
