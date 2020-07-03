package xfileNotify

import (
	"errors"
	"fmt"
	"github.com/go-xe2/x/container/xstackQe"
	_type "github.com/go-xe2/x/sync/type"
	"github.com/go-xe2/x/sync/xsafeMap"
)

// 注册的监听回调方法
type Callback struct {
	Id        int                     // 唯一ID
	Func      func(event *TEvent)     // 回调方法
	Path      string                  // 监听的文件/目录
	elem      *xstackQe.TStackElement // 指向回调函数链表中的元素项位置(便于删除)
	recursive bool                    // 当目录时，是否递归监听(使用在子文件/目录回溯查找回调函数时)
}

// 按位进行识别的操作集合
type Op uint32

// 必须放到一个const分组里面
const (
	CREATE Op = 1 << iota
	WRITE
	REMOVE
	RENAME
	CHMOD
)

const (
	REPEAT_EVENT_FILTER_INTERVAL = 1      // (毫秒)重复事件过滤间隔
	mFSNOTIFY_EVENT_EXIT         = "exit" // 是否退出回调执行
)

var (
	// 默认的Watcher对象
	defaultWatcher, _ = New()
	// 默认的watchers是否初始化，使用时才创建
	watcherInited = _type.NewBool()
	// 回调方法ID与对象指针的映射哈希表，用于根据ID快速查找回调对象
	callbackIdMap = xsafeMap.NewIntAnyMap()
	// 回调函数的ID生成器(原子操作)
	callbackIdGenerator = _type.NewInt()
)

// 添加对指定文件/目录的监听，并给定回调函数；如果给定的是一个目录，默认递归监控。
func Add(path string, callbackFunc func(event *TEvent), recursive ...bool) (callback *Callback, err error) {
	return defaultWatcher.Add(path, callbackFunc, recursive...)
}

// 递归移除对指定文件/目录的所有监听回调
func Remove(path string) error {
	return defaultWatcher.Remove(path)
}

// 根据指定的回调函数ID，移出指定的inotify回调函数
func RemoveCallback(callbackId int) error {
	callback := (*Callback)(nil)
	if r := callbackIdMap.Get(callbackId); r != nil {
		callback = r.(*Callback)
	}
	if callback == nil {
		return errors.New(fmt.Sprintf(`callback for id %d not found`, callbackId))
	}
	defaultWatcher.RemoveCallback(callbackId)
	return nil
}

// 在回调方法中调用该方法退出回调注册
func Exit() {
	panic(mFSNOTIFY_EVENT_EXIT)
}
