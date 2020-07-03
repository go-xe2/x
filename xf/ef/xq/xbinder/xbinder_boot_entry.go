package xbinder

import (
	"github.com/go-xe2/x/type/xstring"
	"github.com/go-xe2/x/xf/ef/xqi"
	"github.com/go-xe2/x/xf/xfboot"
	"strings"
)

type binderEntryItem struct {
	// 绑定器
	binder xqi.DbQueryBinder
	// 默认参数
	options map[string]interface{}
}

type BinderBootEntry interface {
	Register(binderName string, binder xqi.DbQueryBinder, options ...map[string]interface{})
	GetBinder(binderName xqi.TBinderName) xqi.DbQueryBinder
	GetBinderByName(nameOptions string) xqi.DbQueryBinder
}

type xBinderBootEntry struct {
	binders map[string]*binderEntryItem
}

var _ xfboot.BootEntry = (*xBinderBootEntry)(nil)

const BinderBootEntryName = "binderEntry"

var defaultBinder xqi.DbQueryBinder

// 注册启动初始化入口
var BootEntry = xfboot.GetEntryOrRegister(BinderBootEntryName, func() xfboot.BootEntry {
	inst := &xBinderBootEntry{
		binders: make(map[string]*binderEntryItem),
	}
	registerInternalBinder(inst)
	return inst
}).(BinderBootEntry)

func (be *xBinderBootEntry) EntryName() string {
	return BinderBootEntryName
}

func (be *xBinderBootEntry) Entry() interface{} {
	return be
}

func (be *xBinderBootEntry) Register(binderName string, binder xqi.DbQueryBinder, options ...map[string]interface{}) {
	var ops map[string]interface{}
	if len(options) > 0 {
		ops = options[0]
	} else {
		ops = make(map[string]interface{})
	}
	be.binders[binderName] = &binderEntryItem{
		binder:  binder,
		options: ops,
	}
}

func (be *xBinderBootEntry) IsInit() bool {
	return xfboot.IsEntryInit(be.EntryName())
}

func (be *xBinderBootEntry) Init() {
	if be.IsInit() {
		return
	}
	xfboot.InitEntry(be.EntryName())
}

func (be *xBinderBootEntry) GetBinder(binderName xqi.TBinderName) xqi.DbQueryBinder {
	be.Init()
	if item, ok := be.binders[binderName.Name()]; !ok {
		return defaultBinder
	} else {
		opts := make(map[string]interface{})
		// 默认参数
		for k, v := range item.options {
			opts[k] = v
		}
		for k, v := range binderName.Options() {
			opts[k] = v
		}
		return item.binder.NewInstance(opts)
	}
}

func (be *xBinderBootEntry) GetBinderByName(nameOptions string) xqi.DbQueryBinder {
	be.Init()
	szName := strings.Split(nameOptions, ":")
	name := szName[0]
	if item, ok := be.binders[name]; !ok {
		return defaultBinder
	} else {
		options := make(map[string]interface{})
		s := ""
		if len(szName) > 0 {
			s = szName[1]
		}
		for k, v := range item.options {
			options[k] = v
		}
		inOpts := xstring.ParseKeyValue(s)
		for k, v := range inOpts {
			options[k] = v
		}
		return item.binder.NewInstance(options)
	}
}
