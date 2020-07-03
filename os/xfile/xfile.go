package xfile

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/go-xe2/x/type/xstring"
	"github.com/go-xe2/x/utils/xregex"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"time"
)

const (
	// 系统路径分割符
	Separator = string(filepath.Separator)
	// 默认文件操作权限
	mDEFAULT_PERM = 0666
)

var (
	mainPkgPath = xstring.New()
)

// 以路径分隔符连接路径并返回
func Join(subPath ...string) string {
	return xstring.Join(subPath, Separator)
}

// 创建文件夹
func Mkdir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// 创建文件
func Create(path string) (*os.File, error) {
	dir := Dir(path)
	if !Exists(dir) {
		Mkdir(dir)
	}
	return os.Create(path)
}

// Open opens file/directory readonly.
// 以只读方式打开文件或文件夹
func Open(path string) (*os.File, error) {
	return os.Open(path)
}

// 以指定文件打开文件或文件夹
func OpenFile(path string, flag int, perm os.FileMode) (*os.File, error) {
	return os.OpenFile(path, flag, perm)
}

// 以默认权限打开文件或文件夹
func OpenWithFlag(path string, flag int) (*os.File, error) {
	f, err := os.OpenFile(path, flag, mDEFAULT_PERM)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// 以指定权限打开文件或文件夹
func OpenWithFlagPerm(path string, flag int, perm int) (*os.File, error) {
	f, err := os.OpenFile(path, flag, os.FileMode(perm))
	if err != nil {
		return nil, err
	}
	return f, nil
}

// 检查文件是否存在
func Exists(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}
	return false
}

// 检查路径是否是文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 返回当前工作目录
func Pwd() string {
	path, _ := os.Getwd()
	return path
}

// 检查路径是否是文件
func IsFile(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !s.IsDir()
}

// 获取路径信息,Stat别名
func Info(path string) (os.FileInfo, error) {
	return Stat(path)
}

// Stat returns a FileInfo describing the named file.
// If there is an error, it will be of type *PathError.
// 获取路径信息
func Stat(path string) (os.FileInfo, error) {
	return os.Stat(path)
}

// 移动文件或文件夹
func Move(src string, dst string) error {
	return os.Rename(src, dst)
}

// 重命名文件或文件夹
func Rename(src string, dst string) error {
	return Move(src, dst)
}

// 复制文件或文件夹
func Copy(src string, dst string) error {
	if IsFile(src) {
		return CopyFile(src, dst)
	}
	return CopyDir(src, dst)
}

// 复制文件
func CopyFile(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer func() {
		if e := in.Close(); e != nil {
			err = e
		}
	}()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		if e := out.Close(); e != nil {
			err = e
		}
	}()
	_, err = io.Copy(out, in)
	if err != nil {
		return
	}
	err = out.Sync()
	if err != nil {
		return
	}
	si, err := os.Stat(src)
	if err != nil {
		return
	}
	err = os.Chmod(dst, si.Mode())
	if err != nil {
		return
	}
	return
}

// 复制文件夹
func CopyDir(src string, dst string) (err error) {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)
	si, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !si.IsDir() {
		return fmt.Errorf("source is not a directory")
	}
	_, err = os.Stat(dst)
	if err != nil && !os.IsNotExist(err) {
		return
	}
	if err == nil {
		return fmt.Errorf("destination already exists")
	}
	err = os.MkdirAll(dst, si.Mode())
	if err != nil {
		return
	}
	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return
	}
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())
		if entry.IsDir() {
			err = CopyDir(srcPath, dstPath)
			if err != nil {
				return
			}
		} else {
			// Skip symlinks.
			if entry.Mode()&os.ModeSymlink != 0 {
				continue
			}
			err = CopyFile(srcPath, dstPath)
			if err != nil {
				return
			}
		}
	}
	return
}

