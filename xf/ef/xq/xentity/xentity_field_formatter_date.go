package xentity

import (
	"github.com/go-xe2/x/type/t"
)

type FieldDateFormatter struct {
	options map[string]interface{}
}

var _ FieldFormatter = (*FieldDateFormatter)(nil)

const FormatDate = "dateFormat"

// tag字符串为: dateFormat[:y-M-d h:i:s]
var defaultDateFormatOptions = map[string]interface{}{
	"format": "Y-m-d H:i:s",
}

func NewFieldDateFormatter(options ...map[string]interface{}) FieldFormatter {
	inst := new(FieldDateFormatter)
	if len(options) > 0 && options[0] != nil {
		inst.options = options[0]
	} else {
		inst.options = defaultDateFormatOptions
	}
	return inst
}

// 日期格式化
func (fdf *FieldDateFormatter) Options() map[string]interface{} {
	return fdf.options
}

func (fdf *FieldDateFormatter) OptionFromSlice(items ...interface{}) map[string]interface{} {
	opts := make(map[string]interface{})
	for k, v := range fdf.options {
		opts[k] = v
	}
	if len(items) > 0 {
		opts["format"] = t.String(items[0], "Y-m-d H:i:s")
	}
	return opts
}

func (fdf *FieldDateFormatter) Formatter() func(old interface{}, options ...map[string]interface{}) interface{} {
	return func(old interface{}, options ...map[string]interface{}) interface{} {
		opts := defaultDateFormatOptions
		if len(options) > 0 {
			opts = options[0]
		}
		szFormat := t.String(opts["format"])
		tr := t.XTime(old)
		return tr.Format(szFormat)
	}
}
