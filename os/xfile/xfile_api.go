package xfile

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/go-xe2/x/container/xarray"
	"os"
)

// 查找文件，找到返回文件路径，否则返回空字符串
func Search(name string, prioritySearchPaths ...string) (realPath string, err error) {
	realPath = RealPath(name)
	if realPath != "" {
		return
	}
	// TODO move search paths to internal package variable.
	array := xarray.NewStringArray(true)
	array.Append(prioritySearchPaths...)
	array.Append(Pwd(), SelfDir())
	if path := MainPkgPath(); path != "" {
		array.Append(path)
	}
	array.Unique()
	array.RLockFunc(func(array []string) {
		path := ""
		for _, v := range array {
			path = RealPath(v + Separator + name)
			if path != "" {
				realPath = path
				break
			}
		}
	})
	if realPath == "" {
		buffer := bytes.NewBuffer(nil)
		buffer.WriteString(fmt.Sprintf("cannot find file/folder \"%s\" in following paths:", name))
		array.RLockFunc(func(array []string) {
			for k, v := range array {
				buffer.WriteString(fmt.Sprintf("\n%d. %s", k+1, v))
			}
		})
		err = errors.New(buffer.String())
	}
	return
}

// 获取文件大小,返回字节数
func Size(path string) int64 {
	s, e := os.Stat(path)
	if e != nil {
		return 0
	}
	return s.Size()
}

// 获取文件大小，返回格式化后字符串
func ReadableSize(path string) string {
	return FormatSize(float64(Size(path)))
}

// 格式化文件大小
func FormatSize(raw float64) string {
	var t float64 = 1024
	var d float64 = 1

	if raw < t {
		return fmt.Sprintf("%.2fB", raw/d)
	}

	d *= 1024
	t *= 1024

	if raw < t {
		return fmt.Sprintf("%.2fK", raw/d)
	}

	d *= 1024
	t *= 1024

	if raw < t {
		return fmt.Sprintf("%.2fM", raw/d)
	}

	d *= 1024
	t *= 1024

	if raw < t {
		return fmt.Sprintf("%.2fG", raw/d)
	}

	d *= 1024
	t *= 1024

	if raw < t {
		return fmt.Sprintf("%.2fT", raw/d)
	}

	d *= 1024
	t *= 1024

	if raw < t {
		return fmt.Sprintf("%.2fP", raw/d)
	}

	return "TooLarge"
}

// 获取文件修改时间，秒
func MTime(path string) int64 {
	s, e := os.Stat(path)
	if e != nil {
		return 0
	}
	return s.ModTime().Unix()
}

// 获取文件修改时间，毫秒
func MTimeMillisecond(path string) int64 {
	s, e := os.Stat(path)
	if e != nil {
		return 0
	}
	return int64(s.ModTime().Nanosecond() / 1000000)
}
