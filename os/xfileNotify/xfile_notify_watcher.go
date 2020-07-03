package xfileNotify

import (
	"errors"
	"fmt"
	"github.com/go-xe2/third/github.com/fsnotify/fsnotify"
	"github.com/go-xe2/x/container/xqueue"
	"github.com/go-xe2/x/os/xcache"
	"github.com/go-xe2/x/sync/xsafeMap"
	"github.com/go-xe2/x/sync/xsafeStack"
)

// 监听管理对象
type TWatcher struct {
	watcher   *fsnotify.Watcher    // 底层fsnotify对象
	events    *xqueue.TQueue       // 过滤后的事件通知，不会出现重复事件
	cache     *xcache.TCache       // 缓存对象，主要用于事件重复过滤
	callbacks *xsafeMap.TStrAnyMap // 注册的所有绝对路径(文件/目录)及其对应的回调函数列表map
	closeChan chan struct{}        // 关闭事件
}

// 创建监听管理对象，主要注意的是创建监听对象会占用系统的inotify句柄数量，受到 fs.inotify.max_user_instances 的限制
func New() (*TWatcher, error) {
	w := &TWatcher{
		cache:     xcache.New(),
		events:    xqueue.New(),
		closeChan: make(chan struct{}),
		callbacks: xsafeMap.NewStrAnyMap(),
	}
	if watcher, err := fsnotify.NewWatcher(); err == nil {
		w.watcher = watcher
	} else {
		return nil, err
	}
	w.startWatchLoop()
	w.startEventLoop()
	return w, nil
}

// 添加监控，path参数支持文件或者目录路径，recursive为非必需参数，默认为递归监控(当path为目录时)。
// 如果添加目录，这里只会返回目录的callback，按照callback删除时会递归删除。
func (w *TWatcher) Add(path string, callbackFunc func(event *TEvent), recursive ...bool) (callback *Callback, err error) {
	// 首先添加这个文件/目录
	callback, err = w.addWithCallbackFunc(path, callbackFunc, recursive...)
	if err != nil {
		return nil, err
	}
	// 如果需要递归，那么递归添加其下的子级目录，
	// 注意!!
	// 1、这里只递归添加**目录**, 而非文件，因为监控了目录即监控了其下一级的文件;
	// 2、这里只是添加底层监控对象对**子级所有目录**的监控，没有任何回调函数的设置，在事件产生时会回溯查找父级的回调函数；
	if fileIsDir(path) && (len(recursive) == 0 || recursive[0]) {
		for _, subPath := range fileAllDirs(path) {
			if fileIsDir(subPath) {
				w.watcher.Add(subPath)
			}
		}
	}
	return
}

// 添加对指定文件/目录的监听，并给定回调函数
func (w *TWatcher) addWithCallbackFunc(path string, callbackFunc func(event *TEvent), recursive ...bool) (callback *Callback, err error) {
	// 这里统一转换为当前系统的绝对路径，便于统一监控文件名称
	if t := fileRealPath(path); t == "" {
		return nil, errors.New(fmt.Sprintf(`"%s" does not exist`, path))
	} else {
		path = t
	}
	callback = &Callback{
		Id:        callbackIdGenerator.Add(1),
		Func:      callbackFunc,
		Path:      path,
		recursive: true,
	}
	if len(recursive) > 0 {
		callback.recursive = recursive[0]
	}
	// 注册回调函数
	w.callbacks.LockFunc(func(m map[string]interface{}) {
		list := (*xsafeStack.TSafeStackQe)(nil)
		if v, ok := m[path]; !ok {
			list = xsafeStack.New()
			m[path] = list
		} else {
			list = v.(*xsafeStack.TSafeStackQe)
		}
		callback.elem = list.PushBack(callback)
	})
	// 添加底层监听
	w.watcher.Add(path)
	// 添加成功后会注册该callback id到全局的哈希表
	callbackIdMap.Set(callback.Id, callback)
	return
}

// 关闭监听管理对象
func (w *TWatcher) Close() {
	w.events.Close()
	w.watcher.Close()
	close(w.closeChan)
}

// 递归移除对指定文件/目录的所有监听回调
func (w *TWatcher) Remove(path string) error {
	// 首先移除path注册的回调注册，以及callbackIdMap中的ID
	if r := w.callbacks.Remove(path); r != nil {
		list := r.(*xsafeStack.TSafeStackQe)
		for {
			if r := list.PopFront(); r != nil {
				callbackIdMap.Remove(r.(*Callback).Id)
			} else {
				break
			}
		}
	}
	// 其次递归判断所有的子级是否可删除监听
	if subPaths, err := fileScanDir(path, "*", true); err == nil && len(subPaths) > 0 {
		for _, subPath := range subPaths {
			if w.checkPathCanBeRemoved(subPath) {
				w.watcher.Remove(subPath)
			}
		}
	}
	// 最后移除底层的监听
	return w.watcher.Remove(path)
}

