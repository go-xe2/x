package xpath

import (
	"github.com/go-xe2/x/os/xfile"
	"github.com/go-xe2/x/os/xfileNotify"
	"github.com/go-xe2/x/type/xstring"
	"runtime"
	"strings"
)

// 递归添加目录下的文件
func (sp *TPath) updateCacheByPath(path string) {
	if sp.cache == nil {
		return
	}
	sp.addToCache(path, path)
}

// 格式化name返回符合规范的缓存名称，分隔符号统一为'/'，且前缀必须以'/'开头(类似HTTP URI).
func (sp *TPath) formatCacheName(name string) string {
	if runtime.GOOS != "linux" {
		name = xstring.Replace(name, "\\", "/")
	}
	return "/" + strings.Trim(name, "./")
}

// 根据path计算出对应的缓存name, dirPath为检索根目录路径
func (sp *TPath) nameFromPath(filePath, rootPath string) string {
	name := xstring.Replace(filePath, rootPath, "")
	name = sp.formatCacheName(name)
	return name
}

// 按照一定数据结构生成缓存的数据项字符串
func (sp *TPath) makeCacheValue(filePath string, isDir bool) string {
	if isDir {
		return filePath + "_D_"
	}
	return filePath + "_F_"
}

// 按照一定数据结构解析数据项字符串
func (sp *TPath) parseCacheValue(value string) (filePath string, isDir bool) {
	if value[len(value)-2 : len(value)-1][0] == 'F' {
		return value[:len(value)-3], false
	}
	return value[:len(value)-3], true
}

// 添加path到缓存中(递归)
func (sp *TPath) addToCache(filePath, rootPath string) {
	// 首先添加自身
	idDir := xfile.IsDir(filePath)
	sp.cache.SetIfNotExist(sp.nameFromPath(filePath, rootPath), sp.makeCacheValue(filePath, idDir))
	// 如果添加的是目录，那么需要递归添加
	if idDir {
		if files, err := xfile.ScanDir(filePath, "*", true); err == nil {
			//fmt.Println("gspath add to cache:", filePath, files)
			for _, path := range files {
				sp.cache.SetIfNotExist(sp.nameFromPath(path, rootPath), sp.makeCacheValue(path, xfile.IsDir(path)))
			}
		} else {
			//fmt.Errorf(err.Error())
		}
	}
}

// 添加文件目录监控(递归)，当目录下的文件有更新时，会同时更新缓存。
// 这里需要注意的点是，由于添加监听是递归添加的，那么假如删除一个目录，那么该目录下的文件(包括目录)也会产生一条删除事件，总共会产生N条事件。
func (sp *TPath) addMonitorByPath(path string) {
	if sp.cache == nil {
		return
	}
	_, _ = xfileNotify.Add(path, func(event *xfileNotify.TEvent) {
		//glog.Debug(event.String())
		switch {
		case event.IsRemove():
			sp.cache.Remove(sp.nameFromPath(event.Path, path))

		case event.IsRename():
			if !xfile.Exists(event.Path) {
				sp.cache.Remove(sp.nameFromPath(event.Path, path))
			}

		case event.IsCreate():
			sp.addToCache(event.Path, path)
		}
	}, true)
}

// 删除监听(递归)
func (sp *TPath) removeMonitorByPath(path string) {
	if sp.cache == nil {
		return
	}
	_ = xfileNotify.Remove(path)
}
