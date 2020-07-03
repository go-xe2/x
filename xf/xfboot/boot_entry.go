package xfboot

import (
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/sync/xsafeMap"
	"sync"
)

type BootEntry interface {
	// 启动入口名称
	EntryName() string
	// 启动入口实例指针,可以通过该指针转成实际的入口类型
	Entry() interface{}
}

type bootEntryItem struct {
	entry BootEntry
	// 是否已经初始化
	isInit bool
	// 入口绑定的初始化过程列表
	methods []func(entry BootEntry)
	look    sync.RWMutex
}

var entryInstances = xsafeMap.NewStrAnyMap()

// 注册启动初始化入口
func RegisterEntry(entry BootEntry) bool {
	if entryInstances.Contains(entry.EntryName()) {
		return false
	}
	item := &bootEntryItem{
		entry:   entry,
		isInit:  false,
		methods: make([]func(entry BootEntry), 0),
	}
	entryInstances.Set(entry.EntryName(), item)
	return true
}

func GetEntryOrRegister(name string, create func() BootEntry) interface{} {
	if v := entryInstances.Get(name); v != nil {
		return v.(*bootEntryItem).entry
	}
	item := &bootEntryItem{
		entry:   create(),
		isInit:  false,
		methods: make([]func(entry BootEntry), 0),
	}
	entryInstances.Set(name, item)
	return item.entry
}

// 注册启动初始化节点
func RegisterEntryPoint(entryName string, point func(entry BootEntry)) bool {
	v := entryInstances.Get(entryName)
	if v == nil {
		panic(exception.Newf("启动初始化结点%s未注册", entryName))
	}
	entryItem := v.(*bootEntryItem)
	entryItem.methods = append(entryItem.methods, point)
	return true
}

// 检查是否注册了指定入口
func IsRegisterEntry(entryName string) bool {
	return entryInstances.Contains(entryName)
}

// 检查指定入口是否已经初始化
func IsEntryInit(entryName string) bool {
	if v := entryInstances.Get(entryName); v != nil {
		item := v.(*bootEntryItem)
		item.look.RLock()
		defer item.look.RUnlock()
		return v.(*bootEntryItem).isInit
	}
	// 不存在的入口，不需要调用，所以都是true
	return true
}

// 初始化指定入口
func InitEntry(entryName string) {
	if IsEntryInit(entryName) {
		return
	}
	v := entryInstances.Get(entryName)
	if v == nil {
		return
	}
	entryItem := v.(*bootEntryItem)
	entryItem.look.Lock()
	for _, method := range entryItem.methods {
		(func() {
			defer func() {
				if e := recover(); e != nil {
					panic(exception.Newf("初始化入口%s出错:%v", entryItem.entry.EntryName(), e))
				}
			}()
			method(entryItem.entry)
		})()
	}
	entryItem.isInit = true
	entryItem.look.Unlock()
}

// 初始化所有入口
func InitAllEntry() {
	entries := entryInstances.Keys()
	for _, entryName := range entries {
		InitEntry(entryName)
	}
}