// 判断给定的路径是否可以删除监听(只有所有回调函数都没有了才能删除)
func (w *TWatcher) checkPathCanBeRemoved(path string) bool {
	// 首先检索path对应的回调函数
	if v := w.callbacks.Get(path); v != nil {
		return false
	}
	// 其次查找父级目录有无回调注册
	dirPath := fileDir(path)
	if v := w.callbacks.Get(dirPath); v != nil {
		return false
	}
	// 最后回溯查找递归回调函数
	for {
		parentDirPath := fileDir(dirPath)
		if parentDirPath == dirPath {
			break
		}
		if v := w.callbacks.Get(parentDirPath); v != nil {
			return false
		}
		dirPath = parentDirPath
	}
	return true
}

// 根据指定的回调函数ID，移出指定的inotify回调函数
func (w *TWatcher) RemoveCallback(callbackId int) {
	callback := (*Callback)(nil)
	if r := callbackIdMap.Get(callbackId); r != nil {
		callback = r.(*Callback)
	}
	if callback != nil {
		if r := w.callbacks.Get(callback.Path); r != nil {
			r.(*xsafeStack.TSafeStackQe).Remove(callback.elem)
		}
		callbackIdMap.Remove(callbackId)
	}
}

// 监听循环
func (w *TWatcher) startWatchLoop() {
	go func() {
		for {
			select {
			// 关闭事件
			case <-w.closeChan:
				return

			// 监听事件
			case ev := <-w.watcher.Events:
				//fmt.Println("ev:", ev.String())
				w.cache.SetIfNotExist(ev.String(), func() interface{} {
					w.events.Push(&TEvent{
						event:   ev,
						Path:    ev.Name,
						Op:      Op(ev.Op),
						Watcher: w,
					})
					return struct{}{}
				}, REPEAT_EVENT_FILTER_INTERVAL)

			case <-w.watcher.Errors:
				//fmt.Fprintf(os.Stderr, "[gfsnotify] error: %s\n", err.Error())
			}
		}
	}()
}

// 获得文件路径的监听回调，包括层级的监听回调。
func (w *TWatcher) getCallbacks(path string) (callbacks []*Callback) {
	// 首先检索path对应的回调函数
	if v := w.callbacks.Get(path); v != nil {
		for _, v := range v.(*xsafeStack.TSafeStackQe).FrontAll() {
			callback := v.(*Callback)
			callbacks = append(callbacks, callback)
		}
	}
	// 其次查找父级目录有无回调注册
	dirPath := fileDir(path)
	if v := w.callbacks.Get(dirPath); v != nil {
		for _, v := range v.(*xsafeStack.TSafeStackQe).FrontAll() {
			callback := v.(*Callback)
			callbacks = append(callbacks, callback)
		}
	}
	// 最后回溯查找递归回调函数
	for {
		parentDirPath := fileDir(dirPath)
		if parentDirPath == dirPath {
			break
		}
		if v := w.callbacks.Get(parentDirPath); v != nil {
			for _, v := range v.(*xsafeStack.TSafeStackQe).FrontAll() {
				callback := v.(*Callback)
				if callback.recursive {
					callbacks = append(callbacks, callback)
				}
			}
		}
		dirPath = parentDirPath
	}
	return
}

// 事件循环(核心逻辑)
func (w *TWatcher) startEventLoop() {
	go func() {
		for {
			if v := w.events.Pop(); v != nil {
				event := v.(*TEvent)
				// 如果该路径一个回调也没有，那么没有必要执行后续逻辑，删除对该文件的监听
				callbacks := w.getCallbacks(event.Path)
				if len(callbacks) == 0 {
					w.watcher.Remove(event.Path)
					continue
				}
				switch {
				// 如果是删除操作，那么需要判断是否文件真正不存在了，如果存在，那么将此事件认为“假删除”
				case event.IsRemove():
					if fileExists(event.Path) {
						// 底层重新添加监控(不用担心重复添加)
						w.watcher.Add(event.Path)
						// 修改事件操作为重命名(相当于重命名为自身名称，最终名称没变)
						event.Op = RENAME
					}

				// 如果是重命名操作，那么需要判断是否文件真正不存在了，如果存在，那么将此事件认为“假命名”
				// (特别是某些编辑器在编辑文件时会先对文件RENAME再CHMOD)
				case event.IsRename():
					if fileExists(event.Path) {
						// 底层有可能去掉了监控, 这里重新添加监控(不用担心重复添加)
						w.watcher.Add(event.Path)
						// 修改事件操作为修改属性
						event.Op = CHMOD
					}

				// 创建文件/目录
				case event.IsCreate():
					// =========================================
					// 注意这里只是添加底层监听，并没有注册任何的回调函数，
					// 默认的回调函数为父级的递归回调
					// =========================================
					if fileIsDir(event.Path) {
						// 递归添加
						for _, subPath := range fileAllDirs(event.Path) {
							if fileIsDir(subPath) {
								w.watcher.Add(subPath)
							}
						}
					} else {
						// 添加文件监听
						w.watcher.Add(event.Path)
					}

				}
				// 执行回调处理，异步处理
				for _, v := range callbacks {
					go func(callback *Callback) {
						defer func() {
							// 是否退出监控
							if err := recover(); err != nil {
								switch err {
								case mFSNOTIFY_EVENT_EXIT:
									w.RemoveCallback(callback.Id)
								default:
									panic(err)
								}
							}
						}()
						callback.Func(event)
					}(v)
				}
			} else {
				break
			}
		}
	}()
}
