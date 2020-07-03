package structs

import (
	"encoding/json"
	"fmt"
	"github.com/go-xe2/x/type/xstring"
)

type TagValue string

type FieldTag struct {
	items map[string]TagValue `default:"" name:""`
}

func NewFieldTag(szTag ...string) *FieldTag {
	inst := &FieldTag{}
	if len(szTag) > 0 {
		return inst.Parse(szTag[0])
	}
	inst.items = make(map[string]TagValue)
	return inst
}

func (ft *FieldTag) Parse(szTag string) *FieldTag {
	nLen := len(szTag)
	var layMod = 0 // 0 查找key, 1: 值开始, 2: 找值结束
	var i = 0
	var nStart = 0
	var key = ""
	var value = ""
	ft.items = make(map[string]TagValue)
	for i < nLen {
		c := szTag[i]
		if c == ' ' && layMod <= 1 {
			nStart += 1
			i++
			continue
		} else if c == ':' {
			key = szTag[nStart:i]
			layMod = 1
			nStart = i + 1
		} else if c == '"' {
			if layMod == 1 {
				nStart = i + 1
				layMod = 2
			} else if layMod == 2 {
				value = szTag[nStart:i]
				nStart = i + 1
				layMod = 0
				ft.items[key] = TagValue(value)
			}
		}
		i++
	}
	return ft
}

func (ft *FieldTag) Items() map[string]TagValue {
	return ft.items
}

func (ft *FieldTag) String() string {
	result := make([]string, 0)
	for k, v := range ft.items {
		result = append(result, fmt.Sprintf("%s:\"%s\"", k, v))
	}
	return xstring.Join(result, " ")
}

func (ft *FieldTag) HasKey(key string) bool {
	if _, ok := ft.items[key]; ok {
		return true
	}
	return false
}

func (ft *FieldTag) Get(key string) TagValue {
	if s, ok := ft.items[key]; ok {
		return s
	}
	return ""
}

/*
TagValue methods
*/

func (tv TagValue) Values() []string {
	return xstring.Split(string(tv), ",")
}

func (tv TagValue) Value() string {
	cols := tv.Values()
	if len(cols) > 0 {
		return cols[0]
	}
	return ""
}

func (tv TagValue) IsEmpty() bool {
	return tv == ""
}

func (tv TagValue) String() string {
	return string(tv)
}

func (tv TagValue) Raw() string {
	return string(tv)
}

func (tv TagValue) Map() map[string]string {
	result := make(map[string]string)
	if tv.IsEmpty() {
		return result
	}
	err := json.Unmarshal([]byte(tv), &result)
	if err != nil {
		result = xstring.ParseKeyValue(string(tv))
	}
	return result
}
