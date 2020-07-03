package xconfig

import (
	"github.com/go-xe2/x/sync/xsafeMap"
)

const (
	// 默认配置实例名称
	DEFAULT_GROUP_NAME = "default"
)

var (
	// 所有配置实例键值对
	instances = xsafeMap.NewStrAnyMap()
)

// 获取配置实例
func Instance(name ...string) *TConfig {
	key := DEFAULT_GROUP_NAME
	if len(name) > 0 {
		key = name[0]
	}
	return instances.GetOrSetFuncLock(key, func() interface{} {
		return New()
	}).(*TConfig)
}
