package xconfig

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/go-xe2/x/container/xarray"
	"github.com/go-xe2/x/core/cmdenv"
	"github.com/go-xe2/x/encoding/xjson"
	"github.com/go-xe2/x/os/xfile"
	"github.com/go-xe2/x/os/xfileNotify"
	"github.com/go-xe2/x/os/xlog"
	"github.com/go-xe2/x/os/xpath"
	_type "github.com/go-xe2/x/sync/type"
	"github.com/go-xe2/x/sync/xsafeMap"
	"github.com/go-xe2/x/type/xstring"
	"github.com/go-xe2/x/type/xtime"
	"time"
)

const (
	DEFAULT_CONFIG_FILE = "config.yaml"
)

type TConfig struct {
	name  *xstring.String      // 默认配置文件名
	paths *xarray.TStringArray // 配置文件搜索路径
	jsons *xsafeMap.TStrAnyMap
	vc    *_type.TBool // 是否启用访问子层，默认false
}

func New(file ...string) *TConfig {
	name := DEFAULT_CONFIG_FILE
	if len(file) > 0 {
		name = file[0]
	}
	c := &TConfig{
		name:  xstring.New(name),
		paths: xarray.NewStringArray(),
		jsons: xsafeMap.NewStrAnyMap(),
		vc:    _type.NewBool(),
	}
	if envPath := cmdenv.Get("x.config.path").String(); envPath != "" {
		if xfile.Exists(envPath) {
			_ = c.SetPath(envPath)
		} else {
			if errorPrint() {
				xlog.Errorf("Configuration directory path does not exist: %s", envPath)
			}
		}
	} else {
		_ = c.SetPath(xfile.Pwd())
		if selfPath := xfile.SelfDir(); selfPath != "" && xfile.Exists(selfPath) {
			_ = c.AddPath(selfPath)
		}
		if mainPath := xfile.MainPkgPath(); mainPath != "" && xfile.Exists(mainPath) {
			_ = c.AddPath(mainPath)
		}
	}
	return c
}

func (c *TConfig) filePath(file ...string) (path string) {
	name := c.name.String()
	if len(file) > 0 {
		name = file[0]
	}
	path = c.FilePath(name)
	if path == "" {
		buffer := bytes.NewBuffer(nil)
		if c.paths.Len() > 0 {
			buffer.WriteString(fmt.Sprintf("[xconfig] cannot find config file \"%s\" in following paths:", name))
			c.paths.RLockFunc(func(array []string) {
				index := 1
				for _, v := range array {
					buffer.WriteString(fmt.Sprintf("\n%d. %s", index, v))
					index++
					buffer.WriteString(fmt.Sprintf("\n%d. %s", index, v+xfile.Separator+"config"))
					index++
				}
			})
		} else {
			buffer.WriteString(fmt.Sprintf("[xconfig] cannot find config file \"%s\" with no path set/add", name))
		}
		if errorPrint() {
			xlog.Error(buffer.String())
		}
	}
	return path
}

func (c *TConfig) SetPath(path string) error {
	// 获取绝对路径
	realPath := xfile.RealPath(path)
	if realPath == "" {
		c.paths.RLockFunc(func(array []string) {
			for _, v := range array {
				if path, _ := xpath.Search(v, path); path != "" {
					realPath = path
					break
				}
			}
		})
	}
	if realPath == "" {
		buffer := bytes.NewBuffer(nil)
		if c.paths.Len() > 0 {
			buffer.WriteString(fmt.Sprintf("[xconfig] SetPath failed: cannot find directory \"%s\" in following paths:", path))
			c.paths.RLockFunc(func(array []string) {
				for k, v := range array {
					buffer.WriteString(fmt.Sprintf("\n%d. %s", k+1, v))
				}
			})
		} else {
			buffer.WriteString(fmt.Sprintf(`[xconfig] SetPath failed: path "%s" does not exist`, path))
		}
		err := errors.New(buffer.String())
		if errorPrint() {
			xlog.Error(err)
		}
		return err
	}
	if !xfile.IsDir(realPath) {
		err := fmt.Errorf(`[xconfig] SetPath failed: path "%s" should be directory type`, path)
		if errorPrint() {
			xlog.Error(err)
		}
		return err
	}
	// Repeated path check.
	if c.paths.Search(realPath) != -1 {
		return nil
	}
	c.jsons.Clear()
	c.paths.Clear()
	c.paths.Append(realPath)
	return nil
}

// 设置是否检查访问子层，默认false
func (c *TConfig) SetViolenceCheck(check bool) {
	c.vc.Set(check)
	c.Clear()
}

