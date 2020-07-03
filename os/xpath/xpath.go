package xpath

import (
	"errors"
	"fmt"
	"github.com/go-xe2/x/container/xarray"
	"github.com/go-xe2/x/os/xfile"
	"github.com/go-xe2/x/sync/xsafeMap"
	"github.com/go-xe2/x/type/xstring"
	"os"
	"sort"
	"strings"
)

// 文件目录搜索管理对象
type TPath struct {
	paths *xarray.TStringArray // 搜索路径，按照优先级进行排序
	cache *xsafeMap.TStrStrMap // 搜索结果缓存map(如果未nil表示未启用缓存功能)
}

// 文件搜索缓存项
type TPathCacheItem struct {
	path  string // 文件/目录绝对路径
	isDir bool   // 是否目录
}

var (
	// 单个目录路径对应的TPath对象指针，用于路径检索对象复用
	pathsMap      = xsafeMap.NewStrAnyMap()
	pathsCacheMap = xsafeMap.NewStrAnyMap()
)

// 创建一个搜索对象
func New(path string, cache bool) *TPath {
	sp := &TPath{
		paths: xarray.NewStringArray(),
	}
	if cache {
		sp.cache = xsafeMap.NewStrStrMap()
	}
	if len(path) > 0 {
		if _, err := sp.Add(path); err != nil {
			fmt.Println(err.Error())
		}
	}
	return sp
}

// 创建/获取一个单例的搜索对象, root必须为目录的绝对路径
func Get(root string, cache bool) *TPath {
	return pathsMap.GetOrSetFuncLock(root, func() interface{} {
		return New(root, cache)
	}).(*TPath)
}

// 检索root目录(必须为绝对路径)下面的name文件的绝对路径，indexFiles用于指定当检索到的结果为目录时，同时检索是否存在这些indexFiles文件
func Search(root string, name string, indexFiles ...string) (filePath string, isDir bool) {
	return Get(root, false).Search(name, indexFiles...)
}

// 检索root目录(必须为绝对路径)下面的name文件的绝对路径，indexFiles用于指定当检索到的结果为目录时，同时检索是否存在这些indexFiles文件
func SearchWithCache(root string, name string, indexFiles ...string) (filePath string, isDir bool) {
	return Get(root, true).Search(name, indexFiles...)
}

// 设置搜索路径，只保留当前设置项，其他搜索路径被清空
func (sp *TPath) Set(path string) (realPath string, err error) {
	realPath = xfile.RealPath(path)
	if realPath == "" {
		realPath, _ = sp.Search(path)
		if realPath == "" {
			realPath = xfile.RealPath(xfile.Pwd() + xfile.Separator + path)
		}
	}
	if realPath == "" {
		return realPath, errors.New(fmt.Sprintf(`path "%s" does not exist`, path))
	}
	// 设置的搜索路径必须为目录
	if xfile.IsDir(realPath) {
		realPath = strings.TrimRight(realPath, xfile.Separator)
		if sp.paths.Search(realPath) != -1 {
			for _, v := range sp.paths.Slice() {
				sp.removeMonitorByPath(v)
			}
		}
		sp.paths.Clear()
		if sp.cache != nil {
			sp.cache.Clear()
		}
		sp.paths.Append(realPath)
		sp.updateCacheByPath(realPath)
		sp.addMonitorByPath(realPath)
		return realPath, nil
	} else {
		return "", errors.New(path + " should be a folder")
	}
}

// 添加搜索路径
func (sp *TPath) Add(path string) (realPath string, err error) {
	realPath = xfile.RealPath(path)
	if realPath == "" {
		realPath, _ = sp.Search(path)
		if realPath == "" {
			realPath = xfile.RealPath(xfile.Pwd() + xfile.Separator + path)
		}
	}
	if realPath == "" {
		return realPath, errors.New(fmt.Sprintf(`path "%s" does not exist`, path))
	}
	// 添加的搜索路径必须为目录
	if xfile.IsDir(realPath) {
		//fmt.Println("gspath:", realPath, sp.paths.Search(realPath))
		// 如果已经添加则不再添加
		if sp.paths.Search(realPath) < 0 {
			realPath = strings.TrimRight(realPath, xfile.Separator)
			sp.paths.Append(realPath)
			sp.updateCacheByPath(realPath)
			sp.addMonitorByPath(realPath)
		}
		return realPath, nil
	} else {
		return "", errors.New(path + " should be a folder")
	}
}

// 给定的name只是相对文件路径，找不到该文件时，返回空字符串;
// 当给定indexFiles时，如果name是一个目录，那么会进一步检索其下对应的indexFiles文件是否存在，存在则返回indexFile绝对路径；
// 否则返回name目录绝对路径。
func (sp *TPath) Search(name string, indexFiles ...string) (filePath string, isDir bool) {
	// 不使用缓存
	if sp.cache == nil {
		sp.paths.LockFunc(func(array []string) {
			path := ""
			for _, v := range array {
				path = v + xfile.Separator + name
				if stat, err := os.Stat(path); !os.IsNotExist(err) {
					filePath = path
					isDir = stat.IsDir()
					break
				}
			}
		})
		if len(indexFiles) > 0 && isDir {
			if name == "/" {
				name = ""
			}
			path := ""
			for _, file := range indexFiles {
				path = filePath + xfile.Separator + file
				if xfile.Exists(path) {
					filePath = path
					isDir = false
					break
				}
			}
		}
		return
	}
	// 使用缓存功能
	name = sp.formatCacheName(name)
	if v := sp.cache.Get(name); v != "" {
		filePath, isDir = sp.parseCacheValue(v)
		if len(indexFiles) > 0 && isDir {
			if name == "/" {
				name = ""
			}
			for _, file := range indexFiles {
				if v := sp.cache.Get(name + "/" + file); v != "" {
					return sp.parseCacheValue(v)
				}
			}
		}
	}
	return
}

// 从搜索路径中移除指定的文件，这样该文件无法给搜索。
// path可以是绝对路径，也可以相对路径。
func (sp *TPath) Remove(path string) {
	if sp.cache == nil {
		return
	}
	if xfile.Exists(path) {
		for _, v := range sp.paths.Slice() {
			name := xstring.Replace(path, v, "")
			name = sp.formatCacheName(name)
			sp.cache.Remove(name)
		}
	} else {
		name := sp.formatCacheName(path)
		sp.cache.Remove(name)
	}
}

// 返回当前对象搜索目录路径列表
func (sp *TPath) Paths() []string {
	return sp.paths.Slice()
}

// 返回当前对象缓存的所有路径列表
func (sp *TPath) AllPaths() []string {
	if sp.cache == nil {
		return nil
	}
	paths := sp.cache.Keys()
	if len(paths) > 0 {
		sort.Strings(paths)
	}
	return paths
}

// 当前的搜索路径数量
func (sp *TPath) Size() int {
	return sp.paths.Len()
}
