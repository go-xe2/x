package xjson

import (
	"github.com/go-xe2/x/core/rwmutex"
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/type/xstring"
	"reflect"
	"strconv"
	"strings"
)

const (
	mDEFAULT_SPLIT_CHAR = '.'
	// json使用引号
	mJSON_VALUE_QUOTES = "\""
)

type TJson struct {
	mu *rwmutex.RWMutex
	p  *interface{} // Pointer for hierarchical data access, it's the root of data in default.
	c  byte         // 子级访问连接符，默认"."
	vc bool         // 是否允许通过带"."的键名访问子层，默认false
}

func (j *TJson) MarshalJSON() ([]byte, error) {
	return j.ToJson()
}

func (j *TJson) setValue(pattern string, value interface{}, removed bool) error {
	array := strings.Split(pattern, string(j.c))
	length := len(array)
	value = j.convertValue(value)
	// 初始化判断
	if *j.p == nil {
		if xstring.IsNumeric(array[0]) {
			*j.p = make([]interface{}, 0)
		} else {
			*j.p = make(map[string]interface{})
		}
	}
	var pparent *interface{} = nil // Parent pointer.
	var pointer *interface{} = j.p // Current pointer.
	j.mu.Lock()
	defer j.mu.Unlock()
	for i := 0; i < length; i++ {
		switch (*pointer).(type) {
		case map[string]interface{}:
			if i == length-1 {
				if removed && value == nil {
					// Delete item from map.
					delete((*pointer).(map[string]interface{}), array[i])
				} else {
					(*pointer).(map[string]interface{})[array[i]] = value
				}
			} else {
				// If the key does not exit in the map.
				if v, ok := (*pointer).(map[string]interface{})[array[i]]; !ok {
					if removed && value == nil {
						goto done
					}
					// Creating new node.
					if xstring.IsNumeric(array[i+1]) {
						// Creating array node.
						n, _ := strconv.Atoi(array[i+1])
						var v interface{} = make([]interface{}, n+1)
						pparent = j.setPointerWithValue(pointer, array[i], v)
						pointer = &v
					} else {
						// Creating map node.
						var v interface{} = make(map[string]interface{})
						pparent = j.setPointerWithValue(pointer, array[i], v)
						pointer = &v
					}
				} else {
					pparent = pointer
					pointer = &v
				}
			}

		case []interface{}:
			if !xstring.IsNumeric(array[i]) {
				if i == length-1 {
					*pointer = map[string]interface{}{array[i]: value}
				} else {
					var v interface{} = make(map[string]interface{})
					*pointer = v
					pparent = pointer
					pointer = &v
				}
				continue
			}

			valn, err := strconv.Atoi(array[i])
			if err != nil {
				return err
			}
			// Leaf node.
			if i == length-1 {
				if len((*pointer).([]interface{})) > valn {
					if removed && value == nil {
						// Deleting element.
						if pparent == nil {
							*pointer = append((*pointer).([]interface{})[:valn], (*pointer).([]interface{})[valn+1:]...)
						} else {
							j.setPointerWithValue(pparent, array[i-1], append((*pointer).([]interface{})[:valn], (*pointer).([]interface{})[valn+1:]...))
						}
					} else {
						(*pointer).([]interface{})[valn] = value
					}
				} else {
					if removed && value == nil {
						goto done
					}
					if pparent == nil {
						// It is the root node.
						j.setPointerWithValue(pointer, array[i], value)
					} else {
						// It is not the root node.
						s := make([]interface{}, valn+1)
						copy(s, (*pointer).([]interface{}))
						s[valn] = value
						j.setPointerWithValue(pparent, array[i-1], s)
					}
				}
			} else {
				if xstring.IsNumeric(array[i+1]) {
					n, _ := strconv.Atoi(array[i+1])
					if len((*pointer).([]interface{})) > valn {
						(*pointer).([]interface{})[valn] = make([]interface{}, n+1)
						pparent = pointer
						pointer = &(*pointer).([]interface{})[valn]
					} else {
						if removed && value == nil {
							goto done
						}
						var v interface{} = make([]interface{}, n+1)
						pparent = j.setPointerWithValue(pointer, array[i], v)
						pointer = &v
					}
				} else {
					var v interface{} = make(map[string]interface{})
					pparent = j.setPointerWithValue(pointer, array[i], v)
					pointer = &v
				}
			}

		// If the variable pointed to by the <pointer> is not of a reference type,
		// then it modifies the variable via its the parent, ie: pparent.
		default:
			if removed && value == nil {
				goto done
			}
			if xstring.IsNumeric(array[i]) {
				n, _ := strconv.Atoi(array[i])
				s := make([]interface{}, n+1)
				if i == length-1 {
					s[n] = value
				}
				if pparent != nil {
					pparent = j.setPointerWithValue(pparent, array[i-1], s)
				} else {
					*pointer = s
					pparent = pointer
				}
			} else {
				var v interface{} = make(map[string]interface{})
				if i == length-1 {
					v = map[string]interface{}{
						array[i]: value,
					}
				}
				if pparent != nil {
					pparent = j.setPointerWithValue(pparent, array[i-1], v)
				} else {
					*pointer = v
					pparent = pointer
				}
				pointer = &v
			}
		}
	}
done:
	return nil
}