// 增加搜索目录
func (c *TConfig) AddPath(path string) error {
	realPath := xfile.RealPath(path)
	if realPath == "" {
		c.paths.RLockFunc(func(array []string) {
			for _, v := range array {
				if path, _ := xpath.Search(v, path); path != "" {
					realPath = path
					break
				}
			}
		})
	}
	if realPath == "" {
		buffer := bytes.NewBuffer(nil)
		if c.paths.Len() > 0 {
			buffer.WriteString(fmt.Sprintf("[xconfig] AddPath failed: cannot find directory \"%s\" in following paths:", path))
			c.paths.RLockFunc(func(array []string) {
				for k, v := range array {
					buffer.WriteString(fmt.Sprintf("\n%d. %s", k+1, v))
				}
			})
		} else {
			buffer.WriteString(fmt.Sprintf(`[xconfig] AddPath failed: path "%s" does not exist`, path))
		}
		err := errors.New(buffer.String())
		if errorPrint() {
			xlog.Error(err)
		}
		return err
	}
	if !xfile.IsDir(realPath) {
		err := fmt.Errorf(`[xconfig] AddPath failed: path "%s" should be directory type`, path)
		if errorPrint() {
			xlog.Error(err)
		}
		return err
	}
	if c.paths.Search(realPath) != -1 {
		return nil
	}
	c.paths.Append(realPath)
	return nil
}

// 获取配置文件路径
func (c *TConfig) GetFilePath(file ...string) (path string) {
	return c.FilePath(file...)
}

func (c *TConfig) FilePath(file ...string) (path string) {
	name := c.name.String()
	if len(file) > 0 {
		name = file[0]
	}
	c.paths.RLockFunc(func(array []string) {
		for _, v := range array {
			if path, _ = xpath.Search(v, name); path != "" {
				break
			}
			if path, _ = xpath.Search(v+xfile.Separator+"config", name); path != "" {
				break
			}
		}
	})
	return
}

// 设置配置文件名
func (c *TConfig) SetFileName(name string) {
	c.name.Set(name)
}

// 获取配置文件名
func (c *TConfig) GetFileName() string {
	return c.name.String()
}

// 获取配置json对角值， 如果不存在或出错，返回nil
func (c *TConfig) getJson(file ...string) *xjson.TJson {
	name := c.name.String()
	if len(file) > 0 {
		name = file[0]
	}
	r := c.jsons.GetOrSetFuncLock(name, func() interface{} {
		content := ""
		filePath := ""
		if content = GetContent(name); content == "" {
			filePath = c.filePath(name)
			if filePath == "" {
				return nil
			}
			content = xfile.GetContents(filePath)
		}
		if j, err := xjson.LoadContent(content); err == nil {
			j.SetViolenceCheck(c.vc.Val())
			// Add monitor for this configuration file,
			// any changes of this file will refresh its cache in TConfig object.
			if filePath != "" {
				_, err = xfileNotify.Add(filePath, func(event *xfileNotify.TEvent) {
					c.jsons.Remove(name)
				})
				if err != nil && errorPrint() {
					xlog.Error(err)
				}
			}
			return j
		} else {
			if errorPrint() {
				if filePath != "" {
					xlog.Criticalf(`[xconfig] Load config file "%s" failed: %s`, filePath, err.Error())
				} else {
					xlog.Criticalf(`[xconfig] Load configuration failed: %s`, err.Error())
				}
			}
		}
		return nil
	})
	if r != nil {
		return r.(*xjson.TJson)
	}
	return nil
}

func (c *TConfig) Get(pattern string, def ...interface{}) interface{} {
	if j := c.getJson(); j != nil {
		return j.Get(pattern, def...)
	}
	return nil
}

func (c *TConfig) GetVar(pattern string, def ...interface{}) *_type.TVar {
	if j := c.getJson(); j != nil {
		return _type.NewVar(j.Get(pattern, def...), true)
	}
	return _type.NewVar(nil, true)
}

func (c *TConfig) Contains(pattern string) bool {
	if j := c.getJson(); j != nil {
		return j.Contains(pattern)
	}
	return false
}

func (c *TConfig) GetMap(pattern string, def ...interface{}) map[string]interface{} {
	if j := c.getJson(); j != nil {
		return j.GetMap(pattern, def...)
	}
	return nil
}

func (c *TConfig) GetArray(pattern string, def ...interface{}) []interface{} {
	if j := c.getJson(); j != nil {
		return j.GetArray(pattern, def...)
	}
	return nil
}

func (c *TConfig) GetString(pattern string, def ...interface{}) string {
	if j := c.getJson(); j != nil {
		return j.GetString(pattern, def...)
	}
	return ""
}

func (c *TConfig) GetStrings(pattern string, def ...interface{}) []string {
	if j := c.getJson(); j != nil {
		return j.GetStrings(pattern, def...)
	}
	return nil
}

