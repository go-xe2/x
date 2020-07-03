package xf

import (
	"github.com/go-xe2/x/sync/xsafeMap"
	"github.com/go-xe2/x/type/t"
	"reflect"
)

var xfConfigs = xsafeMap.NewStrAnyMap()

// 读写运行配置
// v 支持1-2个参数
// len(v) = 1时，v 可以设置为实际值、func() interface{}获取值的函数、reflect.Kind数据类型,
// 为reflect.kind时，为按kind类型转换返回结果, 当reflect.kind = struct时，转入的第三个参数为struct的reflect.Type类型
func C(key string, v ...interface{}) interface{} {
	if len(v) > 0 {
		val := v[0]
		if fn, ok := val.(func() interface{}); ok {
			val = fn()
		} else if vt, ok := val.(reflect.Kind); ok {
			getV := xfConfigs.Get(key)
			switch vt {
			case String:
				return t.String(getV)
			case Bool:
				return t.Bool(getV)
			case Float32:
				return t.Float32(getV)
			case Float64:
				return t.Float64(getV)
			case Int:
				return t.Int(getV)
			case Int8:
				return t.Int8(getV)
			case Int16:
				return t.Int16(getV)
			case Int32:
				return t.Int32(getV)
			case Int64:
				return t.Int64(getV)
			case Uint:
				return t.Uint(getV)
			case Uint8:
				return t.Uint8(getV)
			case Uint16:
				return t.Uint16(getV)
			case Uint32:
				return t.Uint32(getV)
			case Uint64:
				return t.Uint64(getV)
			case Complex64:
				if cl, ok := getV.(complex64); ok {
					return cl
				}
				return 0
			case Complex128:
				if cl, ok := getV.(complex128); ok {
					return cl
				}
				return 0
			case Slice, Array:
				return t.SliceAny(getV)
			case Map:
				return t.Map(getV)
			case Struct:
				if len(v) == 3 {
					if stv, ok := v[2].(reflect.Type); ok {
						return t.New(getV).ToStructType(stv)
					}
				}
				return getV
			case Any:
				return getV
			default:
				return getV
			}
		}
		xfConfigs.Set(key, val)
		return val
	}
	return xfConfigs.Get(key)
}

// 读配置
func CR(key string, def ...interface{}) t.T {
	if xfConfigs.Contains(key) {
		return t.New(xfConfigs.Get(key))
	}
	if len(def) > 0 {
		return t.New(def[0])
	}
	return t.New(nil)
}

// 写配置
func CW(key string, val interface{}) {
	xfConfigs.Set(key, val)
}
