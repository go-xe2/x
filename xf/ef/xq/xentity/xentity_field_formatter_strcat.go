package xentity

import (
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/type/xstring"
)

type FieldStrCatFormatter struct {
	options map[string]interface{}
}

var _ FieldFormatter = (*FieldStrCatFormatter)(nil)

const FormatStrCat = "strCat"

// tag字符串为: strCat[:from,len]
var defaultStrCatFormatOptions = map[string]interface{}{
	"from": 0,
	"len":  -1,
}

func NewFieldStrCatFormatter(options ...map[string]interface{}) FieldFormatter {
	inst := new(FieldStrCatFormatter)
	if len(options) > 0 && options[0] != nil {
		inst.options = options[0]
	} else {
		inst.options = defaultDateFormatOptions
	}
	return inst
}

// 格式化参数
func (fdf *FieldStrCatFormatter) Options() map[string]interface{} {
	return fdf.options
}

func (fdf *FieldStrCatFormatter) OptionFromSlice(items ...interface{}) map[string]interface{} {
	opts := make(map[string]interface{})
	for k, v := range fdf.options {
		opts[k] = v
	}
	if len(items) > 0 {
		opts["from"] = t.Int(items[0])
	}
	if len(items) > 1 {
		opts["len"] = t.Int(items[1], -1)
	}
	return opts
}

func (fdf *FieldStrCatFormatter) Formatter() func(old interface{}, options ...map[string]interface{}) interface{} {
	return func(old interface{}, options ...map[string]interface{}) interface{} {
		opts := defaultDateFormatOptions
		if len(options) > 0 {
			opts = options[0]
		}
		s := t.String(old)
		nFrom := t.Int(opts["from"])
		nLen := t.Int(opts["len"], -1)
		if nLen == -1 {
			nLen = len([]rune(s))
		}
		return xstring.SubStr(s, nFrom, nLen)
	}
}