func (c *TConfig) GetInterfaces(pattern string, def ...interface{}) []interface{} {
	if j := c.getJson(); j != nil {
		return j.GetInterfaces(pattern, def...)
	}
	return nil
}

func (c *TConfig) GetBool(pattern string, def ...interface{}) bool {
	if j := c.getJson(); j != nil {
		return j.GetBool(pattern, def...)
	}
	return false
}

func (c *TConfig) GetFloat32(pattern string, def ...interface{}) float32 {
	if j := c.getJson(); j != nil {
		return j.GetFloat32(pattern, def...)
	}
	return 0
}

func (c *TConfig) GetFloat64(pattern string, def ...interface{}) float64 {
	if j := c.getJson(); j != nil {
		return j.GetFloat64(pattern, def...)
	}
	return 0
}

func (c *TConfig) GetFloats(pattern string, def ...interface{}) []float64 {
	if j := c.getJson(); j != nil {
		return j.GetFloats(pattern, def...)
	}
	return nil
}

func (c *TConfig) GetInt(pattern string, def ...interface{}) int {
	if j := c.getJson(); j != nil {
		return j.GetInt(pattern, def...)
	}
	return 0
}

func (c *TConfig) GetInt8(pattern string, def ...interface{}) int8 {
	if j := c.getJson(); j != nil {
		return j.GetInt8(pattern, def...)
	}
	return 0
}

func (c *TConfig) GetInt16(pattern string, def ...interface{}) int16 {
	if j := c.getJson(); j != nil {
		return j.GetInt16(pattern, def...)
	}
	return 0
}

func (c *TConfig) GetInt32(pattern string, def ...interface{}) int32 {
	if j := c.getJson(); j != nil {
		return j.GetInt32(pattern, def...)
	}
	return 0
}

func (c *TConfig) GetInt64(pattern string, def ...interface{}) int64 {
	if j := c.getJson(); j != nil {
		return j.GetInt64(pattern, def...)
	}
	return 0
}

func (c *TConfig) GetInts(pattern string, def ...interface{}) []int {
	if j := c.getJson(); j != nil {
		return j.GetInts(pattern, def...)
	}
	return nil
}

func (c *TConfig) GetUint(pattern string, def ...interface{}) uint {
	if j := c.getJson(); j != nil {
		return j.GetUint(pattern, def...)
	}
	return 0
}

func (c *TConfig) GetUint8(pattern string, def ...interface{}) uint8 {
	if j := c.getJson(); j != nil {
		return j.GetUint8(pattern, def...)
	}
	return 0
}

func (c *TConfig) GetUint16(pattern string, def ...interface{}) uint16 {
	if j := c.getJson(); j != nil {
		return j.GetUint16(pattern, def...)
	}
	return 0
}

func (c *TConfig) GetUint32(pattern string, def ...interface{}) uint32 {
	if j := c.getJson(); j != nil {
		return j.GetUint32(pattern, def...)
	}
	return 0
}

func (c *TConfig) GetUint64(pattern string, def ...interface{}) uint64 {
	if j := c.getJson(); j != nil {
		return j.GetUint64(pattern, def...)
	}
	return 0
}

func (c *TConfig) GetTime(pattern string, format ...string) time.Time {
	if j := c.getJson(); j != nil {
		return j.GetTime(pattern, format...)
	}
	return time.Time{}
}

func (c *TConfig) GetDuration(pattern string, def ...interface{}) time.Duration {
	if j := c.getJson(); j != nil {
		return j.GetDuration(pattern, def...)
	}
	return 0
}

func (c *TConfig) GetXTime(pattern string, format ...string) *xtime.Time {
	if j := c.getJson(); j != nil {
		return j.GetGTime(pattern, format...)
	}
	return nil
}

func (c *TConfig) GetStruct(pattern string, pointer interface{}, mapping ...map[string]string) error {
	if j := c.getJson(); j != nil {
		return j.GetStruct(pattern, pointer, mapping...)
	}
	return errors.New("config file not found")
}

func (c *TConfig) GetStructDeep(pattern string, pointer interface{}, mapping ...map[string]string) error {
	if j := c.getJson(); j != nil {
		return j.GetStructDeep(pattern, pointer, mapping...)
	}
	return errors.New("config file not found")
}

func (c *TConfig) GetStructs(pattern string, pointer interface{}, mapping ...map[string]string) error {
	if j := c.getJson(); j != nil {
		return j.GetStructs(pattern, pointer, mapping...)
	}
	return errors.New("config file not found")
}

func (c *TConfig) GetStructsDeep(pattern string, pointer interface{}, mapping ...map[string]string) error {
	if j := c.getJson(); j != nil {
		return j.GetStructsDeep(pattern, pointer, mapping...)
	}
	return errors.New("config file not found")
}

func (c *TConfig) Clear() {
	c.jsons.Clear()
}