func (j *TJson) convertValue(value interface{}) interface{} {
	switch value.(type) {
	case map[string]interface{}:
		return value
	case []interface{}:
		return value
	default:
		rv := reflect.ValueOf(value)
		kind := rv.Kind()
		if kind == reflect.Ptr {
			rv = rv.Elem()
			kind = rv.Kind()
		}
		switch kind {
		case reflect.Array:
			return t.Interfaces(value)
		case reflect.Slice:
			return t.Interfaces(value)
		case reflect.Map:
			return t.Map(value)
		case reflect.Struct:
			return t.Map(value)
		default:
			// Use json decode/encode at last.
			b, _ := Encode(value)
			v, _ := Decode(b)
			return v
		}
	}
}

func (j *TJson) setPointerWithValue(pointer *interface{}, key string, value interface{}) *interface{} {
	switch (*pointer).(type) {
	case map[string]interface{}:
		(*pointer).(map[string]interface{})[key] = value
		return &value
	case []interface{}:
		n, _ := strconv.Atoi(key)
		if len((*pointer).([]interface{})) > n {
			(*pointer).([]interface{})[n] = value
			return &(*pointer).([]interface{})[n]
		} else {
			s := make([]interface{}, n+1)
			copy(s, (*pointer).([]interface{}))
			s[n] = value
			*pointer = s
			return &s[n]
		}
	default:
		*pointer = value
	}
	return pointer
}

func (j *TJson) getPointerByPattern(pattern string) *interface{} {
	if j.vc {
		return j.getPointerByPatternWithViolenceCheck(pattern)
	} else {
		return j.getPointerByPatternWithoutViolenceCheck(pattern)
	}
}

func (j *TJson) getPointerByPatternWithViolenceCheck(pattern string) *interface{} {
	if !j.vc {
		return j.getPointerByPatternWithoutViolenceCheck(pattern)
	}
	index := len(pattern)
	start := 0
	length := 0
	pointer := j.p
	if index == 0 {
		return pointer
	}
	for {
		if r := j.checkPatternByPointer(pattern[start:index], pointer); r != nil {
			length += index - start
			if start > 0 {
				length += 1
			}
			start = index + 1
			index = len(pattern)
			if length == len(pattern) {
				return r
			} else {
				pointer = r
			}
		} else {
			// Get the position for next separator char.
			index = strings.LastIndexByte(pattern[start:index], j.c)
			if index != -1 && length > 0 {
				index += length + 1
			}
		}
		if start >= index {
			break
		}
	}
	return nil
}

func (j *TJson) getPointerByPatternWithoutViolenceCheck(pattern string) *interface{} {
	if j.vc {
		return j.getPointerByPatternWithViolenceCheck(pattern)
	}
	pointer := j.p
	if len(pattern) == 0 {
		return pointer
	}
	array := strings.Split(pattern, string(j.c))
	for k, v := range array {
		if r := j.checkPatternByPointer(v, pointer); r != nil {
			if k == len(array)-1 {
				return r
			} else {
				pointer = r
			}
		} else {
			break
		}
	}
	return nil
}

func (j *TJson) checkPatternByPointer(key string, pointer *interface{}) *interface{} {
	switch (*pointer).(type) {
	case map[string]interface{}:
		if v, ok := (*pointer).(map[string]interface{})[key]; ok {
			return &v
		}
	case []interface{}:
		if xstring.IsNumeric(key) {
			n, err := strconv.Atoi(key)
			if err == nil && len((*pointer).([]interface{})) > n {
				return &(*pointer).([]interface{})[n]
			}
		}
	}
	return nil
}
