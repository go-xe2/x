package xfileNotify

import "github.com/go-xe2/third/github.com/fsnotify/fsnotify"

// 监听事件对象
type TEvent struct {
	event   fsnotify.Event // 底层事件对象
	Path    string         // 文件绝对路径
	Op      Op             // 触发监听的文件操作
	Watcher *TWatcher      // 事件对应的监听对象
}

func (e *TEvent) String() string {
	return e.event.String()
}

// 文件/目录创建
func (e *TEvent) IsCreate() bool {
	return e.Op == 1 || e.Op&CREATE == CREATE
}

// 文件/目录修改
func (e *TEvent) IsWrite() bool {
	return e.Op&WRITE == WRITE
}

// 文件/目录删除
func (e *TEvent) IsRemove() bool {
	return e.Op&REMOVE == REMOVE
}

// 文件/目录重命名
func (e *TEvent) IsRename() bool {
	return e.Op&RENAME == RENAME
}

// 文件/目录修改权限
func (e *TEvent) IsChmod() bool {
	return e.Op&CHMOD == CHMOD
}