// 获取子文件夹列表
func DirNames(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdirnames(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	return list, nil
}

// 按指定规则查找获取文件或文件夹
func Glob(pattern string, onlyNames ...bool) ([]string, error) {
	if list, err := filepath.Glob(pattern); err == nil {
		if len(onlyNames) > 0 && onlyNames[0] && len(list) > 0 {
			array := make([]string, len(list))
			for k, v := range list {
				array[k] = Basename(v)
			}
			return array, nil
		}
		return list, nil
	} else {
		return nil, err
	}
}

// 删除文件或文件夹
func Remove(path string) error {
	return os.RemoveAll(path)
}

// 检查路径是否可读
func IsReadable(path string) bool {
	result := true
	file, err := os.OpenFile(path, os.O_RDONLY, mDEFAULT_PERM)
	if err != nil {
		result = false
	}
	file.Close()
	return result
}

// 检查路径是否可写
func IsWritable(path string) bool {
	result := true
	if IsDir(path) {
		// If it's a directory, create a temporary file to test whether it's writable.
		tmpFile := strings.TrimRight(path, Separator) + Separator + fmt.Sprintf("%d", time.Now().UnixNano())
		if f, err := Create(tmpFile); err != nil || !Exists(tmpFile) {
			result = false
		} else {
			f.Close()
			Remove(tmpFile)
		}
	} else {
		// 如果是文件，那么判断文件是否可打开
		file, err := os.OpenFile(path, os.O_WRONLY, mDEFAULT_PERM)
		if err != nil {
			result = false
		}
		file.Close()
	}
	return result
}

// 修改路径权限
func Chmod(path string, mode os.FileMode) error {
	return os.Chmod(path, mode)
}

// 扫描路径中的文件夹，包含子文件夹中的所有文件，返回文件列表
func ScanDir(path string, pattern string, recursive ...bool) ([]string, error) {
	list, err := doScanDir(path, pattern, recursive...)
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		sort.Strings(list)
	}
	return list, nil
}

// 扫描路径中的所有文件
func doScanDir(path string, pattern string, recursive ...bool) ([]string, error) {
	list := ([]string)(nil)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	names, err := file.Readdirnames(-1)
	if err != nil {
		return nil, err
	}
	for _, name := range names {
		path := fmt.Sprintf("%s%s%s", path, Separator, name)
		if IsDir(path) && len(recursive) > 0 && recursive[0] {
			array, _ := doScanDir(path, pattern, true)
			if len(array) > 0 {
				list = append(list, array...)
			}
		}
		// If it meets pattern, then add it to the result list.
		for _, p := range strings.Split(pattern, ",") {
			if match, err := filepath.Match(strings.TrimSpace(p), name); err == nil && match {
				list = append(list, path)
			}
		}
	}
	return list, nil
}

// 获取路径的绝对路径，如果路径不存在返回空字符串
func RealPath(path string) string {
	p, err := filepath.Abs(path)
	if err != nil {
		return ""
	}
	if !Exists(p) {
		return ""
	}
	return p
}

// 获取当前程序所在路径
func SelfPath() string {
	p, _ := filepath.Abs(os.Args[0])
	return p
}

// 获取当前程序名称
func SelfName() string {
	return Basename(SelfPath())
}

// 获取当前程序所在路径
func SelfDir() string {
	return filepath.Dir(SelfPath())
}

// 获取最内层文件夹或文件名称，如果不存在，返回"."
func Basename(path string) string {
	return filepath.Base(path)
}

// 获取文件所在文件夹，如果path=""，则返回"."
func Dir(path string) string {
	return filepath.Dir(path)
}

// 获取文件后掇，返回文件后缀中包含"."
func Ext(path string) string {
	return filepath.Ext(path)
}

// 获取用户根目录
func Home() (string, error) {
	u, err := user.Current()
	if nil == err {
		return u.HomeDir, nil
	}
	if "windows" == runtime.GOOS {
		return homeWindows()
	}
	return homeUnix()
}

// 获取unix类型系统用户根目录
func homeUnix() (string, error) {
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}

// 获取windows系统用户根目录
func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}

	return home, nil
}

// 获取主项目目录
func MainPkgPath() string {
	path := mainPkgPath.String()
	if path != "" {
		if path == "-" {
			return ""
		}
		return path
	}
	for i := 1; i < 10000; i++ {
		if _, file, _, ok := runtime.Caller(i); ok {
			// <file> is separated by '/'
			if xstring.New(file).Contains("/go-xe2/core/") {
				continue
			}
			if Ext(file) != ".go" {
				continue
			}
			// separator of <file> '/' will be converted to Separator.
			for path = Dir(file); len(path) > 1 && Exists(path) && path[len(path)-1] != os.PathSeparator; {
				files, _ := ScanDir(path, "*.go")
				for _, v := range files {
					if xregex.IsMatchString(`package\s+main`, GetContents(v)) {
						mainPkgPath = xstring.New(path)
						return path
					}
				}
				path = Dir(path)
			}

		} else {
			break
		}
	}
	// If it fails finding the path, then mark it as "-",
	// which means it will never do this search again.
	mainPkgPath = xstring.New("-")
	return ""
}

// 临时目录
func TempDir() string {
	return os.TempDir()
}
