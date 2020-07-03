package xentity

import (
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/xf/xfboot"
)

type FieldFormatter interface {
	Options() map[string]interface{}
	Formatter() func(old interface{}, options ...map[string]interface{}) interface{}
	OptionFromSlice(items ...interface{}) map[string]interface{}
}

type FieldFormatterEntry interface {
	Register(name string, formatter FieldFormatter)
	Formatter(formatName string, old interface{}, options ...map[string]interface{}) interface{}
	GetFormatter(name string) FieldFormatter
}

type fieldFormatterFactoryEntry struct {
	formatters map[string]FieldFormatter
}

var _ FieldFormatterEntry = (*fieldFormatterFactoryEntry)(nil)

var _ xfboot.BootEntry = (*fieldFormatterFactoryEntry)(nil)

const fieldFormatterBootEntry = "fieldFormatBootEntry"

func (ffe *fieldFormatterFactoryEntry) EntryName() string {
	return fieldFormatterBootEntry
}

func (ffe *fieldFormatterFactoryEntry) Entry() interface{} {
	return ffe
}

func (ffe *fieldFormatterFactoryEntry) Register(name string, formatter FieldFormatter) {
	if _, ok := ffe.formatters[name]; ok {
		panic(exception.Newf("实体字段格式化器%s已经存在", name))
	}
	ffe.formatters[name] = formatter
}

func (ffe *fieldFormatterFactoryEntry) IsInit() bool {
	return xfboot.IsEntryInit(ffe.EntryName())
}

func (ffe *fieldFormatterFactoryEntry) Init() {
	if ffe.IsInit() {
		return
	}
	xfboot.InitEntry(ffe.EntryName())
}

func (ffe *fieldFormatterFactoryEntry) GetFormatter(name string) FieldFormatter {
	if f, ok := ffe.formatters[name]; ok {
		return f
	}
	return nil
}

func (ffe *fieldFormatterFactoryEntry) Formatter(formatName string, old interface{}, options ...map[string]interface{}) interface{} {
	if fn := ffe.GetFormatter(formatName); fn != nil {
		opts := fn.Options()
		if opts == nil {
			opts = make(map[string]interface{})
		}
		if len(options) > 0 && options[0] != nil {
			option := options[0]
			for k, v := range option {
				opts[k] = v
			}
		}
		return fn.Formatter()(old, opts)
	}
	return old
}

var fieldFormatters = xfboot.GetEntryOrRegister(fieldFormatterBootEntry, func() xfboot.BootEntry {
	inst := &fieldFormatterFactoryEntry{
		formatters: make(map[string]FieldFormatter),
	}
	registerFieldFormatters(inst)
	return inst
}).(FieldFormatterEntry)

// 注册字段格式化方法
func RegisterFormatter(name string, formatter FieldFormatter) {
	fieldFormatters.Register(name, formatter)
}

func Formatter(formatName string, old interface{}, options ...map[string]interface{}) interface{} {
	return fieldFormatters.Formatter(formatName, old)
}

func GetFieldFormatter(formatName string) FieldFormatter {
	return fieldFormatters.GetFormatter(formatName)
}
