package xjson

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-xe2/x/core/rwmutex"
	"github.com/go-xe2/x/encoding/xtoml"
	"github.com/go-xe2/x/encoding/xxml"
	"github.com/go-xe2/x/encoding/xyaml"
	"github.com/go-xe2/x/os/xfile"
	"github.com/go-xe2/x/os/xfileCache"
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/utils/xregex"
	"reflect"
)

func New(data interface{}, unsafe ...bool) *TJson {
	j := (*TJson)(nil)
	switch data.(type) {
	case string, []byte:
		if r, err := LoadContent(t.Bytes(data)); err == nil {
			j = r
		} else {
			j = &TJson{
				p:  &data,
				c:  byte(mDEFAULT_SPLIT_CHAR),
				vc: false,
			}
		}
	default:
		rv := reflect.ValueOf(data)
		kind := rv.Kind()
		if kind == reflect.Ptr {
			rv = rv.Elem()
			kind = rv.Kind()
		}
		switch kind {
		case reflect.Slice:
			fallthrough
		case reflect.Array:
			i := interface{}(nil)
			i = t.Interfaces(data)
			j = &TJson{
				p:  &i,
				c:  byte(mDEFAULT_SPLIT_CHAR),
				vc: false,
			}
		case reflect.Map:
			fallthrough
		case reflect.Struct:
			i := interface{}(nil)
			i = t.Map(data, "json")
			j = &TJson{
				p:  &i,
				c:  byte(mDEFAULT_SPLIT_CHAR),
				vc: false,
			}
		default:
			j = &TJson{
				p:  &data,
				c:  byte(mDEFAULT_SPLIT_CHAR),
				vc: false,
			}
		}
	}
	j.mu = rwmutex.New(unsafe...)
	return j
}

func NewUnsafe(data ...interface{}) *TJson {
	if len(data) > 0 {
		return New(data[0], true)
	}
	return New(nil, true)
}

// 检查data是否是json数据类型
func Valid(data interface{}) bool {
	return json.Valid(t.Bytes(data))
}

// 编码为json数据
func Encode(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}

// 解析json数据为interface{}
func Decode(data interface{}) (interface{}, error) {
	var value interface{}
	if err := DecodeTo(t.Bytes(data), &value); err != nil {
		return nil, err
	} else {
		return value, nil
	}
}

// 解析json数据为指定格式
func DecodeTo(data interface{}, v interface{}) error {
	decoder := json.NewDecoder(bytes.NewReader(t.Bytes(data)))
	decoder.UseNumber()
	return decoder.Decode(v)
}

// 解码json数据为TJson对象实例
func DecodeToJson(data interface{}, unsafe ...bool) (*TJson, error) {
	if v, err := Decode(t.Bytes(data)); err != nil {
		return nil, err
	} else {
		return New(v, unsafe...), nil
	}
}

// 文件加载json数据并转换为TJson对象实例
// 文件支持json,xml, toml,yaml格式
func Load(path string, unsafe ...bool) (*TJson, error) {
	return doLoadContent(xfile.Ext(path), xfileCache.GetBinContents(path), unsafe...)
}

// 从json字符串加载TJson对象实例
func LoadJson(data interface{}, unsafe ...bool) (*TJson, error) {
	return doLoadContent("json", t.Bytes(data), unsafe...)
}

// 从xml字符串加载TJson对象实例
func LoadXml(data interface{}, unsafe ...bool) (*TJson, error) {
	return doLoadContent("xml", t.Bytes(data), unsafe...)
}

// 从yaml字符串加载TJson对象实例
func LoadYaml(data interface{}, unsafe ...bool) (*TJson, error) {
	return doLoadContent("yaml", t.Bytes(data), unsafe...)
}

// 从toml字符串格式加载为TJson对角实例
func LoadToml(data interface{}, unsafe ...bool) (*TJson, error) {
	return doLoadContent("toml", t.Bytes(data), unsafe...)
}

// 加载解析字符串数据，格式支持json, xml, toml, yaml
func doLoadContent(dataType string, data []byte, unsafe ...bool) (*TJson, error) {
	var err error
	var result interface{}
	if len(data) == 0 {
		return New(nil, unsafe...), nil
	}
	if dataType == "" {
		dataType = checkDataType(data)
	}
	switch dataType {
	case "json", ".json":

	case "xml", ".xml":
		if data, err = xxml.ToJson(data); err != nil {
			return nil, err
		}

	case "yml", "yaml", ".yml", ".yaml":
		if data, err = xyaml.ToJson(data); err != nil {
			return nil, err
		}

	case "toml", ".toml":
		if data, err = xtoml.ToJson(data); err != nil {
			return nil, err
		}

	default:
		err = errors.New("unsupported type for loading")
	}
	if err != nil {
		return nil, err
	}
	if result == nil {
		decoder := json.NewDecoder(bytes.NewReader(data))
		decoder.UseNumber()
		if err := decoder.Decode(&result); err != nil {
			return nil, err
		}
		switch result.(type) {
		case string, []byte:
			return nil, fmt.Errorf(`json decoding failed for content: %s`, string(data))
		}
	}
	return New(result, unsafe...), nil
}

// 将其他数据加载转换成TJson对象实例
func LoadContent(data interface{}, unsafe ...bool) (*TJson, error) {
	content := t.Bytes(data)
	if len(content) == 0 {
		return New(nil, unsafe...), nil
	}
	return doLoadContent(checkDataType(content), content, unsafe...)

}

// 数据格式类型检查
func checkDataType(content []byte) string {
	if json.Valid(content) {
		return "json"
	} else if xregex.IsMatch(`^<.+>[\S\s]+<.+>$`, content) {
		return "xml"
	} else if xregex.IsMatch(`^[\s\t]*[\w\-]+\s*:\s*.+`, content) || xregex.IsMatch(`\n[\s\t]*[\w\-]+\s*:\s*.+`, content) {
		return "yml"
	} else if xregex.IsMatch(`^[\s\t]*[\w\-]+\s*=\s*.+`, content) || xregex.IsMatch(`\n[\s\t]*[\w\-]+\s*=\s*.+`, content) {
		return "toml"
	} else {
		return ""
	}
}
