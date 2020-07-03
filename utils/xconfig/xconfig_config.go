package xconfig

import (
	"github.com/go-xe2/x/sync/xsafeMap"
)

var (
	// Customized configuration content.
	configs = xsafeMap.NewStrStrMap()
)

// 保存配置内容
func SetContent(content string, file ...string) {
	name := DEFAULT_CONFIG_FILE
	if len(file) > 0 {
		name = file[0]
	}
	// Clear file cache for instances which cached <name>.
	instances.LockFunc(func(m map[string]interface{}) {
		if configs.Contains(name) {
			for _, v := range m {
				v.(*TConfig).jsons.Remove(name)
			}
		}
		configs.Set(name, content)
	})
}

// 获取配置文件内容
func GetContent(file ...string) string {
	name := DEFAULT_CONFIG_FILE
	if len(file) > 0 {
		name = file[0]
	}
	return configs.Get(name)
}

// 删除配置文件, 并从缓存中删除
func RemoveConfig(file ...string) {
	name := DEFAULT_CONFIG_FILE
	if len(file) > 0 {
		name = file[0]
	}
	instances.LockFunc(func(m map[string]interface{}) {
		if configs.Contains(name) {
			for _, v := range m {
				v.(*TConfig).jsons.Remove(name)
			}
			configs.Remove(name)
		}
	})
}

func Config(name ...string) *TConfig {
	return Instance(name...)
}

// 清空配置内容
func ClearContent() {
	configs.Clear()
	// Clear cache for all instances.
	instances.LockFunc(func(m map[string]interface{}) {
		for _, v := range m {
			v.(*TConfig).jsons.Clear()
		}
	})
}
